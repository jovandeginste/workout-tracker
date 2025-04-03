package tzfrellite

import (
	_ "embed"
)

//go:embed combined-with-oceans.reduce.bin
var LiteData []byte

//go:embed combined-with-oceans.reduce.compress.bin
var LiteCompressData []byte

//go:embed combined-with-oceans.reduce.preindex.bin
var PreindexData []byte
