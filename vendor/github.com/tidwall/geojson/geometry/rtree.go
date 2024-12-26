package geometry

import (
	"encoding/binary"
	"math"
)

const rDims = 2
const rMaxEntries = 16

type rRect struct {
	data     interface{}
	min, max [rDims]float64
}

type rNode struct {
	count int
	rects [rMaxEntries + 1]rRect
}

type rTree struct {
	height   int
	root     rRect
	count    int
	reinsert []rRect
}

func (r *rRect) expand(b *rRect) {
	for i := 0; i < rDims; i++ {
		if b.min[i] < r.min[i] {
			r.min[i] = b.min[i]
		}
		if b.max[i] > r.max[i] {
			r.max[i] = b.max[i]
		}
	}
}

// Insert inserts an item into the RTree
func (tr *rTree) Insert(min, max []float64, value interface{}) {
	var item rRect
	fit(min, max, value, &item)
	tr.insert(&item)
}

func (tr *rTree) insert(item *rRect) {
	if tr.root.data == nil {
		fit(item.min[:], item.max[:], new(rNode), &tr.root)
	}
	grown := tr.root.insert(item, tr.height)
	if grown {
		tr.root.expand(item)
	}
	if tr.root.data.(*rNode).count == rMaxEntries+1 {
		newRoot := new(rNode)
		tr.root.splitLargestAxisEdgeSnap(&newRoot.rects[1])
		newRoot.rects[0] = tr.root
		newRoot.count = 2
		tr.root.data = newRoot
		tr.root.recalc()
		tr.height++
	}
	tr.count++
}

func (r *rRect) chooseLeastEnlargement(b *rRect) int {
	j, jenlargement, jarea := -1, 0.0, 0.0
	n := r.data.(*rNode)
	for i := 0; i < n.count; i++ {
		// force inline
		area := n.rects[i].max[0] - n.rects[i].min[0]
		for j := 1; j < rDims; j++ {
			area *= n.rects[i].max[j] - n.rects[i].min[j]
		}
		var enlargement float64
		// force inline
		enlargedArea := 1.0
		for j := 0; j < len(n.rects[i].min); j++ {
			if b.max[j] > n.rects[i].max[j] {
				if b.min[j] < n.rects[i].min[j] {
					enlargedArea *= b.max[j] - b.min[j]
				} else {
					enlargedArea *= b.max[j] - n.rects[i].min[j]
				}
			} else {
				if b.min[j] < n.rects[i].min[j] {
					enlargedArea *= n.rects[i].max[j] - b.min[j]
				} else {
					enlargedArea *= n.rects[i].max[j] - n.rects[i].min[j]
				}
			}
		}
		enlargement = enlargedArea - area

		if j == -1 || enlargement < jenlargement {
			j, jenlargement, jarea = i, enlargement, area
		} else if enlargement == jenlargement {
			if area < jarea {
				j, jenlargement, jarea = i, enlargement, area
			}
		}
	}
	return j
}

func (r *rRect) recalc() {
	n := r.data.(*rNode)
	r.min = n.rects[0].min
	r.max = n.rects[0].max
	for i := 1; i < n.count; i++ {
		r.expand(&n.rects[i])
	}
}

// contains return struct when b is fully contained inside of n
func (r *rRect) contains(b *rRect) bool {
	for i := 0; i < rDims; i++ {
		if b.min[i] < r.min[i] || b.max[i] > r.max[i] {
			return false
		}
	}
	return true
}

func (r *rRect) largestAxis() (axis int, size float64) {
	j, jsz := 0, 0.0
	for i := 0; i < rDims; i++ {
		sz := r.max[i] - r.min[i]
		if i == 0 || sz > jsz {
			j, jsz = i, sz
		}
	}
	return j, jsz
}

func (r *rRect) splitLargestAxisEdgeSnap(right *rRect) {
	axis, _ := r.largestAxis()
	left := r
	leftNode := left.data.(*rNode)
	rightNode := new(rNode)
	right.data = rightNode

	var equals []rRect
	for i := 0; i < leftNode.count; i++ {
		minDist := leftNode.rects[i].min[axis] - left.min[axis]
		maxDist := left.max[axis] - leftNode.rects[i].max[axis]
		if minDist < maxDist {
			// stay left
		} else {
			if minDist > maxDist {
				// move to right
				rightNode.rects[rightNode.count] = leftNode.rects[i]
				rightNode.count++
			} else {
				// move to equals, at the end of the left array
				equals = append(equals, leftNode.rects[i])
			}
			leftNode.rects[i] = leftNode.rects[leftNode.count-1]
			leftNode.rects[leftNode.count-1].data = nil
			leftNode.count--
			i--
		}
	}
	for _, b := range equals {
		if leftNode.count < rightNode.count {
			leftNode.rects[leftNode.count] = b
			leftNode.count++
		} else {
			rightNode.rects[rightNode.count] = b
			rightNode.count++
		}
	}
	left.recalc()
	right.recalc()
}

func (r *rRect) insert(item *rRect, height int) (grown bool) {
	n := r.data.(*rNode)
	if height == 0 {
		n.rects[n.count] = *item
		n.count++
		grown = !r.contains(item)
		return grown
	}
	// choose subtree
	index := r.chooseLeastEnlargement(item)
	child := &n.rects[index]
	grown = child.insert(item, height-1)
	if grown {
		child.expand(item)
		grown = !r.contains(item)
	}
	if child.data.(*rNode).count == rMaxEntries+1 {
		child.splitLargestAxisEdgeSnap(&n.rects[n.count])
		n.count++
	}
	return grown
}

