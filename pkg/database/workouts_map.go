package database

import (
	"math"
	"slices"
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/jovandeginste/workout-tracker/v2/pkg/geocoder"
	"github.com/jovandeginste/workout-tracker/v2/pkg/templatehelpers"
	"github.com/labstack/gommon/log"
	"github.com/paulmach/orb"
	"github.com/ringsaturn/tzf"
	"github.com/tkrajina/gpxgo/gpx"
	"github.com/westphae/geomag/pkg/egm96"
	"gorm.io/gorm"
)

var online = true

const UnknownLocation = "(unknown location)"

var correctAltitudeCreators = []string{
	"garmin", "Garmin", "Garmin Connect",
	"Apple Watch", "Open GPX Tracker for iOS",
	"StravaGPX iPhone", "StravaGPX",
	"Workout Tracker",
}

func creatorNeedsCorrection(creator string) bool {
	return !slices.Contains(correctAltitudeCreators, creator)
}

func normalizeDegrees(val float64) float64 {
	if val < 0 {
		return val + 360
	}

	return val
}

func correctAltitude(creator string, lat, long, alt float64) float64 {
	if !creatorNeedsCorrection(creator) {
		return alt
	}

	lat = normalizeDegrees(lat)
	long = normalizeDegrees(long)

	loc := egm96.NewLocationGeodetic(lat, long, alt)

	h, err := loc.HeightAboveMSL()
	if err != nil {
		return alt
	}

	return h
}

type MapData struct {
	Model
	Address *geo.Address    `gorm:"serializer:json" json:"address"`                       // The address of the workout
	Details *MapDataDetails `gorm:"constraint:OnDelete:CASCADE" json:"details,omitempty"` // The details of the workout

	Workout       *Workout  `gorm:"foreignKey:WorkoutID" json:"-"`         // The user who owns this profile
	Creator       string    `json:"creator"`                               // The tool that created this workout
	AddressString string    `json:"addressString"`                         // The generic location of the workout
	Center        MapCenter `gorm:"serializer:json" json:"center"`         // The center of the workout (in coordinates)
	WorkoutID     uint64    `gorm:"not null;uniqueIndex" json:"workoutID"` // The workout this data belongs to
	converters.WorkoutData
}

type MapDataDetails struct {
	Model

	MapData *MapData   `gorm:"foreignKey:MapDataID" json:"-"`
	Points  []MapPoint `gorm:"serializer:json" json:"points"` // The GPS points of the workout

	MapDataID uint64 `gorm:"not null;uniqueIndex" json:"mapDataID"` // The ID of the map data these details belong to
}

// MapCenter is the center of the workout
type MapCenter struct {
	TZ  string  `json:"tz"`  // Timezone
	Lat float64 `json:"lat"` // Latitude
	Lng float64 `json:"lng"` // Longitude
}

type MapPoint struct {
	Time time.Time `json:"time"` // The time the point was recorded

	ExtraMetrics  ExtraMetrics  `json:"extraMetrics"`  // Extra metrics at this point
	Lat           float64       `json:"lat"`           // The latitude of the point
	Lng           float64       `json:"lng"`           // The longitude of the point
	Elevation     float64       `json:"elevation"`     // The elevation of the point
	Distance      float64       `json:"distance"`      // The distance from the previous point
	TotalDistance float64       `json:"totalDistance"` // The total distance of the workout up to this point
	Duration      time.Duration `json:"duration"`      // The duration from the previous point
	TotalDuration time.Duration `json:"totalDuration"` // The total duration of the workout up to this point
}

func (m *MapCenter) ToOrbPoint() *orb.Point {
	return &orb.Point{m.Lng, m.Lat}
}

func (m *MapPoint) ToOrbPoint() *orb.Point {
	return &orb.Point{m.Lng, m.Lat}
}

func (d *MapDataDetails) Save(db *gorm.DB) error {
	return db.Save(d).Error
}

func (m *MapData) UpdateExtraMetrics() {
	if m.Details == nil ||
		len(m.Details.Points) == 0 {
		return
	}

	metrics := []string{}
	found := map[string]bool{}

	for _, d := range m.Details.Points {
		for k := range d.ExtraMetrics {
			if found[k] {
				continue
			}

			metrics = append(metrics, k)
			found[k] = true
		}
	}

	slices.Sort(metrics)

	m.ExtraMetrics = metrics
}

