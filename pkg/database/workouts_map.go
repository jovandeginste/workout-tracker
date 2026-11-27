package database

import (
	"math"
	"slices"
	"time"

	"github.com/biter777/countries"
	"github.com/codingsince1985/geo-golang"
	"github.com/jovandeginste/workout-tracker/v2/pkg/converters"
	"github.com/jovandeginste/workout-tracker/v2/pkg/geocoder"
	"github.com/labstack/gommon/log"
	"github.com/paulmach/orb"
	"github.com/spf13/cast"
	"github.com/tkrajina/gpxgo/gpx"
	"github.com/westphae/geomag/pkg/egm96"
	"gorm.io/gorm"
)

const UnknownLocation = "(unknown location)"

var correctAltitudeCreators = []string{
	"garmin", "Garmin", "Garmin Connect",
	"Apple Watch", "Open GPX Tracker for iOS",
	"StravaGPX iPhone", "StravaGPX",
	"suunto",
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
	if math.IsNaN(lat) || math.IsNaN(long) {
		return 0
	}

	if alt == 0 {
		return 0
	}

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
	Climbs        []Segment `gorm:"serializer:json" json:"climbs"`         // Auto-detected climbs
	converters.WorkoutData
}

type MapDataDetails struct {
	Model

	MapData *MapData   `gorm:"foreignKey:MapDataID" json:"-"`
	Points  []MapPoint `gorm:"serializer:json" json:"points"` // The GPS points of the workout

	HeatmapPoints []HeatmapPoint `gorm:"serializer:json" json:"-"` // Thinned-out points for the heatmap; NULL until calculated

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

	ExtraMetrics    ExtraMetrics  `json:"extraMetrics"`    // Extra metrics at this point
	Lat             float64       `json:"lat"`             // The latitude of the point
	Lng             float64       `json:"lng"`             // The longitude of the point
	Elevation       float64       `json:"elevation"`       // The elevation of the point
	Distance        float64       `json:"distance"`        // The distance from the previous point
	Distance2D      float64       `json:"distance2D"`      // The 2D distance from the previous point
	TotalDistance   float64       `json:"totalDistance"`   // The total distance of the workout up to this point
	TotalDistance2D float64       `json:"totalDistance2D"` // The total 2D distance of the workout up to this point
	Duration        time.Duration `json:"duration"`        // The duration from the previous point
	TotalDuration   time.Duration `json:"totalDuration"`   // The total duration of the workout up to this point
	SlopeGrade      float64       `json:"slopeGrade"`      // The grade of the slope at this point
}

func (m *MapCenter) ToOrbPoint() *orb.Point {
	return &orb.Point{m.Lng, m.Lat}
}

func (m *MapPoint) ToOrbPoint() *orb.Point {
	return &orb.Point{m.Lng, m.Lat}
}

