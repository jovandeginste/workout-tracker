// Copyright 2021 Joshua J Baker. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package rtree

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/tidwall/geoindex/child"
)

// SAFTEY: The unsafe package is used, but with care.
// Using "unsafe" allows for one alloction per node and avoids having to use
// an interface{} type for child nodes; that may either be:
//   - *leafNode[N,T]
//   - *branchNode[N,T]
// This library makes it generally safe by guaranteeing that all references to
// nodes are simply to `*node[N,T]`, which is just the header struct for the
// leaf or branch representation. The difference between a leaf and a branch
// node is that a leaf has an array of item data of generic type T on tail of
// the struct, while a branch has an array of child node pointers on the tail.
// To access the child items `node[N,T].items()` is called; returning a slice,
// or nil if the node is a branch. To access the child nodes
// `node[N,T].children()` is called; returning a slice, or nil if the node is a
// leaf. The `items()` and `children()` methods check the `node[N,T].kind` to
// determine which kind of node it is, which is an enum of `none`, `leaf`, or
// `branch`. The only valid way to create a `*node[N,T]` is
// `RTreeGN[N,T].newNode(leaf bool)` which take a bool that indicates the new
// node kind is a `leaf` or `branch`.

const maxEntries = 64
const orderBranches = true
const orderLeaves = true

// copy-on-write atomic incrementer
var gcow uint64

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type RTreeGN[N numeric, T any] struct {
	icow  uint64
	count int
	rect  rect[N]
	root  *node[N, T]
	empty T
	qpool *sync.Pool
}

type rect[N numeric] struct {
	min [2]N
	max [2]N
}

func (r *rect[N]) expand(b *rect[N]) {
	if b.min[0] < r.min[0] {
		r.min[0] = b.min[0]
	}
	if b.max[0] > r.max[0] {
		r.max[0] = b.max[0]
	}
	if b.min[1] < r.min[1] {
		r.min[1] = b.min[1]
	}
	if b.max[1] > r.max[1] {
		r.max[1] = b.max[1]
	}
}

type kind int8

const (
	none kind = iota
	leaf
	branch
)

type node[N numeric, T any] struct {
	icow  uint64
	kind  kind
	count int16
	rects [maxEntries]rect[N]
}

func (n *node[N, T]) leaf() bool {
	return n.kind == leaf
}

type leafNode[N numeric, T any] struct {
	node[N, T]
	items [maxEntries]T
}

type branchNode[N numeric, T any] struct {
	node[N, T]
	children [maxEntries]*node[N, T]
}

func (n *node[N, T]) children() []*node[N, T] {
	if n.kind != branch {
		// not a branch
		return nil
	}
	return (*branchNode[N, T])(unsafe.Pointer(n)).children[:]
}

func (n *node[N, T]) items() []T {
	if n.kind != leaf {
		// not a leaf
		return nil
	}
	return (*leafNode[N, T])(unsafe.Pointer(n)).items[:]
}

func (tr *RTreeGN[N, T]) newNode(isleaf bool) *node[N, T] {
	if isleaf {
		n := &leafNode[N, T]{node: node[N, T]{icow: tr.icow, kind: leaf}}
		return (*node[N, T])(unsafe.Pointer(n))
	} else {
		n := &branchNode[N, T]{node: node[N, T]{icow: tr.icow, kind: branch}}
		return (*node[N, T])(unsafe.Pointer(n))
	}
}

func (n *node[N, T]) rect() rect[N] {
	rect := n.rects[0]
	for i := 1; i < int(n.count); i++ {
		rect.expand(&n.rects[i])
	}
	return rect
}

// Insert data into tree
func (tr *RTreeGN[N, T]) Insert(min, max [2]N, data T) {
	ir := rect[N]{min, max}
	if tr.root == nil {
		if tr.qpool == nil {
			tr.qpool = &sync.Pool{
				New: func() any { return &queue[N, T]{} },
			}
		}
		tr.root = tr.newNode(true)
		tr.rect = ir
	}
	tr.cow(&tr.root)
	split, grown := tr.nodeInsert(&tr.rect, tr.root, &ir, data)
	if split {
		left := tr.root
		right := tr.splitNode(tr.rect, left)
		tr.root = tr.newNode(false)
		tr.root.rects[0] = left.rect()
		tr.root.rects[1] = right.rect()
		tr.root.children()[0] = left
		tr.root.children()[1] = right
		tr.root.count = 2
		tr.Insert(min, max, data)
		if orderBranches {
			tr.root.sort()
		}
		return
	}
	if grown {
		tr.rect.expand(&ir)
		if orderBranches && !tr.root.leaf() {
			tr.root.sort()
		}
	}
	tr.count++
}

