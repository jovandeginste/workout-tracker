package geometry

import (
	"encoding/binary"
)

const qMaxItems = 32
const qMaxDepth = 16

type qNode struct {
	split bool
	items []uint32
	quads [4]*qNode
}

func (n *qNode) insert(series *baseSeries, bounds, rect Rect, item, depth int) {
	if depth == qMaxDepth {
		// limit depth and insert now
		n.items = append(n.items, uint32(item))
	} else if n.split {
		// qnode is split so try to insert into a quad
		q := n.chooseQuad(bounds, rect)
		if q == -1 {
			// insert into overflow
			n.items = append(n.items, uint32(item))
		} else {
			// insert into quad
			qbounds := quadBounds(bounds, q)
			if n.quads[q] == nil {
				n.quads[q] = new(qNode)
			}
			n.quads[q].insert(series, qbounds, rect, item, depth+1)
		}
	} else if len(n.items) == qMaxItems {
		// split qnode, keep current items in place
		var nitems []uint32
		for i := 0; i < len(n.items); i++ {
			iitem := n.items[i]
			irect := series.SegmentAt(int(iitem)).Rect()
			q := n.chooseQuad(bounds, irect)
			if q == -1 {
				nitems = append(nitems, iitem)
			} else {
				qbounds := quadBounds(bounds, q)
				if n.quads[q] == nil {
					n.quads[q] = new(qNode)
				}
				n.quads[q].insert(series, qbounds, irect, int(iitem), depth+1)
			}
		}
		n.items = nitems
		n.split = true
		n.insert(series, bounds, rect, item, depth)
	} else {
		n.items = append(n.items, uint32(item))
	}
}

func (n *qNode) chooseQuad(bounds, rect Rect) int {
	mid := Point{
		X: (bounds.Min.X + bounds.Max.X) / 2,
		Y: (bounds.Min.Y + bounds.Max.Y) / 2,
	}
	if rect.Max.X < mid.X {
		if rect.Max.Y < mid.Y {
			return 2
		}
		if rect.Min.Y < mid.Y {
			return -1
		}
		return 0
	}
	if rect.Min.X < mid.X {
		return -1
	}
	if rect.Max.Y < mid.Y {
		return 3
	}
	if rect.Min.Y < mid.Y {
		return -1
	}
	return 1
}

func quadBounds(bounds Rect, q int) (qbounds Rect) {
	switch q {
	case 0:
		qbounds.Min.X = bounds.Min.X
		qbounds.Min.Y = (bounds.Min.Y + bounds.Max.Y) / 2
		qbounds.Max.X = (bounds.Min.X + bounds.Max.X) / 2
		qbounds.Max.Y = bounds.Max.Y
	case 1:
		qbounds.Min.X = (bounds.Min.X + bounds.Max.X) / 2
		qbounds.Min.Y = (bounds.Min.Y + bounds.Max.Y) / 2
		qbounds.Max.X = bounds.Max.X
		qbounds.Max.Y = bounds.Max.Y
	case 2:
		qbounds.Min.X = bounds.Min.X
		qbounds.Min.Y = bounds.Min.Y
		qbounds.Max.X = (bounds.Min.X + bounds.Max.X) / 2
		qbounds.Max.Y = (bounds.Min.Y + bounds.Max.Y) / 2
	case 3:
		qbounds.Min.X = (bounds.Min.X + bounds.Max.X) / 2
		qbounds.Min.Y = bounds.Min.Y
		qbounds.Max.X = bounds.Max.X
		qbounds.Max.Y = (bounds.Min.Y + bounds.Max.Y) / 2
	}
	return
}

func (n *qNode) search(
	series *baseSeries,
	bounds, rect Rect,
	iter func(seg Segment, item int) bool,
) bool {
	for _, item := range n.items {
		seg := series.SegmentAt(int(item))
		irect := seg.Rect()
		if irect.IntersectsRect(rect) {
			if !iter(seg, int(item)) {
				return false
			}
		}
	}
	if n.split {
		for q := 0; q < 4; q++ {
			if n.quads[q] != nil {
				qbounds := quadBounds(bounds, q)
				if qbounds.IntersectsRect(rect) {
					if !n.quads[q].search(series, qbounds, rect, iter) {
						return false
					}
				}
			}
		}
	}
	return true
}
func numBytes(n uint32) byte {
	if n <= 0xFF {
		return 1
	}
	if n <= 0xFFFF {
		return 2
	}
	return 4
}

func appendNum(dst []byte, num uint32, ibytes byte) []byte {
	switch ibytes {
	case 1:
		dst = append(dst, byte(num))
	case 2:
		dst = append(dst, 0, 0)
		binary.LittleEndian.PutUint16(dst[len(dst)-2:], uint16(num))
	default:
		dst = append(dst, 0, 0, 0, 0)
		binary.LittleEndian.PutUint32(dst[len(dst)-4:], uint32(num))
	}
	return dst
}

func readNum(data []byte, ibytes byte) uint32 {
	switch ibytes {
	case 1:
		return uint32(data[0])
	case 2:
		return uint32(binary.LittleEndian.Uint16(data))
	default:
		return binary.LittleEndian.Uint32(data)
	}
}
func (n *qNode) compress(dst []byte, bounds Rect) []byte {
	ibytes := numBytes(uint32(len(n.items)))
	for i := 0; i < len(n.items); i++ {
		ibytes2 := numBytes(n.items[i])
		if ibytes2 > ibytes {
			ibytes = ibytes2
		}
	}
	dst = append(dst, ibytes)
	dst = appendNum(dst, uint32(len(n.items)), ibytes)
	for i := 0; i < len(n.items); i++ {
		dst = appendNum(dst, n.items[i], ibytes)
	}
	if !n.split {
		dst = append(dst, 0)
		return dst
	}
	// store the quads
	dst = append(dst, 1)
	// first make the address space
	var mark [4]int
	for q := 0; q < 4; q++ {
		if n.quads[q] == nil {
			// no quad, no address
			dst = append(dst, 0)
		} else {
			// yes quad, plus addres
			dst = append(dst, 1)
			mark[q] = len(dst)
			dst = append(dst, 0, 0, 0, 0)
		}
	}
	// next add each quad
	for q := 0; q < 4; q++ {
		if n.quads[q] != nil {
			binary.LittleEndian.PutUint32(dst[mark[q]:], uint32(len(dst)))
			dst = n.quads[q].compress(dst, quadBounds(bounds, q))
		}
	}
	return dst
}

func qCompressSearch(
	data []byte,
	addr int,
	series *baseSeries,
	bounds, rect Rect,
	iter func(seg Segment, item int) bool,
) bool {
	ibytes := data[addr]
	addr++
	nItems := int(readNum(data[addr:], ibytes))
	addr += int(ibytes)
	for i := 0; i < nItems; i++ {
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
	split := data[addr] == 1
	addr++
	if split {
		for q := 0; q < 4; q++ {
			use := data[addr] == 1
			addr++
			if !use {
				continue
			}
			naddr := int(binary.LittleEndian.Uint32(data[addr:]))
			addr += 4
			qbounds := quadBounds(bounds, q)
			if qbounds.IntersectsRect(rect) {
				if !qCompressSearch(data, naddr, series, qbounds, rect, iter) {
					return false
				}
			}
		}
	}
	return true
}
