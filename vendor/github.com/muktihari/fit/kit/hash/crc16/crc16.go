// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc16

import (
	"github.com/muktihari/fit/kit/hash"
)

var table = [16]uint16{
	0x0000, 0xCC01, 0xD801, 0x1400, 0xF001, 0x3C00, 0x2800, 0xE401,
	0xA001, 0x6C00, 0x7800, 0xB401, 0x5000, 0x9C01, 0x8801, 0x4400,
}

type crc16 uint16

// New creates a new hash.Hash16 computing the CRC-16 checksum using the polynomial represented by FIT table.
// The computing algorithm is defined in [https://developer.garmin.com/fit/protocol].
// Its Sum method will lay the value out in big-endian byte order.
func New() hash.Hash16 { return new(crc16) }

func (c *crc16) Write(p []byte) (n int, err error) {
	crc := uint16(*c) // PERF: Reduce pointer dereference every time 'crc' is computed.
	for _, b := range p {
		crc = c.compute(crc, b)
	}
	*c = crc16(crc)
	return len(p), nil
}

func (c *crc16) compute(crc uint16, b byte) uint16 {
	var tmp uint16

	// compute checksum of lower four bits of byte
	tmp = table[crc&0xF]
	crc = (crc >> 4) & 0x0FFF
	crc = crc ^ tmp ^ table[b&0xF]

	// now compute checksum of upper four bits of byte
	tmp = table[crc&0xF]
	crc = (crc >> 4) & 0x0FFF
	crc = crc ^ tmp ^ table[(b>>4)&0xF]

	return crc
}

func (c *crc16) Sum16() uint16       { return uint16(*c) }
func (c *crc16) Sum(b []byte) []byte { return append(b, byte(*c>>8), byte(*c)) }
func (c *crc16) Reset()              { *c = 0 }
func (c *crc16) Size() int           { return 2 }
func (c *crc16) BlockSize() int      { return 1 }