func (tr *RTreeGN[N, T]) splitNode(r rect[N], left *node[N, T],
) (right *node[N, T]) {
	return tr.splitNodeLargestAxisEdgeSnap(r, left)
}

func (n *node[N, T]) orderToRight(idx int) int {
	for idx < int(n.count)-1 && n.rects[idx+1].min[0] < n.rects[idx].min[0] {
		n.swap(idx+1, idx)
		idx++
	}
	return idx
}

func (n *node[N, T]) orderToLeft(idx int) int {
	for idx > 0 && n.rects[idx].min[0] < n.rects[idx-1].min[0] {
		n.swap(idx, idx-1)
		idx--
	}
	return idx
}

// This operation should not be inlined because it's expensive and rarely
// called outside of heavy copy-on-write situations. Marking it "noinline"
// allows for the parent cowLoad to be inlined.
// go:noinline
func (tr *RTreeGN[N, T]) copy(n *node[N, T]) *node[N, T] {
	n2 := tr.newNode(n.leaf())
	*n2 = *n
	if n2.leaf() {
		copy(n2.items()[:n.count], n.items()[:n.count])
	} else {
		copy(n2.children()[:n.count], n.children()[:n.count])
	}
	return n2
}

// cow ensures the provided node is not being shared with other R-trees.
// Performs a copy-on-write, if needed.
func (tr *RTreeGN[N, T]) cow(n **node[N, T]) {
	if (*n).icow != tr.icow {
		*n = tr.copy(*n)
	}
}

func (n *node[N, T]) rsearch(key N) int {
	rects := n.rects[:n.count]
	for i := 0; i < len(rects); i++ {
		if !(n.rects[i].min[0] < key) {
			return i
		}
	}
	return int(n.count)
}

func (tr *RTreeGN[N, T]) nodeInsert(nr *rect[N], n *node[N, T], ir *rect[N],
	data T,
) (split, grown bool) {
	if n.leaf() {
		if n.count == maxEntries {
			return true, false
		}
		items := n.items()
		index := int(n.count)
		if orderLeaves {
			index = n.rsearch(ir.min[0])
			copy(n.rects[index+1:int(n.count)+1], n.rects[index:int(n.count)])
			copy(items[index+1:int(n.count)+1], items[index:int(n.count)])
		}
		n.rects[index] = *ir
		items[index] = data
		n.count++
		grown = !nr.contains(ir)
		return false, grown
	}

	// choose a subtree
	rects := n.rects[:n.count]
	index := -1
	var narea N
	// take a quick look for any nodes that contain the rect
	for i := 0; i < len(rects); i++ {
		if rects[i].contains(ir) {
			area := rects[i].area()
			if index == -1 || area < narea {
				index = i
				narea = area
			}
		}
	}
	if index == -1 {
		index = n.chooseLeastEnlargement(ir)
	}

	children := n.children()
	tr.cow(&children[index])
	split, grown = tr.nodeInsert(&n.rects[index], children[index], ir, data)
	if split {
		if n.count == maxEntries {
			return true, false
		}
		// split the child node
		left := children[index]
		right := tr.splitNode(n.rects[index], left)
		n.rects[index] = left.rect()
		if orderBranches {
			copy(n.rects[index+2:int(n.count)+1],
				n.rects[index+1:int(n.count)])
			copy(children[index+2:int(n.count)+1],
				children[index+1:int(n.count)])
			n.rects[index+1] = right.rect()
			children[index+1] = right
			n.count++
			if n.rects[index].min[0] > n.rects[index+1].min[0] {
				n.swap(index+1, index)
			}
			index++
			_ = n.orderToRight(index)
		} else {
			n.rects[n.count] = right.rect()
			children[n.count] = right
			n.count++
		}
		return tr.nodeInsert(nr, n, ir, data)
	}
	if grown {
		// The child rectangle must expand to accomadate the new item.
		n.rects[index].expand(ir)
		if orderBranches {
			n.orderToLeft(index)
		}
		grown = !nr.contains(ir)
	}
	return false, grown
}

