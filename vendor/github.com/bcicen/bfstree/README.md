# bfstree

Simple go package providing breadth-first search functions for arbitrary structs

## Usage

```go
package main

import (
	"fmt"
	"github.com/bcicen/bfstree"
)

type FlightRoute struct {
	id       int
	fromCity string
	toCity   string
}

// FlightRoute implements the bfstree.Edge interface
func (f FlightRoute) From() string { return f.fromCity }
func (f FlightRoute) To() string   { return f.toCity }

func main() {
	tree := bfstree.New(
		FlightRoute{0, "New York", "Chicago"},
		FlightRoute{1, "New York", "Los Angeles"},
		FlightRoute{2, "Los Angeles", "Houston"},
		FlightRoute{3, "Chicago", "Tokyo"},
	)

	path, err := tree.FindPath("New York", "Tokyo")
	if err != nil {
		panic(err)
	}

	fmt.Println(path)

	for n, edge := range path.Edges() {
		fmt.Printf("flight %d: %s -> %s\n", n+1, edge.From(), edge.To())
	}
}
```

output:
```
New York->Chicago->Tokyo
flight 1: New York -> Chicago
flight 2: Chicago -> Tokyo
```
