# rtree

[![GoDoc](https://godoc.org/github.com/tidwall/rtree?status.svg)](https://godoc.org/github.com/tidwall/rtree)

This package provides an in-memory R-Tree implementation for Go. It's designed
for [Tile38](https://github.com/tidwall/tile38) and is optimized for fast rect 
inserts and replacements.

<img src="cities.png" width="512" border="0" alt="Cities">

## Usage

### Installing

To start using rtree, install Go and run `go get`:

```sh
$ go get -u github.com/tidwall/rtree
```

### Basic operations

```go
// create a 2D RTree
var tr rtree.RTree

// insert a point
tr.Insert([2]float64{-112.0078, 33.4373}, [2]float64{-112.0078, 33.4373}, "PHX")

// insert a rect
tr.Insert([2]float64{10, 10}, [2]float64{20, 20}, "rect")

// search 
tr.Search([2]float64{-112.1, 33.4}, [2]float64{-112.0, 33.5}, 
 	func(min, max [2]float64, data interface{}) bool {
		println(data.(string)) // prints "PHX"
	},
)

// delete 
tr.Delete([2]float64{-112.0078, 33.4373}, [2]float64{-112.0078, 33.4373}, "PHX")
```

### Support for Generics (Go 1.18+)

```go
// create a 2D RTree
var tr rtree.RTreeG[string]

// insert a point
tr.Insert([2]float64{-112.0078, 33.4373}, [2]float64{-112.0078, 33.4373}, "PHX")

// insert a rect
tr.Insert([2]float64{10, 10}, [2]float64{20, 20}, "rect")

// search 
tr.Search([2]float64{-112.1, 33.4}, [2]float64{-112.0, 33.5}, 
 	func(min, max [2]float64, data string) bool {
		println(data) // prints "PHX"
	},
)

// delete 
tr.Delete([2]float64{-112.0078, 33.4373}, [2]float64{-112.0078, 33.4373}, "PHX")
```


### Support for generic numeric types, like int, float32, etc.

```go
// create a 2D RTree
var tr rtree.RTreeGN[float32, string]

// insert a point
tr.Insert([2]float32{-112.0078, 33.4373}, [2]float32{-112.0078, 33.4373}, "PHX")

// insert a rect
tr.Insert([2]float32{10, 10}, [2]float32{20, 20}, "rect")

// search 
tr.Search([2]float32{-112.1, 33.4}, [2]float32{-112.0, 33.5}, 
 	func(min, max [2]float32, data string) bool {
		println(data) // prints "PHX"
	},
)

// delete 
tr.Delete([2]float32{-112.0078, 33.4373}, [2]float32{-112.0078, 33.4373}, "PHX")
```


## Algorithms

This implementation is a variant of the original paper:  
[R-TREES. A DYNAMIC INDEX STRUCTURE FOR SPATIAL SEARCHING](https://www.cs.princeton.edu/courses/archive/fall08/cos597B/papers/rtrees.pdf)

### Inserting

Similar to the original algorithm. From the root to the leaf, the rects which
will incur the least enlargment are chosen. Ties go to rects with the smallest
area. 

Added to this implementation: when a rect does not incur any enlargement at
all, it's chosen immediately and without further checks on other rects in the 
same node. This make point insertion faster.

### Deleting

Same as the original algorithm. A target rect is deleted directly. When the number of children in a rect falls below it's minumum entries, it is removed from the tree and it's items are re-inserted.

### Searching

Same as the original algorithm.

### Splitting

This is a custom algorithm.
It attempts to minimize intensive operations such as pre-sorting the children and comparing overlaps & area sizes.
The desire is to do simple single axis distance calculations each child only once, with a target 50/50 chance that the child might be moved in-memory.

When a rect has reached it's max number of entries it's largest axis is calculated and the rect is split into two smaller rects, named `left` and `right`.
Each child rects is then evaluated to determine which smaller rect it should be placed into.
Two values, `min-dist` and `max-dist`, are calcuated for each child. 

- `min-dist` is the distance from the parent's minumum value of it's largest axis to the child's minumum value of the parent largest axis.
- `max-dist` is the distance from the parent's maximum value of it's largest axis to the child's maximum value of the parent largest axis.

When the `min-dist` is less than `max-dist` then the child is placed into the `left` rect. 
When the `max-dist` is less than `min-dist` then the child is placed into the `right` rect. 
When the `min-dist` is equal to `max-dist` then the child is placed into an `equal` bucket until all of the children are evaluated.
Each `equal` rect is then one-by-one placed in either `left` or `right`, whichever has less children.

Finally, sort all the rects in the parent node of the split rect by their
minimum x value.

## License

rtree source code is available under the MIT License.