func (r *rect[N]) area() N {
	return (r.max[0] - r.min[0]) * (r.max[1] - r.min[1])
}

// contains return struct when b is fully contained inside of n
func (r *rect[N]) contains(b *rect[N]) bool {
	if b.min[0] < r.min[0] || b.max[0] > r.max[0] {
		return false
	}
	if b.min[1] < r.min[1] || b.max[1] > r.max[1] {
		return false
	}
	return true
}

// intersects returns true if both rects intersect each other.
func (r *rect[N]) intersects(b *rect[N]) bool {
	if b.min[0] > r.max[0] || b.max[0] < r.min[0] {
		return false
	}
	if b.min[1] > r.max[1] || b.max[1] < r.min[1] {
		return false
	}
	return true
}

func (n *node[N, T]) chooseLeastEnlargement(ir *rect[N]) (index int) {
	rects := n.rects[:int(n.count)]
	var j = -1
	var jenlargement N
	var jarea N
	for i := 0; i < len(rects); i++ {
		// calculate the enlarged area
		uarea := rects[i].unionedArea(ir)
		area := rects[i].area()
		enlargement := uarea - area
		if j == -1 || enlargement < jenlargement ||
			(!(enlargement > jenlargement) && area < jarea) {
			j, jenlargement, jarea = i, enlargement, area
		}
	}
	return j
}

func fmin[N numeric](a, b N) N {
	if a < b {
		return a
	}
	return b
}
func fmax[N numeric](a, b N) N {
	if a > b {
		return a
	}
	return b
}

// unionedArea returns the area of two rects expanded
func (r *rect[N]) unionedArea(b *rect[N]) N {
	return (fmax(r.max[0], b.max[0]) - fmin(r.min[0], b.min[0])) *
		(fmax(r.max[1], b.max[1]) - fmin(r.min[1], b.min[1]))
}

func (r rect[N]) largestAxis() (axis int) {
	if r.max[1]-r.min[1] > r.max[0]-r.min[0] {
		return 1
	}
	return 0
}

func (tr *RTreeGN[N, T]) splitNodeLargestAxisEdgeSnap(r rect[N], left *node[N, T],
) (right *node[N, T]) {
	axis := r.largestAxis()
	right = tr.newNode(left.leaf())
	for i := 0; i < int(left.count); i++ {
		minDist := left.rects[i].min[axis] - r.min[axis]
		maxDist := r.max[axis] - left.rects[i].max[axis]
		if minDist < maxDist {
			// stay left
		} else {
			// move to right
			tr.moveRectAtIndexInto(left, i, right)
			i--
		}
	}
	// Make sure that both left and right nodes have at least
	// two by moving items into underflowed nodes.
	if left.count < 2 {
		// reverse sort by min axis
		right.sortByAxis(axis, true, false)
		for left.count < 2 {
			tr.moveRectAtIndexInto(right, int(right.count)-1, left)
		}
	} else if right.count < 2 {
		// reverse sort by max axis
		left.sortByAxis(axis, true, true)
		for right.count < 2 {
			tr.moveRectAtIndexInto(left, int(left.count)-1, right)
		}
	}

	if (orderBranches && !right.leaf()) || (orderLeaves && right.leaf()) {
		// It's not uncommon that the nodes to be already ordered.
		if !right.issorted() {
			right.sort()
		}
		if !left.issorted() {
			left.sort()
		}
	}
	return right
}

func (tr *RTreeGN[N, T]) moveRectAtIndexInto(from *node[N, T], index int,
	into *node[N, T],
) {
	into.rects[into.count] = from.rects[index]
	from.rects[index] = from.rects[from.count-1]
	if from.leaf() {
		into.items()[into.count] = from.items()[index]
		from.items()[index] = from.items()[from.count-1]
		from.items()[from.count-1] = tr.empty
	} else {
		into.children()[into.count] = from.children()[index]
		from.children()[index] = from.children()[from.count-1]
		from.children()[from.count-1] = nil
	}
	from.count--
	into.count++
}

func (n *node[N, T]) search(target rect[N],
	iter func(min, max [2]N, data T) bool,
) bool {
	rects := n.rects[:n.count]
	if n.leaf() {
		items := n.items()
		for i := 0; i < len(rects); i++ {
			if rects[i].intersects(&target) {
				if !iter(rects[i].min, rects[i].max, items[i]) {
					return false
				}
			}
		}
		return true
	}
	children := n.children()
	for i := 0; i < len(rects); i++ {
		if target.intersects(&rects[i]) {
			if !children[i].search(target, iter) {
				return false
			}
		}
	}
	return true
}

