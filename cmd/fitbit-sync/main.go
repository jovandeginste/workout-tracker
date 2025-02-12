package main

import (
	"fmt"
	"os"
)

func main() {
	c := newCLI()
	if err := c.rootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
