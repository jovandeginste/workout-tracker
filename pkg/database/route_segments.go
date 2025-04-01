package database

import (
	"crypto/sha256"
	"path/filepath"
	"sort"
	"strings"

	"github.com/codingsince1985/geo-golang"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/microcosm-cc/bluemonday"
	"github.com/tkrajina/gpxgo/gpx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RoutSegmentCreationParams struct {
	Name  string `form:"name"`
	Start int    `form:"start"`
	End   int    `form:"end"`
}

func (rscp *RoutSegmentCreationParams) Filename() string {
	if rscp.Name == "" {
		return "noname.gpx"
	}

	if strings.HasSuffix(rscp.Name, ".gpx") {
		return rscp.Name
	}

	return rscp.Name + ".gpx"
}

type RouteSegment struct {
	Model
	GeoAddress    *geo.Address `gorm:"serializer:json" json:"geoAddress"` // The address of the workout
	Name          string       `gorm:"not null" json:"name"`              // The name of the workout
	Notes         string       `json:"notes"`                             // The notes associated with the workout, in markdown
	AddressString string       `json:"addressString"`                     // The generic location of the workout
	Filename      string       `json:"filename"`                          // The filename of the file

	Points []MapPoint `gorm:"serializer:json" json:"points"` // The GPS points of the workout

	Content             []byte               `gorm:"type:bytes" json:"content"`            // The file content
	Checksum            []byte               `gorm:"not null;uniqueIndex" json:"checksum"` // The checksum of the content
	RouteSegmentMatches []*RouteSegmentMatch `json:"routeSegmentMatches"`                  // The matches of the route segment
	Center              MapCenter            `gorm:"serializer:json" json:"center"`        // The center of the workout (in coordinates)

	TotalDistance float64 `json:"totalDistance"` // The total distance of the workout
	MinElevation  float64 `json:"minElevation"`  // The minimum elevation of the workout
	MaxElevation  float64 `json:"maxElevation"`  // The maximum elevation of the workout
	TotalUp       float64 `json:"totalUp"`       // The total distance up of the workout
	TotalDown     float64 `json:"totalDown"`     // The total distance down of the workout
	Bidirectional bool    `json:"bidirectional"` // Whether the route segment is bidirectional
	Circular      bool    `json:"circular"`      // Whether the route segment is circular

	Dirty bool `json:"dirty"` // Whether the route segment should be recalculated
}

func (rs *RouteSegment) HasFile() bool {
	return rs.Filename != "" && rs.Content != nil
}

func NewRouteSegment(notes string, filename string, content []byte) (*RouteSegment, error) {
	filename = filepath.Base(filename)
	name := strings.TrimSuffix(filename, ".gpx")

	h := sha256.New()
	h.Write(content)

	rs := &RouteSegment{
		Name:  name,
		Notes: notes,
		Dirty: true,

		Content:  content,
		Checksum: h.Sum(nil),
		Filename: filename,
	}

	if err := rs.UpdateFromContent(); err != nil {
		return nil, err
	}

	return rs, nil
}

func RouteSegmentFromPoints(workout *Workout, params *RoutSegmentCreationParams) ([]byte, error) {
	points := workout.Data.Details.Points[params.Start-1 : params.End-1]

	s := gpx.GPXTrackSegment{}

	for _, p := range points {
		gpxPoint := gpx.Point{
			Latitude:  p.Lat,
			Longitude: p.Lng,
			Elevation: *gpx.NewNullableFloat64(p.ExtraMetrics.Get("elevation")),
		}

		pt := gpx.GPXPoint{Point: gpxPoint}
		s.AppendPoint(&pt)
	}

	newFile := &gpx.GPX{
		Creator: "Workout Tracker",
		Tracks:  []gpx.GPXTrack{{Segments: []gpx.GPXTrackSegment{s}}},
	}

	content, err := newFile.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (rs *RouteSegment) UpdateFromContent() error {
	gpxContent, err := converters.Parse(rs.Filename, rs.Content)
	if err != nil {
		return err
	}

	gpxContent.GPX.SimplifyTracks(MaxDeltaMeter / 2)

	data := gpxAsMapData(gpxContent.GPX)

	if !data.Center.IsZero() {
		data.Address = data.Center.Address()

		rs.GeoAddress = data.Address
		rs.AddressString = data.addressString()
		rs.Center = data.Center
	}

	rs.TotalDistance = data.TotalDistance
	rs.MinElevation = data.MinElevation
	rs.MaxElevation = data.MaxElevation
	rs.TotalUp = data.TotalUp
	rs.TotalDown = data.TotalDown
	rs.Points = data.Details.Points

	return nil
}

func GetRouteSegment(db *gorm.DB, id int) (*RouteSegment, error) {
	var rs RouteSegment

	if err := db.Preload("RouteSegmentMatches.Workout.User").First(&rs, id).Error; err != nil {
		return nil, err
	}

	sort.Slice(rs.RouteSegmentMatches, func(i, j int) bool {
		return rs.RouteSegmentMatches[i].Workout.GetDate().Before(rs.RouteSegmentMatches[j].Workout.GetDate())
	})

	return &rs, nil
}

func (rs *RouteSegment) Delete(db *gorm.DB) error {
	return db.Select(clause.Associations).Delete(rs).Error
}

func (rs *RouteSegment) Create(db *gorm.DB) error {
	if rs.Content == nil {
		return ErrInvalidData
	}

	return db.Create(rs).Error
}

func (rs *RouteSegment) Save(db *gorm.DB) error {
	if rs.Content == nil {
		return ErrInvalidData
	}

	if rs.RouteSegmentMatches != nil {
		if err := db.Model(rs).Association("RouteSegmentMatches").Replace(rs.RouteSegmentMatches); err != nil {
			return err
		}
	}

	return db.Save(rs).Error
}

func GetRouteSegments(db *gorm.DB) ([]*RouteSegment, error) {
	var rs []*RouteSegment

	if err := db.Preload("RouteSegmentMatches.Workout").Order("created_at DESC").Find(&rs).Error; err != nil {
		return nil, err
	}

	return rs, nil
}

func (rs *RouteSegment) MarkdownNotes() string {
	doc := parser.NewWithExtensions(parser.CommonExtensions).Parse([]byte(rs.Notes))
	renderer := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})
	safeHTML := bluemonday.UGCPolicy().SanitizeBytes(markdown.Render(doc, renderer))

	return string(safeHTML)
}

func (rs *RouteSegment) Address() string {
	if rs.AddressString != "" {
		return rs.AddressString
	}

	if rs.GeoAddress != nil {
		return rs.GeoAddress.FormattedAddress
	}

	return UnknownLocation
}
