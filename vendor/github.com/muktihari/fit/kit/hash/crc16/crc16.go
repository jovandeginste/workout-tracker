// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc16

import (
	"github.com/muktihari/fit/kit/hash"
)

// Table is [16]uint16
type Table [16]uint16

const (
	Size = 2 // An uint16 requires 2 bytes to be represented in its binary form.
)

var fitTable = Table{
	0x0000, 0xCC01, 0xD801, 0x1400, 0xF001, 0x3C00, 0x2800, 0xE401,
	0xA001, 0x6C00, 0x7800, 0xB401, 0x5000, 0x9C01, 0x8801, 0x4400,
}

// MakeFITTable makes new table as defined in [https://developer.garmin.com/fit/protocol]
func MakeFITTable() *Table {
	t := fitTable
	return &t
}

// New creates a new hash.Hash16 computing the CRC-16 checksum using the polynomial represented by the Table.
// If table is nil, default FIT table will be used. The computing algorithm is using FIT algorithm defined in
// [https://developer.garmin.com/fit/protocol]. Its Sum method will lay the value out in big-endian byte order.
func New(table *Table) hash.Hash16 {
	if table == nil {
		table = &fitTable
	}
	return &crc16{table: table}
}

type crc16 struct {
	table *Table
	crc   uint16
}

func (c *crc16) Write(p []byte) (n int, err error) {
	crc := c.crc // PERF: Reduce pointer dereference every time 'c.crc' is computed.
	for _, b := range p {
		crc = c.compute(crc, b)
	}
	c.crc = crc
	return len(p), nil
}

func (c *crc16) Sum(b []byte) []byte {
	s := c.Sum16()
	return append(b, byte(s>>8), byte(s))
}

func (c *crc16) Reset() { c.crc = 0 }

func (c *crc16) Size() int { return Size }

func (c *crc16) BlockSize() int { return 1 }

func (c *crc16) Sum16() uint16 { return c.crc }

func (c *crc16) compute(crc uint16, b byte) uint16 {
	var tmp uint16

	// compute checksum of lower four bits of byte
	tmp = c.table[crc&0xF]
	crc = (crc >> 4) & 0x0FFF
	crc = crc ^ tmp ^ c.table[b&0xF]

	// now compute checksum of upper four bits of byte
	tmp = c.table[crc&0xF]
	crc = (crc >> 4) & 0x0FFF
	crc = crc ^ tmp ^ c.table[(b>>4)&0xF]

	return crc
}
