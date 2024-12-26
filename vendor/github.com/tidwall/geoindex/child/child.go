package child

// Child represents a child of a 2d geospatial tree.
// The Min and Max fields are the bounds of the Child.
// Data is whatever the child consists of.
// Item is true when the Data is a leaf item, otherwise it's probably a node.
type Child struct {
	Min, Max [2]float64
	Data     interface{}
	Item     bool
}