// Len returns the number of items in tree
func (tr *RTreeGN[N, T]) Len() int {
	return tr.count
}

// Search for items in tree that intersect the provided rectangle
func (tr *RTreeGN[N, T]) Search(min, max [2]N,
	iter func(min, max [2]N, data T) bool,
) {
	target := rect[N]{min, max}
	if tr.root == nil {
		return
	}
	if target.intersects(&tr.rect) {
		tr.root.search(target, iter)
	}
}

// Scane all items in the tree
func (tr *RTreeGN[N, T]) Scan(iter func(min, max [2]N, data T) bool) {
	if tr.root != nil {
		tr.root.scan(iter)
	}
}

func (n *node[N, T]) scan(iter func(min, max [2]N, data T) bool) bool {
	if n.leaf() {
		for i := 0; i < int(n.count); i++ {
			if !iter(n.rects[i].min, n.rects[i].max, n.items()[i]) {
				return false
			}
		}
	} else {
		for i := 0; i < int(n.count); i++ {
			if !n.children()[i].scan(iter) {
				return false
			}
		}
	}
	return true
}

// Copy the tree.
// This is a copy-on-write operation and is very fast because it only performs
// a shadowed copy.
func (tr *RTreeGN[N, T]) Copy() *RTreeGN[N, T] {
	tr2 := new(RTreeGN[N, T])
	*tr2 = *tr
	tr.icow = atomic.AddUint64(&gcow, 1)
	tr2.icow = atomic.AddUint64(&gcow, 1)
	return tr2
}

// swap two rectanlges
func (n *node[N, T]) swap(i, j int) {
	n.rects[i], n.rects[j] = n.rects[j], n.rects[i]
	if n.leaf() {
		n.items()[i], n.items()[j] = n.items()[j], n.items()[i]
	} else {
		n.children()[i], n.children()[j] = n.children()[j], n.children()[i]
	}
}

func (n *node[N, T]) sortByAxis(axis int, rev, max bool) {
	n.qsort(0, int(n.count), axis, rev, max)
}

func (n *node[N, T]) sort() {
	n.qsort(0, int(n.count), 0, false, false)
}

func (n *node[N, T]) issorted() bool {
	rects := n.rects[:n.count]
	for i := 1; i < len(rects); i++ {
		if rects[i].min[0] < rects[i-1].min[0] {
			return false
		}
	}
	return true
}

func (n *node[N, T]) qsort(s, e int, axis int, rev, max bool) {
	nrects := e - s
	if nrects < 2 {
		return
	}
	left, right := 0, nrects-1
	pivot := nrects / 2 // rand and mod not worth it
	n.swap(s+pivot, s+right)
	rects := n.rects[s:e]
	if !rev {
		if !max {
			for i := 0; i < len(rects); i++ {
				if rects[i].min[axis] < rects[right].min[axis] {
					n.swap(s+i, s+left)
					left++
				}
			}
		} else {
			for i := 0; i < len(rects); i++ {
				if rects[i].max[axis] < rects[right].max[axis] {
					n.swap(s+i, s+left)
					left++
				}
			}
		}
	} else {
		if !max {
			for i := 0; i < len(rects); i++ {
				if rects[right].min[axis] < rects[i].min[axis] {
					n.swap(s+i, s+left)
					left++
				}
			}
		} else {
			for i := 0; i < len(rects); i++ {
				if rects[right].max[axis] < rects[i].max[axis] {
					n.swap(s+i, s+left)
					left++
				}
			}
		}
	}
	n.swap(s+left, s+right)
	n.qsort(s, s+left, axis, rev, max)
	n.qsort(s+left+1, e, axis, rev, max)
}

// Delete data from tree
func (tr *RTreeGN[N, T]) Delete(min, max [2]N, data T) {
	tr.delete(min, max, data)
}

