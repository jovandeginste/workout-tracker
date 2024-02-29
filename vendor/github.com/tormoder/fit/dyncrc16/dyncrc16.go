// Package dyncrc16 implements the Dynastream CRC-16 checksum.
package dyncrc16

const size = 2

var crcTable = [...]uint16{
	0x0000, 0xCC01, 0xD801, 0x1400,
	0xF001, 0x3C00, 0x2800, 0xE401,
	0xA001, 0x6C00, 0x7800, 0xB401,
	0x5000, 0x9C01, 0x8801, 0x4400,
}

// crc16 represents the partial evaluation of a checksum.
type crc16 uint16

// New returns a new hash.Hash16 computing the Dynastream CRC-16 checksum.
func New() Hash16 {
	c := new(crc16)
	return c
}

func (c *crc16) Sum16() uint16 {
	return uint16(*c)
}

func (c *crc16) Reset() {
	*c = 0
}

func (c *crc16) Size() int {
	return size
}

func (c *crc16) BlockSize() int {
	return 1
}

func (c *crc16) Sum(in []byte) []byte {
	s := c.Sum16()
	return append(in, byte(s>>8), byte(s))
}

func (c *crc16) Write(data []byte) (int, error) {
	*c = update(*c, data)
	return len(data), nil
}

// Checksum returns the Dynastream CRC-16 checksum of data.
func Checksum(data []byte) uint16 {
	var c crc16
	c = update(c, data)
	return c.Sum16()
}

// Add data to the running checksum c.
func update(c crc16, data []byte) crc16 {
	for _, d := range data {
		c = updateByte(c, d)
	}
	return c
}

// Add data to the running checksum c.
func updateByte(c crc16, data byte) crc16 {
	d := uint16(c)

	// compute checksum of lower four bits of byte
	tmp := crcTable[d&0x0F]
	d = (d >> 4) & 0x0FFF
	d = d ^ tmp ^ crcTable[data&0x0F]

	// now compute checksum of upper four bits of byte
	tmp = crcTable[d&0x0F]
	d = (d >> 4) & 0x0FFF
	d = d ^ tmp ^ crcTable[(data>>4)&0x0F]

	return crc16(d)
}
