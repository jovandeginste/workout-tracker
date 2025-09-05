// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import (
	"io"
	"math"
)

const (
	// reservedbuf is the maximum bytes that will be requested by the Decoder in one read.
	// The value is obtained from the maximum n field definition in a mesg is 255 and
	// we need 3 byte per field. So 255 * 3 = 765.
	reservedbuf = 765

	minReadBufferSize     = reservedbuf // should not less than this
	maxReadBufferSize     = math.MaxUint32 & math.MaxInt
	defaultReadBufferSize = 4096
)

// readBuffer is a custom buffered reader that will automatically handle buffering, allowing us to
// read bytes directly from the buffer without extra copying, unlike *bufio.Reader which requires us
// to copy the bytes on every Read() method call. When using *bufio.reader we might receive fewer bytes
// than requested, readBuffer returns exactly n requested bytes, otherwise, it returns an error.
type readBuffer struct {
	r io.Reader // reader provided by the client

	// buf is buffer bytes for client reading.
	//
	// Memory layout:
	// [reserved section] + [resizable section]:
	// [0, 1, 2,..., 765] + [766, 767, 768,...]
	//
	// reserved section is used to memmove remaining bytes when remaining < n on reading.
	// resizable section is the space for reading from io.Reader.
	//
	// This way, fragmented remaining bytes is handled and we can always try
	// reading exactly n size bytes from io.Reader.
	//
	// This should be allocated upon creation with minimum len 2*reservedbuf.
	buf []byte

	cur, last int // cur and last of buf positions
}

// Reset resets readBuffer with the new reader and size.
func (b *readBuffer) Reset(r io.Reader, size int) {
	b.r, b.cur, b.last = r, 0, 0

	if size < minReadBufferSize {
		size = minReadBufferSize
	} else if size > maxReadBufferSize {
		size = maxReadBufferSize
	}

	oldsize := cap(b.buf) - reservedbuf
	if size > oldsize {
		b.buf = make([]byte, reservedbuf+size)
	}
	b.buf = b.buf[:reservedbuf+size]
}

// ReadN reads bytes from the buffer and return exactly n bytes.
// If the remaining bytes in the buffer is less than n bytes requested, it will automatically fill the buffer.
// And if it got less than n, an error will be returned.
//
// NOTE: n should be >= 0 and n <= reservedbuf, however, we don't enforce it for efficiency.
func (b *readBuffer) ReadN(n int) ([]byte, error) {
	remaining := b.last - b.cur
	if n > remaining { // fill buf
		cur := reservedbuf
		if remaining != 0 {
			cur = reservedbuf - remaining               // cursor is now pointing at index on 'reserved section'
			copy(b.buf[cur:], b.buf[b.last-remaining:]) // memmove remaining bytes to 'reserved section'.
		}

		nr, err := io.ReadAtLeast(b.r, b.buf[reservedbuf:], n-remaining)
		if err != nil {
			return nil, err
		}
		b.cur = cur
		b.last = reservedbuf + nr
	}

	buf := b.buf[b.cur : b.cur+n]
	b.cur += n
	return buf, nil
}