func (tr *RTreeGN[N, T]) delete(min, max [2]N, data T) bool {
	ir := rect[N]{min, max}
	if tr.root == nil || !tr.rect.contains(&ir) {
		return false
	}
	var reinsert []*node[N, T]
	tr.cow(&tr.root)
	removed, _ := tr.nodeDelete(&tr.rect, tr.root, &ir, data, &reinsert)
	if !removed {
		return false
	}
	tr.count--
	if len(reinsert) > 0 {
		for _, n := range reinsert {
			tr.count -= n.deepCount()
		}
	}
	if tr.count == 0 {
		tr.root = nil
		tr.rect.min = [2]N{0, 0}
		tr.rect.max = [2]N{0, 0}
	} else {
		for !tr.root.leaf() && tr.root.count == 1 {
			tr.root = tr.root.children()[0]
		}
	}
	if len(reinsert) > 0 {
		for i := range reinsert {
			tr.nodeReinsert(reinsert[i])
		}
	}
	return true
}

func compare[T any](a, b T) bool {
	return (interface{})(a) == (interface{})(b)
}

func (tr *RTreeGN[N, T]) nodeDelete(nr *rect[N], n *node[N, T], ir *rect[N], data T,
	reinsert *[]*node[N, T],
) (removed, shrunk bool) {
	rects := n.rects[:n.count]
	if n.leaf() {
		items := n.items()
		for i := 0; i < len(rects); i++ {
			if ir.contains(&rects[i]) && compare(items[i], data) {
				// found the target item to delete
				if orderLeaves {
					copy(n.rects[i:n.count], n.rects[i+1:n.count])
					copy(items[i:n.count], items[i+1:n.count])
				} else {
					n.rects[i] = n.rects[n.count-1]
					items[i] = items[n.count-1]
				}
				items[len(rects)-1] = tr.empty
				n.count--
				shrunk = ir.onedge(nr)
				if shrunk {
					*nr = n.rect()
				}
				return true, shrunk
			}
		}
		return false, false
	}
	children := n.children()
	for i := 0; i < len(rects); i++ {
		if !rects[i].contains(ir) {
			continue
		}
		crect := rects[i]
		tr.cow(&children[i])
		removed, shrunk = tr.nodeDelete(&rects[i], children[i], ir, data,
			reinsert)
		if !removed {
			continue
		}
		if children[i].count == 0 {
			*reinsert = append(*reinsert, children[i])
			if orderBranches {
				copy(n.rects[i:n.count], n.rects[i+1:n.count])
				copy(children[i:n.count], children[i+1:n.count])
			} else {
				n.rects[i] = n.rects[n.count-1]
				children[i] = children[n.count-1]
			}
			children[n.count-1] = nil
			n.count--
			*nr = n.rect()
			return true, true
		}
		if shrunk {
			shrunk = !rects[i].equals(&crect)
			if shrunk {
				*nr = n.rect()
			}
			if orderBranches {
				_ = n.orderToRight(i)
			}
		}
		return true, shrunk
	}
	return false, false
}

func (r *rect[N]) equals(b *rect[N]) bool {
	return !(r.min[0] < b.min[0] || r.min[0] > b.min[0] ||
		r.min[1] < b.min[1] || r.min[1] > b.min[1] ||
		r.max[0] < b.max[0] || r.max[0] > b.max[0] ||
		r.max[1] < b.max[1] || r.max[1] > b.max[1])
}

func (n *node[N, T]) deepCount() int {
	if n.leaf() {
		return int(n.count)
	}
	var count int
	children := n.children()[:n.count]
	for i := 0; i < len(children); i++ {
		count += children[i].deepCount()
	}
	return count
}

func (tr *RTreeGN[N, T]) nodeReinsert(n *node[N, T]) {
	if n.leaf() {
		rects := n.rects[:n.count]
		items := n.items()[:n.count]
		for i := range rects {
			tr.Insert(rects[i].min, rects[i].max, items[i])
		}
	} else {
		children := n.children()[:n.count]
		for i := 0; i < len(children); i++ {
			tr.nodeReinsert(children[i])
		}
	}
}

// onedge returns true when r is on the edge of b
func (r *rect[N]) onedge(b *rect[N]) bool {
	return !(r.min[0] > b.min[0] && r.min[1] > b.min[1] &&
		r.max[0] < b.max[0] && r.max[1] < b.max[1])
}

// Replace an item.
// If the old item does not exist then the new item is not inserted.
func (tr *RTreeGN[N, T]) Replace(
	oldMin, oldMax [2]N, oldData T,
	newMin, newMax [2]N, newData T,
) {
	if tr.delete(oldMin, oldMax, oldData) {
		tr.Insert(newMin, newMax, newData)
	}
}

