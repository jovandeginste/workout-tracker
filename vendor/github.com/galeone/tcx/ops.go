package tcx

import (
	"math"
	"time"

	"github.com/philhofer/vec"
)

type TrackSpline struct {
	lat  *vec.CubicSplineInterpolation
	long *vec.CubicSplineInterpolation
	spd  *vec.CubicSplineInterpolation
	alt  *vec.CubicSplineInterpolation
}

func (t *TrackSpline) Lat(dist float64) float64 {
	return t.lat.F(dist)
}

func (t *TrackSpline) Long(dist float64) float64 {
	return t.long.F(dist)
}

func (t *TrackSpline) Speed(dist float64) float64 {
	return t.spd.F(dist)
}

func (t *TrackSpline) Alt(dist float64) float64 {
	return t.alt.F(dist)
}

type LapZeroDistError struct{}

func (l *LapZeroDistError) Error() string {
	return "Lap Distance Cannot Be Zero."
}

type LapZeroTimeError struct{}

func (l *LapZeroTimeError) Error() string {
	return "Lap Time Cannot be <= 0"
}

// Match lap time to "sec" (in-place)
func (lap *Lap) MatchTime(sec float64) error {
	isec := int(math.Floor(sec))
	if isec <= 0 {
		return new(LapZeroTimeError)
	}

	ldist := lap.Dist
	if ldist == 0 {
		return new(LapZeroDistError)
	}

	maxs := lap.MaxSpeed
	startTime := lap.Trk.Pt[0].Time
	startDist := lap.Trk.Pt[0].Dist
	vfactor := lap.TotalTime / sec
	trkspl := Spline(lap.Trk)
	newtrk := make([]Trackpoint, isec)
	onesec, err := time.ParseDuration("1.0s")
	if err != nil {
		panic(err)
	}

	for i := 0; i < isec; i++ {
		thisSpeed := trkspl.Speed(startDist) * vfactor
		if thisSpeed > maxs {
			lap.MaxSpeed = thisSpeed
		}
		newtrk[i] = Trackpoint{
			Time:  startTime,
			Lat:   trkspl.Lat(startDist),
			Long:  trkspl.Long(startDist),
			Speed: thisSpeed,
			Alt:   trkspl.Alt(startDist),
		}
		startDist += thisSpeed
		startTime.Add(onesec)
	}

	lap.TotalTime = float64(isec)
	lap.Trk = &Track{Pt: newtrk}
	return nil
}

func Spline(trk *Track) *TrackSpline {
	Npts := len(trk.Pt)
	lats := make([]float64, Npts)
	longs := make([]float64, Npts)
	spds := make([]float64, Npts)
	alts := make([]float64, Npts)
	dists := make([]float64, Npts)
	for i := 0; i < Npts; i++ {
		lats[i] = trk.Pt[i].Lat
		longs[i] = trk.Pt[i].Long
		spds[i] = trk.Pt[i].Speed
		alts[i] = trk.Pt[i].Alt
		dists[i] = trk.Pt[i].Dist
	}
	latsplin := vec.CubicSpline(vec.MakeBiVariateData(dists, lats))
	longssplin := vec.CubicSpline(vec.MakeBiVariateData(dists, longs))
	spdssplin := vec.CubicSpline(vec.MakeBiVariateData(dists, spds))
	altssplin := vec.CubicSpline(vec.MakeBiVariateData(dists, alts))
	return &TrackSpline{lat: latsplin, long: longssplin, spd: spdssplin, alt: altssplin}
}
