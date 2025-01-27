package fitbit

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"time"

	"github.com/anyappinc/fitbit/logger"
)

var (
	pseudoRand    *mrand.Rand
	RandomByteSet = NumberLetters + UppercaseAlphabetLetters + LowercaseAlphabetLetters // RandomByteSet is a set of characters used to construct random bytes
)

func init() {
	pseudoRand = mrand.New(mrand.NewSource(time.Now().UnixNano()))
}

func randomBytes(length uint64) []byte {
	var (
		cryptoRandErr               error
		randomByteSetLength         = len(RandomByteSet)
		randomByteSetLengthAsBitInt = big.NewInt(int64(randomByteSetLength))
		randBytes                   = make([]byte, length)
	)
	for i := range randBytes {
		if cryptoRandErr == nil {
			idx, err := crand.Int(crand.Reader, randomByteSetLengthAsBitInt)
			if err == nil {
				randBytes[i] = RandomByteSet[idx.Int64()]
			} else {
				logger.Warn.Printf("Can not generate secure random number: %v", err)
				cryptoRandErr = err
			}
		}
		if cryptoRandErr != nil {
			randBytes[i] = RandomByteSet[pseudoRand.Intn(randomByteSetLength)]
		}
	}
	return randBytes
}
