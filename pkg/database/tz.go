package database

import "github.com/ringsaturn/tzf"

var (
	tzFinder tzf.F
	tzInit   = false
)

func initTZFinder() {
	if tzInit {
		return
	}

	tzFinder, _ = tzf.NewDefaultFinder()
	tzInit = true
}
