// Package polyline implements a Google Maps Encoding Polyline encoder and
// decoder. See
// https://developers.google.com/maps/documentation/utilities/polylinealgorithm.
//
// The default codec encodes and decodes two-dimensional coordinates scaled by
// 1e5. For other dimensionalities and scales create a custom Codec.
//
// The package operates on byte slices. Encoding functions take an existing byte
// slice as input (which can be nil) and return a new byte slice with the
// encoded value appended to it, similarly to how Go's append function works. To
// increase performance, you can pre-allocate byte slices, for example by
// passing make([]byte, 0, 128) as the input byte slice. Similarly, decoding
// functions take a byte slice as input and return the remaining unconsumed
// bytes as output.
package polyline

import (
	"errors"
	"math"
	"strconv"
)

// Errors.
var (
	ErrDimensionalMismatch  = errors.New("dimensional mismatch")
	ErrEmpty                = errors.New("empty")
	ErrInvalidByte          = errors.New("invalid byte")
	ErrOverflow             = errors.New("overflow")
	ErrUnterminatedSequence = errors.New("unterminated sequence")
)

func round(x float64) int {
	if x < 0 {
		return int(-math.Floor(-x + 0.5))
	}
	return int(math.Floor(x + 0.5))
}

// A Codec represents an encoder.
type Codec struct {
	Dim   int     // Dimensionality, normally 2
	Scale float64 // Scale, normally 1e5
}

var defaultCodec = Codec{Dim: 2, Scale: 1e5}

// DecodeUint decodes a single unsigned integer from buf. It returns the decoded
// uint, the remaining unconsumed bytes of buf, and any error.
func DecodeUint(buf []byte) (uint, []byte, error) {
	if len(buf) == 0 {
		return 0, nil, ErrEmpty
	}
	n := strconv.IntSize / 5
	if n > len(buf) {
		n = len(buf)
	}
	var u, shift uint
	for i := 0; i < n; i++ {
		switch b := buf[i]; {
		case 95 <= b && b < 127:
			u += (uint(b) - 95) << shift
			shift += 5
		case 63 <= b && b < 95:
			u += (uint(b) - 63) << shift
			return u, buf[i+1:], nil
		default:
			return 0, nil, ErrInvalidByte
		}
	}
	if len(buf) <= strconv.IntSize/5 {
		return 0, nil, ErrUnterminatedSequence
	}
	max := byte(1<<(strconv.IntSize-5*(strconv.IntSize/5)) - 1)
	switch b := buf[n]; {
	case 63 <= b && b <= 63+max:
		u += (uint(b) - 63) << shift
		return u, buf[n+1:], nil
	case b < 127:
		return 0, nil, ErrOverflow
	default:
		return 0, nil, ErrInvalidByte
	}
}

// DecodeInt decodes a single signed integer from buf. It returns the decoded
// int, the remaining unconsumed bytes of buf, and any error.
func DecodeInt(buf []byte) (int, []byte, error) {
	switch u, buf, err := DecodeUint(buf); {
	case err != nil:
		return 0, nil, err
	case u&1 == 0:
		return int(u >> 1), buf, nil
	case u == maxUint:
		return minInt, buf, nil
	default:
		return -int((u + 1) >> 1), buf, nil
	}
}

// EncodeUint appends the encoding of a single unsigned integer u to buf and
// returns the new buf.
func EncodeUint(buf []byte, u uint) []byte {
	for u >= 32 {
		buf = append(buf, byte((u&31)+95))
		u >>= 5
	}
	buf = append(buf, byte(u+63))
	return buf
}

// EncodeInt appends the encoding of a single signed integer i to buf and
// returns the new buf.
func EncodeInt(buf []byte, i int) []byte {
	var u uint
	if i < 0 {
		u = uint(^(i << 1))
	} else {
		u = uint(i << 1)
	}
	return EncodeUint(buf, u)
}