func (m *MapData) UpdateAddress() {
	if m.Address == nil && !m.Center.IsZero() {
		m.Address = m.Center.Address()
	}

	if m.Address == nil && m.hasAddressString() {
		return
	}

	m.AddressString = m.addressString()
}

func (m *MapData) hasAddressString() bool {
	switch m.AddressString {
	case "", UnknownLocation:
		return false
	default:
		return true
	}
}

func (m *MapData) addressString() string {
	if m.Address == nil {
		return UnknownLocation
	}

	r := ""
	if m.Address.CountryCode != "" {
		r += templatehelpers.CountryToFlag(m.Address.CountryCode) + " "
	}

	switch {
	case m.Address.City != "":
		r += m.Address.City
	case m.Address.Street != "":
		r += m.Address.Street
	default:
		return r + m.Address.FormattedAddress
	}

	if shouldAddState(m.Address) {
		r += ", " + m.Address.State
	}

	return r
}

func shouldAddState(address *geo.Address) bool {
	return address.CountryCode == "US"
}

func (m *MapData) Save(db *gorm.DB) error {
	return db.Save(m).Error
}

func (m *MapPoint) AverageSpeed() float64 {
	if m.Duration.Seconds() == 0 {
		return 0
	}

	return m.Distance / m.Duration.Seconds()
}

func (m *MapPoint) DistanceTo(m2 *MapPoint) float64 {
	if m == nil || m2 == nil {
		return math.Inf(1)
	}

	return m.AsGPXPoint().Distance2D(m2.AsGPXPoint())
}

func (m *MapPoint) AsGPXPoint() *gpx.Point {
	ele := gpx.NewNullableFloat64(m.Elevation)

	return &gpx.Point{Latitude: m.Lat, Longitude: m.Lng, Elevation: *ele}
}

// center returns the center point (lat, lng) of gpx points
func center(gpxContent *gpx.GPX) MapCenter {
	points := allGPXPoints(gpxContent)

	if len(points) == 0 {
		return MapCenter{}
	}

	lat, lng := 0.0, 0.0

	for _, pt := range points {
		lat += pt.Point.Latitude
		lng += pt.Point.Longitude
	}

	size := float64(len(points))

	mc := MapCenter{
		Lat: lat / size,
		Lng: lng / size,
	}

	mc.updateTimezone()

	return mc
}

func (m *MapCenter) updateTimezone() {
	finder, err := tzf.NewDefaultFinder()
	if err != nil {
		m.TZ = time.UTC.String()
		return
	}

	tz := finder.GetTimezoneName(m.Lng, m.Lat)
	if tz == "" {
		m.TZ = time.UTC.String()
		return
	}

	m.TZ = tz
}

func (m *MapCenter) IsZero() bool {
	return m.Lat == 0 && m.Lng == 0
}

func (m *MapCenter) Address() *geo.Address {
	if !online || m.IsZero() {
		return nil
	}

	r, err := geocoder.Reverse(geocoder.Query{
		Lat:    m.Lat,
		Lon:    m.Lng,
		Format: "json",
	})
	if err != nil {
		log.Warn("Error performing reverse geocode: ", err)
		return nil
	}

	return r
}

// allGPXPoints returns the first track segment's points
func allGPXPoints(gpxContent *gpx.GPX) []gpx.GPXPoint {
	if gpxContent == nil {
		return nil
	}

	var points []gpx.GPXPoint

	for _, track := range gpxContent.Tracks {
		for _, segment := range track.Segments {
			for _, p := range segment.Points {
				if !pointHasDistance(p) {
					continue
				}

				points = append(points, p)
			}
		}
	}

	return points
}

func pointHasDistance(p gpx.GPXPoint) bool {
	if math.IsNaN(p.Latitude) || math.IsNaN(p.Longitude) {
		return false
	}

	return true
}

// Determines the date to use for the workout
func gpxDate(gpxContent *gpx.GPX) *time.Time {
	// Use the first track's first segment's timestamp if it exists
	// This is the best time to use as a start time, since converters shouldn't
	// touch this timestamp
	if len(gpxContent.Tracks) > 0 {
		if t := gpxContent.Tracks[0]; len(t.Segments) > 0 {
			if s := t.Segments[0]; len(s.Points) > 0 {
				if !s.Points[0].Timestamp.IsZero() {
					return &s.Points[0].Timestamp
				}
			}
		}
	}

	// Otherwise, return the timestamp from the metadata, use that (not all apps have
	// this, notably Workoutdoors doesn't)
	// If this is nil, this should result in an error and the user will be alerted.
	return gpxContent.Time
}