// Bounds returns the minimum bounding rect
func (tr *RTreeGN[N, T]) Bounds() (min, max [2]N) {
	return tr.rect.min, tr.rect.max
}

func (tr *RTreeGN[N, T]) LeftMost() (min, max [2]N, data T) {
	if tr.root == nil {
		return
	}
	return tr.root.minist(0)
}
func (tr *RTreeGN[N, T]) BottomMost() (min, max [2]N, data T) {
	if tr.root == nil {
		return
	}
	return tr.root.minist(1)
}
func (tr *RTreeGN[N, T]) RightMost() (min, max [2]N, data T) {
	if tr.root == nil {
		return
	}
	return tr.root.maxist(0)
}

func (tr *RTreeGN[N, T]) TopMost() (min, max [2]N, data T) {
	if tr.root == nil {
		return
	}
	return tr.root.maxist(1)
}

func (n *node[N, T]) minist(dim int) (min, max [2]N, data T) {
	var j int
	var m N
	for i, r := range n.rects[:n.count] {
		if i == 0 || r.min[dim] < m {
			j, m = i, r.min[dim]
		}
	}
	if n.leaf() {
		return n.rects[j].min, n.rects[j].max, n.items()[j]
	}
	return n.children()[j].minist(dim)
}

func (n *node[N, T]) maxist(dim int) (min, max [2]N, data T) {
	var j int
	var m N
	for i, r := range n.rects[:n.count] {
		if i == 0 || r.max[dim] > m {
			j, m = i, r.max[dim]
		}
	}
	if n.leaf() {
		return n.rects[j].min, n.rects[j].max, n.items()[j]
	}
	return n.children()[j].maxist(dim)
}

// Nearby performs a kNN-type operation on the index.
// It's expected that the caller provides its own the `dist` function, which
// is used to calculate a distance to rectangles and data.
// The `iter` function will return all items from the smallest distance to the
// largest distance.
//
// BoxDist is included with this package for simple box-distance
// calculations. For example, say you want to return the closest items to
// Point(10 20):
//
//	tr.Nearby(
//		rtree.BoxDist([2]float64{10, 20}, [2]float64{10, 20}, nil),
//		func(min, max [2]float64, data int, dist float64) bool {
//			return true
//		},
//	)
func (tr *RTreeGN[N, T]) Nearby(
	dist func(min, max [2]N, data T, item bool) float64,
	iter func(min, max [2]N, data T, dist float64) bool,
) {
	if tr.root == nil {
		return
	}
	q := tr.qpool.Get().(*queue[N, T])
	defer func() {
		*q = (*q)[:0]
		tr.qpool.Put(q)
	}()

	q.push(qnode[N, T]{
		dist: 0,
		rect: tr.rect,
		node: tr.root,
	})
	for {
		qn, ok := q.pop()
		if !ok {
			return
		}
		if qn.node == nil {
			if !iter(qn.rect.min, qn.rect.max, qn.data, qn.dist) {
				return
			}
		} else {
			rects := qn.node.rects[:qn.node.count]
			if qn.node.leaf() {
				items := qn.node.items()[:qn.node.count]
				for i := 0; i < len(items); i++ {
					q.push(qnode[N, T]{
						dist: dist(rects[i].min, rects[i].max, items[i], true),
						rect: rects[i],
						data: items[i],
					})
				}
			} else {
				children := qn.node.children()[:qn.node.count]
				for i := 0; i < len(children); i++ {
					q.push(qnode[N, T]{
						dist: dist(rects[i].min, rects[i].max, tr.empty, false),
						rect: rects[i],
						node: children[i],
					})
				}
			}
		}
	}
}

type qnode[N numeric, T any] struct {
	dist float64     // distance to
	rect rect[N]     // item or node rect
	data T           // item data (or empty for node)
	node *node[N, T] // node (or nil for leaf data)
}

type queue[N numeric, T any] []qnode[N, T]

func (q *queue[N, T]) push(node qnode[N, T]) {
	*q = append(*q, node)
	nodes := *q
	i := len(nodes) - 1
	parent := (i - 1) / 2
	for ; i != 0 && nodes[parent].dist > nodes[i].dist; parent = (i - 1) / 2 {
		nodes[parent], nodes[i] = nodes[i], nodes[parent]
		i = parent
	}
}

