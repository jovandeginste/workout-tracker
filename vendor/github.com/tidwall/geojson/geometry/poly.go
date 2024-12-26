// Copyright 2018 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package geometry

type Poly struct {
	Exterior Ring
	Holes    []Ring
}

func NewPoly(exterior []Point, holes [][]Point, opts *IndexOptions) *Poly {
	poly := new(Poly)
	poly.Exterior = newRing(exterior, opts)
	if len(holes) > 0 {
		poly.Holes = make([]Ring, len(holes))
		for i := range holes {
			poly.Holes[i] = newRing(holes[i], opts)
		}
	}
	return poly
}

func (poly *Poly) Clockwise() bool {
	if poly == nil || poly.Exterior == nil {
		return false
	}
	return poly.Exterior.Clockwise()
}

func (poly *Poly) Empty() bool {
	if poly == nil || poly.Exterior == nil {
		return true
	}
	return poly.Exterior.Empty()
}

func (poly *Poly) Valid() bool {
	if !poly.Exterior.Valid() {
		return false
	}
	for _, hole := range poly.Holes {
		if !hole.Valid() {
			return false
		}
	}
	return true
}

func (poly *Poly) Rect() Rect {
	if poly == nil || poly.Exterior == nil {
		return Rect{}
	}
	return poly.Exterior.Rect()
}

// Move the polygon by delta. Returns a new polygon
func (poly *Poly) Move(deltaX, deltaY float64) *Poly {
	if poly == nil {
		return nil
	}
	if poly.Exterior == nil {
		return new(Poly)
	}
	npoly := new(Poly)
	if series, ok := poly.Exterior.(*baseSeries); ok {
		npoly.Exterior = Ring(series.Move(deltaX, deltaY))
	} else {
		nseries := makeSeries(
			seriesCopyPoints(poly.Exterior), false, true, DefaultIndexOptions)
		npoly.Exterior = Ring(nseries.Move(deltaX, deltaY))
	}
	if len(poly.Holes) > 0 {
		npoly.Holes = make([]Ring, len(poly.Holes))
		for i, hole := range poly.Holes {
			if series, ok := hole.(*baseSeries); ok {
				npoly.Holes[i] = Ring(series.Move(deltaX, deltaY))
			} else {
				nseries := makeSeries(
					seriesCopyPoints(hole), false, true, DefaultIndexOptions)
				npoly.Holes[i] = Ring(nseries.Move(deltaX, deltaY))
			}
		}
	}
	return npoly
}

func (poly *Poly) ContainsPoint(point Point) bool {
	if poly == nil || poly.Exterior == nil {
		return false
	}
	if !ringContainsPoint(poly.Exterior, point, true).hit {
		return false
	}
	contains := true
	for _, hole := range poly.Holes {
		if ringContainsPoint(hole, point, false).hit {
			contains = false
			break
		}
	}
	return contains
}

func (poly *Poly) IntersectsPoint(point Point) bool {
	if poly == nil {
		return false
	}
	return poly.ContainsPoint(point)
}

func (poly *Poly) ContainsRect(rect Rect) bool {
	if poly == nil {
		return false
	}
	// convert rect into a polygon
	return poly.ContainsPoly(&Poly{Exterior: rect})
}

func (poly *Poly) IntersectsRect(rect Rect) bool {
	if poly == nil {
		return false
	}
	// convert rect into a polygon
	return poly.IntersectsPoly(&Poly{Exterior: rect})
}

func (poly *Poly) ContainsLine(line *Line) bool {
	if poly == nil || poly.Exterior == nil || line == nil {
		return false
	}
	if !ringContainsLine(poly.Exterior, line, true) {
		return false
	}
	for _, polyHole := range poly.Holes {
		if ringIntersectsLine(polyHole, line, false) {
			return false
		}
	}
	return true
}

func (poly *Poly) IntersectsLine(line *Line) bool {
	if poly == nil || poly.Exterior == nil || line == nil {
		return false
	}
	if !ringIntersectsLine(poly.Exterior, line, true) {
		return false
	}
	for _, hole := range poly.Holes {
		if ringContainsLine(hole, line, false) {
			return false
		}
	}
	return true
}

func (poly *Poly) ContainsPoly(other *Poly) bool {
	if poly == nil || poly.Exterior == nil ||
		other == nil || other.Exterior == nil {
		return false
	}
	// 1) other exterior must be fully contained inside of the poly exterior.
	if !ringContainsRing(poly.Exterior, other.Exterior, true) {
		return false
	}
	// 2) ring cannot intersect poly holes
	contains := true
	for _, polyHole := range poly.Holes {
		if ringIntersectsRing(polyHole, other.Exterior, false) {
			contains = false
			// 3) unless the poly hole is contain inside of a other hole
			for _, otherHole := range other.Holes {
				if ringContainsRing(otherHole, polyHole, true) {
					contains = true
					// println(4)
					break
				}
			}
			if !contains {
				break
			}
		}
	}
	return contains
}

func (poly *Poly) IntersectsPoly(other *Poly) bool {
	if poly == nil || poly.Exterior == nil ||
		other == nil || other.Exterior == nil {
		return false
	}
	if !ringIntersectsRing(other.Exterior, poly.Exterior, true) {
		return false
	}
	for _, hole := range poly.Holes {
		if ringContainsRing(hole, other.Exterior, false) {
			return false
		}
	}
	for _, hole := range other.Holes {
		if ringContainsRing(hole, poly.Exterior, false) {
			return false
		}
	}
	return true
}