func (m *MapPoint) IsZero() bool {
	return m.Lat == 0 && m.Lng == 0
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

func addressIsUnset(a *geo.Address) bool {
	if a == nil {
		return true
	}

	if a.Country == "" {
		return true
	}

	return false
}

func (m *MapData) UpdateAddress() {
	if addressIsUnset(m.Address) && !m.Center.IsZero() {
		m.Address = m.Center.Address()
	}

	if addressIsUnset(m.Address) && m.hasAddressString() {
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
	if addressIsUnset(m.Address) {
		return UnknownLocation
	}

	r := ""
	if m.Address.CountryCode != "" {
		e := countries.ByName(m.Address.CountryCode).Emoji()
		r += e + " "
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
	if m.Details != nil {
		if err := db.Save(m.Details).Error; err != nil {
			return err
		}
	}

	return db.Save(m).Error
}

func (m *MapPoint) AverageSpeed() float64 {
	if m.Duration.Seconds() == 0 {
		return 0
	}

	return m.Distance / m.Duration.Seconds()
}

// stoppedSpeedThreshold is the speed (in km/h) at or below which the athlete is
// considered to be standing still. It matches the threshold gpxgo uses to
// calculate the stopped time.
const stoppedSpeedThreshold = 1.0

// IsMoving reports whether the athlete was moving between the previous point
// and this point.
func (m *MapPoint) IsMoving() bool {
	return m.AverageSpeed()*3.6 > stoppedSpeedThreshold
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

	var lat, lng, tot float64

	for _, pt := range points {
		if !pointHasLocation(&pt) {
			continue
		}

		lat += pt.Point.Latitude
		lng += pt.Point.Longitude
		tot++
	}

	if tot == 0 {
		return MapCenter{}
	}

	mc := MapCenter{
		Lat: lat / tot,
		Lng: lng / tot,
	}

	mc.updateTimezone()

	return mc
}

func (m *MapCenter) updateTimezone() {
	m.TZ = ""

	if tzFinder != nil {
		m.TZ = tzFinder.GetTimezoneName(m.Lng, m.Lat)
	}

	if m.TZ == "" {
		m.TZ = time.UTC.String()
	}
}

func (m *MapCenter) IsZero() bool {
	return m.Lat == 0 && m.Lng == 0
}

func (m *MapCenter) Address() *geo.Address {
	if m.IsZero() {
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
				points = append(points, p)
			}
		}
	}

	return points
}

func totalDistanceFromExtraMetrics(p *gpx.GPXPoint) float64 {
	extraMetrics := ExtraMetrics{}
	extraMetrics.ParseGPXExtensions(p.Extensions)

	d, ok := extraMetrics["distance"]
	if !ok || d <= 0 {
		return 0
	}

	return d
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

func distance2DBetween(p1 gpx.GPXPoint, p2 gpx.GPXPoint) float64 {
	return p2.Distance2D(&p1)
}

func distance3DBetween(p1 gpx.GPXPoint, p2 gpx.GPXPoint) float64 {
	return p2.Distance3D(&p1)
}

func maxSpeedForSegment(segment gpx.GPXTrackSegment, includeMovingData bool) float64 {
	ms := 0.0

	if includeMovingData {
		ms = segment.MovingData().MaxSpeed
	}

	for _, p := range segment.Points {
		extraMetrics := ExtraMetrics{}
		extraMetrics.ParseGPXExtensions(p.Extensions)
		if newMS, ok := extraMetrics["speed"]; ok {
			if newMS > ms {
				ms = newMS
			}
		}
	}

	return ms
}

func stoppedDuration2D(segment gpx.GPXTrackSegment) float64 {
	if len(segment.Points) < 2 {
		return 0
	}

	var stopped float64

	for i, pt := range segment.Points[1:] {
		prev := &segment.Points[i]

		if speed := pt.SpeedBetween(prev, false) * 3.6; speed <= stoppedSpeedThreshold {
			stopped += pt.TimeDiff(prev)
		}
	}

	return stopped
}

func createMapData(gpxContent *gpx.GPX, ignoreElevation bool) *MapData {
	if len(gpxContent.Tracks) == 0 {
		return nil
	}

	var (
		maxElevation, uphill, downhill, maxSpeed float64
		pauseDuration                            time.Duration
	)

	minElevation := 100000.0 // This should be high enough for Earthly workouts

	for _, track := range gpxContent.Tracks {
		for _, segment := range track.Segments {
			if len(segment.Points) == 0 {
				continue
			}

			if ignoreElevation {
				pauseDuration += time.Duration(stoppedDuration2D(segment) * float64(time.Second))
				maxSpeed = max(maxSpeed, maxSpeedForSegment(segment, false))
				continue
			}

			pauseDuration += time.Duration(segment.MovingData().StoppedTime) * time.Second
			minElevation = min(minElevation, segment.ElevationBounds().MinElevation)
			maxElevation = max(maxElevation, segment.ElevationBounds().MaxElevation)
			uphill += segment.UphillDownhill().Uphill
			downhill += segment.UphillDownhill().Downhill
			maxSpeed = max(maxSpeed, maxSpeedForSegment(segment, true))
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
	}

	if !mapCenter.IsZero() {
		data.WorkoutData = converters.WorkoutData{
			MaxSpeed:      maxSpeed,
			PauseDuration: pauseDuration,
			MinElevation:  correctAltitude(gpxContent.Creator, mapCenter.Lat, mapCenter.Lng, minElevation),
			MaxElevation:  correctAltitude(gpxContent.Creator, mapCenter.Lat, mapCenter.Lng, maxElevation),
			TotalUp:       uphill,
			TotalDown:     downhill,
		}
	}

	addExtraMetrics(gpxContent, &data.WorkoutData)

	data.correctNaN()

	return data
}

func addExtraMetrics(gpxContent *gpx.GPX, data *converters.WorkoutData) {
	if tc, ok := gpxContent.Extensions.GetNode(gpx.AnyNamespace, "total-calories"); ok {
		data.TotalCalories = cast.ToFloat64(tc.Data)
	}
}

func (m *MapData) correctNaN() {
	for _, value := range []*float64{
		&m.MinElevation,
		&m.MaxElevation,
		&m.TotalDistance,
		&m.TotalDistance2D,
		&m.TotalDown,
		&m.TotalUp,
	} {
		if math.IsNaN(*value) {
			*value = 0
		}
	}
}

func maxSpeedForPoints(points []MapPoint, maxDeltaMeter float64) float64 {
	if len(points) < 2 {
		return 0
	}

	if maxDeltaMeter <= 0 {
		maxDeltaMeter = DefaultMaxDeltaMeter
	}

	maxSpeed := 0.0

	for center := range points {
		left, ok := maxSpeedWindow(points, center, maxDeltaMeter, -1)
		if !ok {
			continue
		}

		right, ok := maxSpeedWindow(points, center, maxDeltaMeter, 1)
		if !ok {
			continue
		}

		leftPoint := points[left]
		rightPoint := points[right]

		dist := rightPoint.TotalDistance2D - leftPoint.TotalDistance2D
		dur := rightPoint.TotalDuration - leftPoint.TotalDuration
		if dur <= 0 {
			continue
		}

		speed := dist / dur.Seconds()
		if speed > maxSpeed {
			maxSpeed = speed
		}
	}

	return maxSpeed
}

func maxSpeedWindow(points []MapPoint, center int, maxDeltaMeter float64, step int) (int, bool) {
	centerPoint := points[center]
	index := center

	for index+step >= 0 && index+step < len(points) {
		next := index + step
		if math.Abs(points[next].TotalDistance2D-centerPoint.TotalDistance2D) >= maxDeltaMeter {
			index = next
			break
		}

		index = next
	}

	if math.Abs(points[index].TotalDistance2D-centerPoint.TotalDistance2D) < maxDeltaMeter {
		return 0, false
	}

	return index, true
}

// distanceTracker accumulates the distance between consecutive track points.
// It uses the cumulative distance that the device recorded when available and
// falls back to the track geometry otherwise.
type distanceTracker struct {
	ignoreElevation    bool
	ignoreDeviceValues bool
	totalDistance      float64
	totalDistance2D    float64
	prevDeviceDistance float64
	hasPrevDevice      bool
}

func (t *distanceTracker) add(distance, distance2D float64) {
	t.totalDistance += distance
	t.totalDistance2D += distance2D
}

func (t *distanceTracker) observe(pt *gpx.GPXPoint) {
	if t.ignoreDeviceValues {
		return
	}

	t.prevDeviceDistance, t.hasPrevDevice = deviceDistance(pt)
}

func (t *distanceTracker) step(prev, pt *gpx.GPXPoint) (float64, float64) {
	if !t.ignoreDeviceValues {
		deviceDist, hasDeviceDist := deviceDistance(pt)

		if hasDeviceDist && t.hasPrevDevice && deviceDist >= t.prevDeviceDistance {
			distance := deviceDist - t.prevDeviceDistance

			t.prevDeviceDistance = deviceDist
			t.hasPrevDevice = true

			return distance, distance
		}

		t.prevDeviceDistance = deviceDist
		t.hasPrevDevice = hasDeviceDist
	}

	if pointHasLocation(prev) && pointHasLocation(pt) {
		distance2D := distance2DBetween(*prev, *pt)
		if t.ignoreElevation {
			return distance2D, distance2D
		}

		return distance3DBetween(*prev, *pt), distance2D
	}

	distance := totalDistanceFromExtraMetrics(pt) - t.totalDistance
	if distance < 0 {
		distance = 0
	}

	return distance, distance
}

func gpxAsMapData(gpxContent *gpx.GPX, workoutType WorkoutType, ignoreElevation, ignoreDeviceValues bool) *MapData {
	data := createMapData(gpxContent, ignoreElevation)

	points := allGPXPoints(gpxContent)
	if len(points) < 2 {
		return data
	}

	tracker := &distanceTracker{
		ignoreElevation:    ignoreElevation,
		ignoreDeviceValues: ignoreDeviceValues,
	}
	tracker.observe(&points[0])

	var (
		totalTime time.Duration
		prevPt    *gpx.GPXPoint
	)

	data.Details = &MapDataDetails{}

	for _, pt := range points {
		newPoint := MapPoint{
			Elevation: pt.Elevation.Value(),
			Time:      pt.Timestamp,
		}

		validGPS := pointHasLocation(&pt)

		if validGPS {
			newPoint.Lat = pt.Point.Latitude
			newPoint.Lng = pt.Point.Longitude
		}

		if prevPt != nil {
			newPoint.Distance, newPoint.Distance2D = tracker.step(prevPt, &pt)
			newPoint.Duration = pt.Timestamp.Sub(prevPt.Timestamp)

			tracker.add(newPoint.Distance, newPoint.Distance2D)
			totalTime += newPoint.Duration
		}

		newPoint.TotalDistance = tracker.totalDistance
		newPoint.TotalDistance2D = tracker.totalDistance2D
		newPoint.TotalDuration = totalTime

		extraMetrics := ExtraMetrics{}

		if validGPS && pt.Elevation.NotNull() {
			extraMetrics.Set("elevation", correctAltitude(gpxContent.Creator, pt.Point.Latitude, pt.Point.Longitude, pt.Elevation.Value()))
		}

		extraMetrics.ParseGPXExtensions(pt.Extensions)

		newPoint.ExtraMetrics = extraMetrics
		prevPt = &pt

		data.Details.Points = append(data.Details.Points, newPoint)
	}

	data.TotalDistance = tracker.totalDistance
	data.TotalDistance2D = tracker.totalDistance2D
	data.TotalDuration = totalTime

	applyDeviceValues(gpxContent, data, ignoreElevation, ignoreDeviceValues)
	data.updateMaxSpeed(gpxContent, workoutType, ignoreDeviceValues)

	if totalTime > 0 {
		data.AverageSpeed = data.TotalDistance / totalTime.Seconds()

		pdDiff := totalTime - data.PauseDuration
		if pdDiff > 0 {
			data.AverageSpeedNoPause = data.TotalDistance / pdDiff.Seconds()
		}
	}

	data.correctNaN()
	data.UpdateExtraMetrics()
	data.Details.UpdateHeatmapPoints()

	return data
}

func pointHasLocation(pt *gpx.GPXPoint) bool {
	if pt == nil {
		return false
	}

	if math.IsNaN(pt.Latitude) || math.IsNaN(pt.Longitude) {
		return false
	}

	if pt.Latitude == 0 && pt.Longitude == 0 {
		return false
	}

	return true
}