// DecodeCoord decodes a single coordinate from buf. It returns the coordinate,
// the remaining unconsumed bytes of buf, and any error.
func (c Codec) DecodeCoord(buf []byte) ([]float64, []byte, error) {
	coord := make([]float64, c.Dim)
	for i := range coord {
		var err error
		var j int
		j, buf, err = DecodeInt(buf)
		if err != nil {
			return nil, nil, err
		}
		coord[i] = float64(j) / c.Scale
	}
	return coord, buf, nil
}

// DecodeCoords decodes an array of coordinates from buf. It returns the
// coordinates, the remaining unconsumed bytes of buf, and any error.
func (c Codec) DecodeCoords(buf []byte) ([][]float64, []byte, error) {
	if len(buf) == 0 {
		return nil, buf, nil
	}
	var coord []float64
	var err error
	coord, buf, err = c.DecodeCoord(buf)
	if err != nil {
		return nil, nil, err
	}
	coords := [][]float64{coord}
	for i := 1; len(buf) > 0; i++ {
		coord, buf, err = c.DecodeCoord(buf)
		if err != nil {
			return nil, nil, err
		}
		for j := range coord {
			coord[j] += coords[i-1][j]
		}
		coords = append(coords, coord)
	}
	return coords, nil, nil
}

// DecodeFlatCoords decodes coordinates from buf, appending them to a
// one-dimensional array. It returns the coordinates, the remaining unconsumed
// bytes in buf, and any error.
func (c Codec) DecodeFlatCoords(flatCoords []float64, buf []byte) ([]float64, []byte, error) {
	if len(flatCoords)%c.Dim != 0 {
		return nil, nil, ErrDimensionalMismatch
	}
	last := make([]int, c.Dim)
	for len(buf) > 0 {
		for j := 0; j < c.Dim; j++ {
			var err error
			var k int
			k, buf, err = DecodeInt(buf)
			if err != nil {
				return nil, nil, err
			}
			last[j] += k
			flatCoords = append(flatCoords, float64(last[j])/c.Scale)
		}
	}
	return flatCoords, nil, nil
}

// EncodeCoord encodes a single coordinate to buf and returns the new buf.
func (c Codec) EncodeCoord(buf []byte, coord []float64) []byte {
	for _, x := range coord {
		buf = EncodeInt(buf, round(c.Scale*x))
	}
	return buf
}

// EncodeCoords appends the encoding of an array of coordinates coords to buf
// and returns the new buf.
func (c Codec) EncodeCoords(buf []byte, coords [][]float64) []byte {
	last := make([]int, c.Dim)
	for _, coord := range coords {
		for i, x := range coord {
			ex := round(c.Scale * x)
			buf = EncodeInt(buf, ex-last[i])
			last[i] = ex
		}
	}
	return buf
}

// EncodeFlatCoords encodes a one-dimensional array of coordinates to buf. It
// returns the new buf and any error.
func (c Codec) EncodeFlatCoords(buf []byte, flatCoords []float64) ([]byte, error) {
	if len(flatCoords)%c.Dim != 0 {
		return nil, ErrDimensionalMismatch
	}
	last := make([]int, c.Dim)
	for i, x := range flatCoords {
		ex := round(c.Scale * x)
		j := i % c.Dim
		buf = EncodeInt(buf, ex-last[j])
		last[j] = ex
	}
	return buf, nil
}

// DecodeCoord decodes a single coordinate from buf using the default codec. It
// returns the coordinate, the remaining bytes in buf, and any error.
func DecodeCoord(buf []byte) ([]float64, []byte, error) {
	return defaultCodec.DecodeCoord(buf)
}

// DecodeCoords decodes an array of coordinates from buf using the default
// codec. It returns the coordinates, the remaining bytes in buf, and any error.
func DecodeCoords(buf []byte) ([][]float64, []byte, error) {
	return defaultCodec.DecodeCoords(buf)
}

// EncodeCoord returns the encoding of an array of coordinates using the default
// codec.
func EncodeCoord(coord []float64) []byte {
	return defaultCodec.EncodeCoord(nil, coord)
}

// EncodeCoords returns the encoding of an array of coordinates using the
// default codec.
func EncodeCoords(coords [][]float64) []byte {
	return defaultCodec.EncodeCoords(nil, coords)
}