func (q *queue[N, T]) pop() (qnode[N, T], bool) {
	nodes := *q
	if len(nodes) == 0 {
		return qnode[N, T]{}, false
	}
	var n qnode[N, T]
	n, nodes[0] = nodes[0], nodes[len(*q)-1]
	nodes = nodes[:len(nodes)-1]
	*q = nodes
	i := 0
	for {
		smallest := i
		left := i*2 + 1
		right := i*2 + 2
		if left < len(nodes) && nodes[left].dist <= nodes[smallest].dist {
			smallest = left
		}
		if right < len(nodes) && nodes[right].dist <= nodes[smallest].dist {
			smallest = right
		}
		if smallest == i {
			break
		}
		nodes[smallest], nodes[i] = nodes[i], nodes[smallest]
		i = smallest
	}
	return n, true
}

// BoxDist performs simple box-distance algorithm on rectangles.
// This is the default algorithm for Nearby.
func BoxDist[N numeric, T any](targetMin, targetMax [2]N,
	itemDist func(min, max [2]N, data T) N,
) (dist func(min, max [2]N, data T, item bool) N) {
	targ := rect[N]{targetMin, targetMax}
	return func(min, max [2]N, data T, item bool) (dist N) {
		if item && itemDist != nil {
			return itemDist(min, max, data)
		}
		return targ.boxDist(&rect[N]{min, max})
	}
}

func (r *rect[N]) boxDist(b *rect[N]) N {
	var dist N
	squared := fmax(r.min[0], b.min[0]) - fmin(r.max[0], b.max[0])
	if squared > 0 {
		dist += squared * squared
	}
	squared = fmax(r.min[1], b.min[1]) - fmin(r.max[1], b.max[1])
	if squared > 0 {
		dist += squared * squared
	}
	return dist
}

// Clear will delete all items.
func (tr *RTreeGN[N, T]) Clear() {
	tr.count = 0
	tr.rect = rect[N]{}
	tr.root = nil
}

////////////////////////////////////////////////////////////////////////////////
// Inherited wrapped types
////////////////////////////////////////////////////////////////////////////////

type RTreeG[T any] struct {
	base RTreeGN[float64, T]
}

// Insert data into tree
func (tr *RTreeG[T]) Insert(min, max [2]float64, data T) {
	tr.base.Insert(min, max, data)
}

// Len returns the number of items in tree
func (tr *RTreeG[T]) Len() int {
	return tr.base.Len()
}

// Search for items in tree that intersect the provided rectangle
func (tr *RTreeG[T]) Search(min, max [2]float64,
	iter func(min, max [2]float64, data T) bool,
) {
	tr.base.Search(min, max, iter)
}

// Scan all items in the tree
func (tr *RTreeG[T]) Scan(iter func(min, max [2]float64, data T) bool) {
	tr.base.Scan(iter)
}

// Copy the tree.
// This is a copy-on-write operation and is very fast because it only performs
// a shadowed copy.
func (tr *RTreeG[T]) Copy() *RTreeG[T] {
	return &RTreeG[T]{*tr.base.Copy()}
}

// Delete data from tree
func (tr *RTreeG[T]) Delete(min, max [2]float64, data T) {
	tr.base.Delete(min, max, data)
}

// Replace an item.
// If the old item does not exist then the new item is not inserted.
func (tr *RTreeG[T]) Replace(
	oldMin, oldMax [2]float64, oldData T,
	newMin, newMax [2]float64, newData T,
) {
	tr.base.Replace(
		oldMin, oldMax, oldData,
		newMin, newMax, newData,
	)
}

// Bounds returns the minimum bounding rect
func (tr *RTreeG[T]) Bounds() (min, max [2]float64) {
	return tr.base.Bounds()
}

// children is a utility function that returns all children for parent node.
// If parent node is nil then the root nodes should be returned. The min, max,
// data, and items slices all must have the same lengths. And, each element
// from all slices must be associated. Returns true for `items` when the the
// item at the leaf level. The reuse buffers are empty length slices that can
// optionally be used to avoid extra allocations.
func (tr *RTreeG[T]) children(parent interface{}, reuse []child.Child,
) (children []child.Child) {
	children = reuse
	if parent == nil {
		if tr.Len() > 0 {
			// fill with the root
			children = append(children, child.Child{
				Min:  tr.base.rect.min,
				Max:  tr.base.rect.max,
				Data: tr.base.root,
				Item: false,
			})
		}
	} else {
		// fill with child items
		n := parent.(*node[float64, T])
		for i := 0; i < int(n.count); i++ {
			c := child.Child{
				Min: n.rects[i].min, Max: n.rects[i].max, Item: n.leaf(),
			}
			if c.Item {
				c.Data = n.items()[i]
			} else {
				c.Data = n.children()[i]
			}
			children = append(children, c)
		}
	}
	return children
}

