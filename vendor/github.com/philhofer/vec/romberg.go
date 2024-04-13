package vec

import (
	"math"
	"sync"
	//	"fmt"
)

//Trapezoidal rule - 'f(x)' from 'a' to 'b' with N steps
func trap(f Mathop, a float64, b float64, N int) float64 {
	if N <= 0 {
		return a
	}
	if N == 1 {
		return (b - a) * (f(b) + f(a)) / 2
	}

	out := 0.0
	h := (b - a) / float64(N)
	xs := Arange(a, b+h, N+1)
	PPmap(f, xs)
	for i := 0; i < N; i++ {
		out += xs[i] + xs[i+1]
	}
	return out * (h / 2.0)
}

/*
Performs Romberg integration on a Mathop

Returns the integral evaluated from 'a' to 'b', and convergence.
Supports (+/-)Inf as bounds.

If conv = false, the integral did not converge with
If conv = true, the integral is accurate to at least 15 decimal places.
*/
func Integral(f Mathop, a float64, b float64) (out float64, conv bool) {
	conv = false
	out = 0.0
	const K int = 10
	if a == b {
		return out, true
	}
	if math.IsNaN(a) || math.IsNaN(b) {
		return math.NaN(), false
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return rangeTransform(f, a, b)
	}

	var Ip []float64
	for k := 0; k < K; k++ {
		Ik := make([]float64, K-k)

		if k == 0 {
			wg := new(sync.WaitGroup)
			wg.Add(len(Ik))
			for j := range Ik {
				go func(i int) {
					Ik[i] = trap(f, a, b, int(math.Pow(2, float64(i))))
					wg.Done()
				}(j)
			}
			wg.Wait()
		} else {
			for i := range Ik {
				j := i + 1
				m := math.Pow(4.0, float64(k))
				Ik[i] = (m*Ip[j] - Ip[j-1]) / (m - 1.0)
			}
		}

		out = Ik[K-k-1]
		if k > 0 {
			err := math.Abs(Ik[K-k-1] - Ip[K-k])
			if err <= 1E-16 {
				return out, true
			}
		} else {
			if math.Abs(Ik[K-k-2]-Ik[K-k-1]) <= 1E-16 {
				return out, true
			}
		}

		Ip = Ik
	}

	return out, false
}

func rangeTransform(f Mathop, a float64, b float64) (float64, bool) {

	flipped := false
	if a > b {
		a, b = b, a
		flipped = true
	}

	X := func(z float64) float64 {
		return -z / ((z - 1) * (z + 1))
	}

	fnew := func(z float64) float64 {
		return f(X(z)) * (z*z + 1) / math.Pow((z*z-1), 2.0)
	}

	var anew, bnew float64
	if a == 0 {
		anew = 0.0
	} else if math.IsInf(a, -1) {
		anew = -1.0 + (1E-15)
	} else {
		anew = (math.Sqrt(4*a*a+1) - 1.0) / (2 * a)
	}
	if b == 0 {
		bnew = 0.0
	} else if math.IsInf(b, 1) {
		bnew = 1.0 - (1E-15)
	} else {
		bnew = (math.Sqrt(4*b*b+1) - 1.0) / (2 * b)
	}

	out, conv := Integral(fnew, anew, bnew)
	if flipped {
		out *= -1.0
	}

	return out, conv
}
