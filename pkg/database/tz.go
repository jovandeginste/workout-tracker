package database

import "github.com/ringsaturn/tzf/v2"

var tzFinder tzf.F

func InitTZFinder() error {
	var err error

	tzFinder, err = tzf.NewDefaultFinder()

	return err
}