// fit an external item into a rect type
func fit(min, max []float64, value interface{}, target *rRect) {
	if max == nil {
		max = min
	}
	if len(min) != len(max) {
		panic("min/max dimension mismatch")
	}
	if len(min) != rDims {
		panic("invalid number of dimensions")
	}
	for i := 0; i < rDims; i++ {
		target.min[i] = min[i]
		target.max[i] = max[i]
	}
	target.data = value
}

func (r *rRect) intersects(b *rRect) bool {
	for i := 0; i < rDims; i++ {
		if b.min[i] > r.max[i] || b.max[i] < r.min[i] {
			return false
		}
	}
	return true
}

func (r *rRect) search(
	target *rRect, height int,
	iter func(min, max []float64, value interface{}) bool,
) bool {
	n := r.data.(*rNode)
	if height == 0 {
		for i := 0; i < n.count; i++ {
			if target.intersects(&n.rects[i]) {
				if !iter(n.rects[i].min[:], n.rects[i].max[:],
					n.rects[i].data) {
					return false
				}
			}
		}
	} else {
		for i := 0; i < n.count; i++ {
			if target.intersects(&n.rects[i]) {
				if !n.rects[i].search(target, height-1, iter) {
					return false
				}
			}
		}
	}
	return true
}

func (tr *rTree) search(
	target *rRect,
	iter func(min, max []float64, value interface{}) bool,
) {
	if tr.root.data == nil {
		return
	}
	if target.intersects(&tr.root) {
		tr.root.search(target, tr.height, iter)
	}
}

func (tr *rTree) Search(
	min, max []float64,
	iter func(min, max []float64, value interface{}) bool,
) {
	var target rRect
	fit(min, max, nil, &target)
	tr.search(&target, iter)
}

func appendFloat(dst []byte, num float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(num))
	return append(dst, buf[:]...)
}

func (tr *rTree) compress(dst []byte) []byte {
	if tr.root.data == nil {
		return dst
	}
	dst = append(dst, byte(tr.height))
	return tr.root.compress(dst, tr.height)
}

func (r *rRect) compress(dst []byte, height int) []byte {
	n := r.data.(*rNode)
	dst = appendFloat(dst, r.min[0])
	dst = appendFloat(dst, r.min[1])
	dst = appendFloat(dst, r.max[0])
	dst = appendFloat(dst, r.max[1])
	dst = append(dst, byte(n.count))
	if height == 0 {
		var ibytes byte = 1
		for i := 0; i < n.count; i++ {
			ibytes2 := numBytes(uint32(n.rects[i].data.(int)))
			if ibytes2 > ibytes {
				ibytes = ibytes2
			}
		}
		dst = append(dst, ibytes)
		for i := 0; i < n.count; i++ {
			dst = appendNum(dst, uint32(n.rects[i].data.(int)), ibytes)
		}
		return dst
	}
	mark := make([]int, n.count)
	for i := 0; i < n.count; i++ {
		mark[i] = len(dst)
		dst = append(dst, 0, 0, 0, 0)
	}
	for i := 0; i < n.count; i++ {
		binary.LittleEndian.PutUint32(dst[mark[i]:], uint32(len(dst)))
		dst = n.rects[i].compress(dst, height-1)
	}
	return dst
}

func rCompressSearch(
	data []byte,
	addr int,
	series *baseSeries,
	rect Rect,
	iter func(seg Segment, item int) bool,
) bool {
	if int(addr) == len(data) {
		return true
	}
	height := int(data[addr])
	addr++
	return rnCompressSearch(data, addr, series, rect, height, iter)
}

func rnCompressSearch(
	data []byte,
	addr int,
	series *baseSeries,
	rect Rect,
	height int,
	iter func(seg Segment, item int) bool,
) bool {
	var nrect Rect
	nrect.Min.X = math.Float64frombits(binary.LittleEndian.Uint64(data[addr:]))
	addr += 8
	nrect.Min.Y = math.Float64frombits(binary.LittleEndian.Uint64(data[addr:]))
	addr += 8
	nrect.Max.X = math.Float64frombits(binary.LittleEndian.Uint64(data[addr:]))
	addr += 8
	nrect.Max.Y = math.Float64frombits(binary.LittleEndian.Uint64(data[addr:]))
	addr += 8
	if !rect.IntersectsRect(nrect) {
		return true
	}
	count := int(data[addr])
	addr++
	if height == 0 {
		ibytes := data[addr]
		addr++
		for i := 0; i < count; i++ {
			item := int(readNum(data[addr:], ibytes))
			addr += int(ibytes)
			seg := series.SegmentAt(int(item))
			irect := seg.Rect()
			if irect.IntersectsRect(rect) {
				if !iter(seg, int(item)) {
					return false
				}
			}
		}
		return true
	}
	for i := 0; i < count; i++ {
		naddr := int(binary.LittleEndian.Uint32(data[addr:]))
		addr += 4
		if !rnCompressSearch(data, naddr, series, rect, height-1, iter) {
			return false
		}
	}
	return true
}
