package main

import (
	"fmt"
	"os"
)

var (
	gitRef     = "0.0.0-dev"
	gitRefName = "local"
	gitRefType = "local"
	gitCommit  = "local"
	buildTime  = "now"
)

func main() {
	c, err := newCLI()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := c.rootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