// Nearby performs a kNN-type operation on the index.
// It's expected that the caller provides its own the `dist` function, which
// is used to calculate a distance to rectangles and data.
// The `iter` function will return all items from the smallest distance to the
// largest distance.
//
// BoxDist is included with this package for simple box-distance
// calculations. For example, say you want to return the closest items to
// Point(10 20):
//
//	tr.Nearby(
//		rtree.BoxDist([2]float64{10, 20}, [2]float64{10, 20}, nil),
//		func(min, max [2]float64, data int, dist float64) bool {
//			return true
//		},
//	)
func (tr *RTreeG[T]) Nearby(
	dist func(min, max [2]float64, data T, item bool) float64,
	iter func(min, max [2]float64, data T, dist float64) bool,
) {
	tr.base.Nearby(dist, iter)
}

// Clear will delete all items.
func (tr *RTreeG[T]) Clear() {
	tr.base.Clear()
}

// Generic RTree
// Deprecated: use RTreeG
type Generic[T any] struct {
	RTreeG[T]
}

func (tr *Generic[T]) Copy() *Generic[T] {
	return &Generic[T]{*tr.RTreeG.Copy()}
}

type RTree struct {
	base RTreeG[any]
}

// Insert an item into the structure
func (tr *RTree) Insert(min, max [2]float64, data interface{}) {
	tr.base.Insert(min, max, data)
}

// Delete an item from the structure
func (tr *RTree) Delete(min, max [2]float64, data interface{}) {
	tr.base.Delete(min, max, data)
}

// Replace an item in the structure. This is effectively just a Delete
// followed by an Insert. But for some structures it may be possible to
// optimize the operation to avoid multiple passes
func (tr *RTree) Replace(
	oldMin, oldMax [2]float64, oldData interface{},
	newMin, newMax [2]float64, newData interface{},
) {
	tr.base.Replace(
		oldMin, oldMax, oldData,
		newMin, newMax, newData,
	)
}

// Search the structure for items that intersects the rect param
func (tr *RTree) Search(
	min, max [2]float64,
	iter func(min, max [2]float64, data interface{}) bool,
) {
	tr.base.Search(min, max, iter)
}

// Scan iterates through all data in tree in no specified order.
func (tr *RTree) Scan(iter func(min, max [2]float64, data interface{}) bool) {
	tr.base.Scan(iter)
}

// Len returns the number of items in tree
func (tr *RTree) Len() int {
	return tr.base.Len()
}

// Bounds returns the minimum bounding box
func (tr *RTree) Bounds() (min, max [2]float64) {
	return tr.base.Bounds()
}

// Children returns all children for parent node. If parent node is nil
// then the root nodes should be returned.
// The reuse buffer is an empty length slice that can optionally be used
// to avoid extra allocations.
func (tr *RTree) Children(parent interface{}, reuse []child.Child) (children []child.Child) {
	return tr.base.children(parent, reuse)
}

// Nearby performs a kNN-type operation on the index.
// It's expected that the caller provides its own the `dist` function, which
// is used to calculate a distance to rectangles and data.
// The `iter` function will return all items from the smallest distance to the
// largest distance.
//
// BoxDist is included with this package for simple box-distance
// calculations. For example, say you want to return the closest items to
// Point(10 20):
//
//	tr.Nearby(
//		rtree.BoxDist([2]float64{10, 20}, [2]float64{10, 20}, nil),
//		func(min, max [2]float64, data int, dist float64) bool {
//			return true
//		},
//	)
func (tr *RTree) Nearby(
	algo func(min, max [2]float64, data interface{}, item bool) (dist float64),
	iter func(min, max [2]float64, data interface{}, dist float64) bool,
) {
	tr.base.Nearby(algo, iter)
}

// Copy the tree.
// This is a copy-on-write operation and is very fast because it only performs
// a shadowed copy.
func (tr *RTree) Copy() *RTree {
	return &RTree{base: *tr.base.Copy()}
}

// Clear will delete all items.
func (tr *RTree) Clear() {
	tr.base.Clear()
}
