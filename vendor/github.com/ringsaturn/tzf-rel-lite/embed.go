package tzfrellite

import (
	_ "embed"
)

//go:embed combined-with-oceans.reduce.pb
var LiteData []byte

//go:embed combined-with-oceans.reduce.compress.pb
var LiteCompressData []byte

//go:embed combined-with-oceans.reduce.preindex.pb
var PreindexData []byte
