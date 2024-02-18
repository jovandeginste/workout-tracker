package database

import (
	"fmt"
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/openstreetmap"
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
	Date          string
	Center        MapCenter
	Address       *geo.Address
	TotalDistance float64
	TotalDuration time.Duration
	AverageSpeed  float64
	AverageTempo  float64
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
	Lat           float64
	Lng           float64
	Distance      float64
	TotalDistance float64
	Duration      time.Duration
	TotalDuration time.Duration
	Title         string
}

// center returns the center point (lat, lng) of gpx points
func center(gpxContent *gpx.GPX) MapCenter {
	lat, lng := 0.0, 0.0
	points := allGPXPoints(gpxContent)

	for _, pt := range points {
		lat += pt.Point.Latitude
		lng += pt.Point.Longitude
	}

	size := float64(len(points))

	return MapCenter{
		Lat: lat / size,
		Lng: lng / size,
	}
}

func (m *MapCenter) Address() *geo.Address {
	geocoder := openstreetmap.Geocoder()

	address, err := geocoder.ReverseGeocode(m.Lat, m.Lng)
	if err != nil {
		return nil
	}

	return address
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

func createMapData(gpxContent *gpx.GPX) *MapData {
	mapCenter := center(gpxContent)

	totalDistance := gpxContent.Tracks[0].Segments[0].Length3D()
	totalDuration := time.Duration(gpxContent.Tracks[0].Segments[0].Duration()) * time.Second

	updown := gpxContent.Tracks[0].Segments[0].UphillDownhill()

	data := &MapData{
		Name:          gpxName(gpxContent),
		Center:        mapCenter,
		Address:       mapCenter.Address(),
		TotalDistance: totalDistance,
		TotalDuration: totalDuration,
		AverageSpeed:  totalDistance / totalDuration.Seconds(),
		MaxSpeed:      gpxContent.Tracks[0].Segments[0].MovingData().MaxSpeed,
		PauzeDuration: time.Duration(gpxContent.Tracks[0].Segments[0].MovingData().StoppedTime) * time.Second,
		MinElevation:  gpxContent.Tracks[0].Segments[0].ElevationBounds().MinElevation,
		MaxElevation:  gpxContent.Tracks[0].Segments[0].ElevationBounds().MaxElevation,
		TotalUp:       updown.Uphill,
		TotalDown:     updown.Downhill,
	}

	return data
}

func gpxAsMapData(gpxContent *gpx.GPX) MapData {
	data := createMapData(gpxContent)
	points := allGPXPoints(gpxContent)

	totalDist := 0.0
	totalTime := 0.0
	prevPoint := points[0]
	speedMPS := 0.0

	for i, pt := range points {
		dist := 0.0
		t := 0.0

		if i > 0 {
			dist = distanceBetween(prevPoint, pt)
			t = pt.TimeDiff(&prevPoint)
			speedMPS = pt.SpeedBetween(&prevPoint, true)

			prevPoint = pt
		}

		totalDist += dist
		totalTime += t

		title := fmt.Sprintf(
			"<b>Time:</b> %s<br/><b>Distance:</b> %.2f km<br/><b>Duration:</b> %s<br/><b>Speed:</b> %.2f km/h<br /><b>Height:</b> %.2f m",
			pt.Timestamp.Format("15:04"), // HH:MM
			totalDist/1000,
			time.Duration(totalTime)*time.Second,
			speedMPS,
			pt.Elevation.Value(),
		)

		data.Points = append(data.Points, MapPoint{
			Lat:           pt.Point.Latitude,
			Lng:           pt.Point.Longitude,
			Distance:      dist,
			TotalDistance: totalDist,
			Duration:      time.Duration(t) * time.Second,
			TotalDuration: time.Duration(totalTime) * time.Second,
			Title:         title,
		})
	}

	return *data
}
