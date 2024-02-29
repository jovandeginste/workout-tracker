//go:build gofuzz
// +build gofuzz

package fit

import (
	"bytes"
	"encoding/binary"
)

func Fuzz(data []byte) int {
	fitFile, err := Decode(bytes.NewReader(data))
	if err != nil {
		return 0
	}

	var w bytes.Buffer
	err = Encode(&w, fitFile, binary.LittleEndian)
	if err != nil {
		panic(err)
	}

	return 1
}
