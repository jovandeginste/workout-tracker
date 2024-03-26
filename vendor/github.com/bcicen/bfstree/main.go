package bfstree

import (
	"fmt"
	"strings"
)

var (
	Debug = false // enable debug message printing
)

func dbg(format string, a ...interface{}) {
	if Debug {
		fmt.Printf(format, a...)
	}
}

type Edge interface {
	From() string
	To() string
}

type BFSTree struct {
	edges []Edge
}

func New(edges ...Edge) *BFSTree { return &BFSTree{edges} }

func (b *BFSTree) Len() int       { return len(b.edges) }
func (b *BFSTree) Edges() []Edge  { return b.edges }
func (b *BFSTree) AddEdge(e Edge) { b.edges = append(b.edges, e) }

// Return unique node names
func (b *BFSTree) Nodes() []string {
	var names []string
	for _, e := range b.edges {
		names = append(names, e.To())
		names = append(names, e.From())
	}
	return uniq(names)
}

// return edges from a given start point
func (b *BFSTree) fromNode(start string) (res []Edge) {
	for _, e := range b.edges {
		if e.From() == start {
			res = append(res, e)
		}
	}
	return res
}

func (b *BFSTree) FindPath(start string, end string) (path *Path, err error) {
	var iter int
	var paths []*Path

	// Create start paths from origin node
	for _, e := range b.fromNode(start) {
		p := newPath(e)
		if e.To() == end {
			return p, nil
		}
		paths = append(paths, p)
	}

	for len(paths) > 0 {
		var newPaths []*Path

		dbg("iter %d\n", iter)

		for _, p := range paths {
			dbg("  %s\n", p)
			children := b.fromNode(p.Last().To())

			// maximum path depth reached, drop
			if len(children) == 0 {
				dbg("    dropped path (no children): %s\n", p)
				continue
			}

			// branch path for each child node
			for _, e := range children {
				// drop circular paths
				if p.IsCircular(e) {
					dbg("    dropped circular child: %s->%s\n", e.From(), e.To())
					continue
				}

				np := newPath(p.edges...)
				np.AddEdge(e)
				dbg("    new path branch: %s\n", np)

				if e.To() == end {
					return np, nil
				}
				newPaths = append(newPaths, np)
			}
		}
		iter++
		paths = newPaths
	}

	return path, fmt.Errorf("no path found")
}

type Path struct {
	*BFSTree
}

func newPath(edges ...Edge) *Path {
	np := &Path{&BFSTree{}}
	for _, e := range edges {
		np.AddEdge(e)
	}
	return np
}

func (p *Path) String() string { return strings.Join(p.Nodes(), "->") }
func (p *Path) Last() Edge     { return p.edges[len(p.edges)-1] }

// Returns names for all path nodes in the order they are transversed
func (p *Path) Nodes() []string {
	names := []string{p.edges[0].From()}
	for _, e := range p.edges {
		names = append(names, e.To())
	}
	return names
}

// Return whether a given edge, if added, would result in
// a circular or recursive path
func (p *Path) IsCircular(edge Edge) bool {
	child := edge.To()
	for _, e := range p.edges {
		if e.From() == child || e.To() == child {
			return true
		}
	}
	return false
}

// Return whether this path transverses a given node name
func (p *Path) HasNode(s string) bool {
	for _, e := range p.edges {
		if e.From() == s || e.To() == s {
			return true
		}
	}
	return false
}

// uniq returns a unique subset of the string slice provided.
func uniq(a []string) []string {
	u := make([]string, 0, len(a))
	m := make(map[string]bool)

	for _, val := range a {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}
