package user

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/tkrajina/gpxgo/gpx"
)

// parseGPX parses a GPX file, returns GPX.
func parseGPX(gpxBytes []byte) (*gpx.GPX, error) {
	gpxContent, err := gpx.ParseBytes(gpxBytes)
	if err != nil {
		return nil, err
	}

	return gpxContent, nil
}

type MapData struct {
	Name          string
	Type          string
	Date          string
	Center        MapCenter
	TotalDistance float64
	TotalDuration time.Duration
	AverageSpeed  float64
	MaxSpeed      float64
	PauzeDuration time.Duration
	MinElevation  float64
	MaxElevation  float64
	TotalUp       float64
	TotalDown     float64
	Points        []MapPoint
}
type MapCenter struct {
	Lat float64
	Lng float64
}
type MapPoint struct {
	Lat   float64
	Lng   float64
	Title string
}

// center returns the center point (lat, lng) of gpx points
func center(gpxContent *gpx.GPX) (float64, float64) {
	lat, lng := 0.0, 0.0
	points := allGPXPoints(gpxContent)

	for _, pt := range points {
		lat += pt.Point.Latitude
		lng += pt.Point.Longitude
	}

	size := float64(len(points))

	return lat / size, lng / size
}

// allGPXPoints returns the first track segment's points
func allGPXPoints(gpxContent *gpx.GPX) []gpx.GPXPoint {
	return gpxContent.Tracks[0].Segments[0].Points
}

func gpxName(gpxContent *gpx.GPX) string {
	if gpxContent.Name != "" {
		return gpxContent.Name
	}

	if len(gpxContent.Tracks) == 0 {
		return "(no name)"
	}

	return gpxContent.Tracks[0].Name
}

func distanceBetween(p1 gpx.GPXPoint, p2 gpx.GPXPoint) float64 {
	return gpx.HaversineDistance(p1.Latitude, p1.Longitude, p2.Latitude, p2.Longitude)
}

func gpxAsMapData(gpxContent *gpx.GPX) MapData {
	points := allGPXPoints(gpxContent)

	clat, clng := center(gpxContent)

	totalDistance := gpxContent.Tracks[0].Segments[0].Length3D()
	totalDuration := time.Duration(gpxContent.Tracks[0].Segments[0].Duration()) * time.Second

	updown := gpxContent.Tracks[0].Segments[0].UphillDownhill()

	data := MapData{
		Name: gpxName(gpxContent),
		Type: "running",
		Date: fmt.Sprintf("%s (%s)", gpxContent.Time.Local().Format("2006-01-02 15:04"), humanize.Time(*gpxContent.Time)),
		Center: MapCenter{
			Lat: clat,
			Lng: clng,
		},
		TotalDistance: totalDistance / 1000.0,
		TotalDuration: totalDuration,
		AverageSpeed:  3.6 * totalDistance / totalDuration.Seconds(),
		MaxSpeed:      3.6 * gpxContent.Tracks[0].Segments[0].MovingData().MaxSpeed,
		PauzeDuration: time.Duration(gpxContent.Tracks[0].Segments[0].MovingData().StoppedTime) * time.Second,
		MinElevation:  gpxContent.Tracks[0].Segments[0].ElevationBounds().MinElevation,
		MaxElevation:  gpxContent.Tracks[0].Segments[0].ElevationBounds().MaxElevation,
		TotalUp:       updown.Uphill,
		TotalDown:     updown.Downhill,
	}

	dist := 0.0
	t := 0.0
	prevPoint := points[0]
	speedMPS := 0.0

	for i, pt := range points {
		if i > 0 {
			dist += distanceBetween(prevPoint, pt)
			t += pt.TimeDiff(&prevPoint)
			speedMPS = pt.SpeedBetween(&prevPoint, true)

			prevPoint = pt
		}

		title := fmt.Sprintf(
			"<b>Time:</b> %s<br/><b>Distance:</b> %.2f km<br/><b>Duration:</b> %s<br/><b>Speed:</b> %.2f km/h<br /><b>Height:</b> %.2f m",
			pt.Timestamp.Format("15:04"), // HH:MM
			dist/1000,
			time.Duration(t)*time.Second,
			3.6*speedMPS,
			pt.Elevation.Value(),
		)

		data.Points = append(data.Points, MapPoint{
			Lat:   pt.Point.Latitude,
			Lng:   pt.Point.Longitude,
			Title: title,
		})
	}

	return data
}
