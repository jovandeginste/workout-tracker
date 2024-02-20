package rand

import (
	cRand "crypto/rand"
	"math/big"
	mRand "math/rand"
	"time"
)

func init() {
	mRand.Seed(time.Now().UTC().UnixNano())
}

// returns pseudo-random int between min and max, inclusive
func Int(min int, max int) int {
	return mRand.Intn(max-min+1) + min
}

// returns pseudo-random bool
func Bool() bool {
	return Int(0, 1) == 1
}

// returns n pseudo-random bytes
func Bytes(n int) []byte {
	b := make([]byte, n)
	// returned error can be safely ignored as it cannot be non-nil
	// ref https://golang.org/pkg/math/rand/#Read
	mRand.Read(b)
	return b
}

// returns single pseudo-random rune from pool
func Rune(pool []rune) rune {
	return pool[Int(0, len(pool)-1)]
}

// returns string of pseudo-random runes from pool
func String(length int, pool []rune) string {
	out := make([]rune, 0)
	for i := 0; i < length; i++ {
		out = append(out, Rune(pool))
	}
	return string(out)
}

// returns cryptographically secure int between min and max, inclusive
func SecureInt(min int, max int) (int, error) {
	nBig, err := cRand.Int(cRand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()) + min, nil
}

// returns cryptographically secure bool
func SecureBool() (bool, error) {
	randInt, err := SecureInt(0, 1)
	if err != nil || randInt == 0 {
		return false, err
	}
	return true, nil
}

// returns n cryptographically secure bytes
func SecureBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := cRand.Read(b)
	if err != nil {
		return b, err
	}
	return b, nil
}

// returns single cryptographically secure rune
func SecureRune(pool []rune) (rune, error) {
	randInt, err := SecureInt(0, len(pool)-1)
	if err != nil {
		return ' ', err
	}
	return pool[randInt], nil
}

// returns string of cryptographically secure runes from pool
func SecureString(length int, pool []rune) (string, error) {
	out := make([]rune, 0)
	for i := 0; i < length; i++ {
		randRune, err := SecureRune(pool)
		if err != nil {
			return "", err
		}
		out = append(out, randRune)
	}
	return string(out), nil
}

// returns []rune of 0-9
func GetNumericPool() []rune {
	return []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
}

// returns []rune of a-z
func GetAlphabeticLowercasePool() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
}

// returns []rune a-z A-Z
func GetAlphabeticPool() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
}

// returns []rune 0-9 a-z
func GetAlphaNumericLowercasePool() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
}

// returns []rune 0-9 a-z A-Z
func GetAlphaNumericPool() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
}

// returns []rune 0-9 a-z A-Z $ - _ ! ( ) [ ] { } ~ + *
// intended for use in randomly generated tokens such as session ids or api keys
func GetTokenPool() []rune {
	return []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '$', '-', '_', '!', '(', ')', '[', ']', '{', '}', '~', '+', '*'}
}

// returns []rune 1-9, a-z but with characters 1, i, l, 0, a, e, i, o, u removed to prevent ambiguous or offensive output
// useful for use in human-readable ids like shortened urls
func GetUnambiguousLowercasePool() []rune {
	return []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z', '2', '3', '4', '5', '6', '7', '8', '9'}
}

// returns []rune 1-9, a-z, A-Z but with characters 1, i, I, l, 0, a, e, i, o, u, A, E, I, O, U removed to prevent ambiguous or offensive output
// useful for use in human-readable ids like shortened urls
func GetUnambiguousPool() []rune {
	return []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z', '2', '3', '4', '5', '6', '7', '8', '9', 'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z'}
}
