package dyncrc16

import "hash"

// Hash16 is the common interface implemented by some 16-bit hash functions.
type Hash16 interface {
	hash.Hash
	Sum16() uint16
}
