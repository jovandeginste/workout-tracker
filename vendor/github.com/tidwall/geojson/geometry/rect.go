// Copyright 2018 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package geometry

type Rect struct {
	Min, Max Point
}

func (rect Rect) Move(deltaX, deltaY float64) Rect {
	return Rect{
		Min: Point{X: rect.Min.X + deltaX, Y: rect.Min.Y + deltaY},
		Max: Point{X: rect.Max.X + deltaX, Y: rect.Max.Y + deltaY},
	}
}

func (rect Rect) Index() interface{} {
	return nil
}

func (rect Rect) Clockwise() bool {
	return false
}

func (rect Rect) Center() Point {
	return Point{(rect.Max.X + rect.Min.X) / 2, (rect.Max.Y + rect.Min.Y) / 2}
}

func (rect Rect) Area() float64 {
	return (rect.Max.X - rect.Min.X) * (rect.Max.Y - rect.Min.Y)
}

func (rect Rect) NumPoints() int {
	return 5
}

func (rect Rect) NumSegments() int {
	return 4
}

func (rect Rect) PointAt(index int) Point {
	switch index {
	default:
		return []Point{}[0]
	case 0:
		return Point{rect.Min.X, rect.Min.Y}
	case 1:
		return Point{rect.Max.X, rect.Min.Y}
	case 2:
		return Point{rect.Max.X, rect.Max.Y}
	case 3:
		return Point{rect.Min.X, rect.Max.Y}
	case 4:
		return Point{rect.Min.X, rect.Min.Y}
	}
}

func (rect Rect) SegmentAt(index int) Segment {
	switch index {
	default:
		return []Segment{}[0]
	case 0:
		return Segment{
			Point{rect.Min.X, rect.Min.Y},
			Point{rect.Max.X, rect.Min.Y},
		}
	case 1:
		return Segment{
			Point{rect.Max.X, rect.Min.Y},
			Point{rect.Max.X, rect.Max.Y},
		}
	case 2:
		return Segment{
			Point{rect.Max.X, rect.Max.Y},
			Point{rect.Min.X, rect.Max.Y},
		}
	case 3:
		return Segment{
			Point{rect.Min.X, rect.Max.Y},
			Point{rect.Min.X, rect.Min.Y},
		}
	}
}

func (rect Rect) Search(target Rect, iter func(seg Segment, idx int) bool) {
	var idx int
	rectNumSegments := rect.NumSegments()
	for i := 0; i < rectNumSegments; i++ {
		seg := rect.SegmentAt(i)
		if seg.Rect().IntersectsRect(target) {
			if !iter(seg, idx) {
				break
			}
		}
		idx++
	}
}

func (rect Rect) Empty() bool {
	return false
}

func (rect Rect) Valid() bool {
	return rect.Min.Valid() && rect.Max.Valid()
}

func (rect Rect) Rect() Rect {
	return rect
}

func (rect Rect) Convex() bool {
	return true
}

func (rect Rect) ContainsPoint(point Point) bool {
	return point.X >= rect.Min.X && point.X <= rect.Max.X &&
		point.Y >= rect.Min.Y && point.Y <= rect.Max.Y
}

func (rect Rect) IntersectsPoint(point Point) bool {
	return rect.ContainsPoint(point)
}

func (rect Rect) ContainsRect(other Rect) bool {
	if other.Min.X < rect.Min.X || other.Max.X > rect.Max.X {
		return false
	}
	if other.Min.Y < rect.Min.Y || other.Max.Y > rect.Max.Y {
		return false
	}
	return true
}

func (rect Rect) IntersectsRect(other Rect) bool {
	if rect.Min.Y > other.Max.Y || rect.Max.Y < other.Min.Y {
		return false
	}
	if rect.Min.X > other.Max.X || rect.Max.X < other.Min.X {
		return false
	}
	return true
}

func (rect Rect) ContainsLine(line *Line) bool {
	if line == nil {
		return false
	}
	return !line.Empty() && rect.ContainsRect(line.Rect())
}

func (rect Rect) IntersectsLine(line *Line) bool {
	if line == nil {
		return false
	}
	return ringIntersectsLine(rect, line, true)
}

func (rect Rect) ContainsPoly(poly *Poly) bool {
	if poly == nil {
		return false
	}
	return !poly.Empty() && rect.ContainsRect(poly.Rect())
}

func (rect Rect) IntersectsPoly(poly *Poly) bool {
	if poly == nil {
		return false
	}
	return poly.IntersectsRect(rect)
}
