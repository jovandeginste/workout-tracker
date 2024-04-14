package vec

import "math"

/* Numerical Derivative
Calculates the numerical derivative of f(x) using Richardson
Extrapolation of the central difference. Returns a boolean convergence
value to indicate if the derivative is numerically stable.
*/
func Deriv(f Mathop, x float64) (out float64, conv bool) {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return math.NaN(), false
	}
	//"Optimum" h0 to reduce roundoff error @ O(h^6)
	h0 := math.Sqrt(10.0) / 1000.0
	A0 := func(h float64) float64 {
		return (f(x+h) - f(x-h)) / (2 * h)
	}
	A1 := func(h float64) float64 {
		return (4*A0(h/2) - A0(h)) / 3
	}
	A2 := func(h float64) float64 {
		return (16*A1(h/4) - A1(h)) / 15
	}

	a1, a2, a3 := A0(h0), A1(h0), A2(h0)
	out = a3
	//test for convergence
	if math.Abs(a1-a2) < math.Abs(a2-a3) {
		conv = false
	}
	conv = true
	return
}