func distanceBetween(p1 gpx.GPXPoint, p2 gpx.GPXPoint) float64 {
	return p2.Distance3D(&p1)
}

func createMapData(gpxContent *gpx.GPX) *MapData {
	if len(gpxContent.Tracks) == 0 {
		return nil
	}

	var (
		totalDistance, maxElevation, uphill, downhill, maxSpeed float64
		totalDuration, pauseDuration                            time.Duration
	)

	minElevation := 100000.0 // This should be high enough for Earthly workouts

	for _, track := range gpxContent.Tracks {
		for _, segment := range track.Segments {
			if len(segment.Points) == 0 {
				continue
			}

			totalDistance += segment.Length3D()
			totalDuration += time.Duration(segment.Duration()) * time.Second
			pauseDuration += (time.Duration(segment.MovingData().StoppedTime)) * time.Second
			minElevation = min(minElevation, segment.ElevationBounds().MinElevation)
			maxElevation = max(maxElevation, segment.ElevationBounds().MaxElevation)
			uphill += segment.UphillDownhill().Uphill
			downhill += segment.UphillDownhill().Downhill
			maxSpeed = max(maxSpeed, segment.MovingData().MaxSpeed)
			pauseDuration += time.Duration(segment.MovingData().StoppedTime)
		}
	}

	// Make sure minElevation is never higher than maxElevation
	minElevation = min(minElevation, maxElevation)

	// Now reduce the whole GPX to a single track to calculate the center
	gpxContent.ReduceGpxToSingleTrack()
	mapCenter := center(gpxContent)

	data := &MapData{
		Creator: gpxContent.Creator,
		Center:  mapCenter,
		WorkoutData: converters.WorkoutData{
			TotalDistance:       totalDistance,
			TotalDuration:       totalDuration,
			MaxSpeed:            maxSpeed,
			AverageSpeed:        totalDistance / totalDuration.Seconds(),
			AverageSpeedNoPause: totalDistance / (totalDuration - pauseDuration).Seconds(),
			PauseDuration:       pauseDuration,
			MinElevation:        correctAltitude(gpxContent.Creator, mapCenter.Lat, mapCenter.Lng, minElevation),
			MaxElevation:        correctAltitude(gpxContent.Creator, mapCenter.Lat, mapCenter.Lng, maxElevation),
			TotalUp:             uphill,
			TotalDown:           downhill,
		},
	}

	data.UpdateAddress()
	data.UpdateExtraMetrics()
	data.correctNaN()

	return data
}

func (m *MapData) correctNaN() {
	if math.IsNaN(m.MinElevation) {
		m.MinElevation = 0
	}

	if math.IsNaN(m.MaxElevation) {
		m.MaxElevation = 0
	}

	if math.IsNaN(m.TotalDistance) {
		m.TotalDistance = 0
	}

	if math.IsNaN(m.TotalDown) {
		m.TotalDown = 0
	}

	if math.IsNaN(m.TotalUp) {
		m.TotalUp = 0
	}
}

func gpxAsMapData(gpxContent *gpx.GPX) *MapData {
	data := createMapData(gpxContent)

	points := allGPXPoints(gpxContent)
	if len(points) == 0 {
		return data
	}

	totalDist := 0.0
	totalTime := 0.0
	prevPoint := points[0]

	data.Details = &MapDataDetails{}

	for i, pt := range points {
		if !pointHasDistance(pt) {
			continue
		}

		dist := 0.0
		t := 0.0

		if i > 0 {
			dist = distanceBetween(prevPoint, pt)
			t = pt.TimeDiff(&prevPoint)

			prevPoint = pt

			totalDist += dist
			totalTime += t
		}

		extraMetrics := ExtraMetrics{}
		extraMetrics.Set("elevation", correctAltitude(gpxContent.Creator, pt.Point.Latitude, pt.Point.Longitude, pt.Elevation.Value()))
		extraMetrics.ParseGPXExtensions(pt.Extensions)

		data.Details.Points = append(data.Details.Points, MapPoint{
			Lat:           pt.Point.Latitude,
			Lng:           pt.Point.Longitude,
			Elevation:     pt.Elevation.Value(),
			Time:          pt.Timestamp,
			Distance:      dist,
			TotalDistance: totalDist,
			Duration:      time.Duration(t) * time.Second,
			TotalDuration: time.Duration(totalTime) * time.Second,
			ExtraMetrics:  extraMetrics,
		})
	}

	data.correctNaN()

	return data
}
