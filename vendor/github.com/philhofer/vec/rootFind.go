package vec

import (
	"math"
	//"fmt"
)

//'true' if opposite signs; false otherwise
func opp(x float64, y float64) bool {
	if math.Signbit(x) == math.Signbit(y) {
		return false
	} else {
		return true
	}
}

//reverse quadratic interpolation
func rqi(f Mathop, a float64, b float64, c float64) float64 {
	s1 := a * (f(b) * f(c)) / ((f(a) - f(b)) * (f(a) - f(c)))
	s2 := b * (f(a) * f(c)) / ((f(b) - f(a)) * (f(b) - f(c)))
	s3 := c * (f(a) * f(b)) / ((f(c) - f(a)) * (f(c) - f(b)))
	return s1 + s2 + s3
}

/*
FindRoot implements "Brent's Method" for numerical root-finding.
Returns x such that f(x) is 0.
Returns conv=false if the root-finding failed.
For success, f(a) and f(b) must have opposite signs, and f(x) must be continuous.
*/
func FindRoot(f Mathop, a float64, b float64) (out float64, conv bool) {
	const acc float64 = 10E-15
	//check for opposite sign
	if !opp(f(a), f(b)) {
		return math.NaN(), false
	}

	//returns magnitude of f(x)
	fmag := func(x float64) float64 {
		return math.Abs(f(x))
	}

	//'b' must be the better guess
	if fmag(a) < fmag(b) {
		a, b = b, a
	}

	s := 0.0
	d := a
	c := a
	flag := true
	check := func() bool {
		if (s > b && s < ((3*a+b)/4) && b > ((3*a+b)/4)) || (s < b && s > ((3*a+b)/4) && b < ((3*a+b)/4)) {
			return true
		} else if flag && math.Abs(s-b) >= (math.Abs(b-c)/2.0) {
			return true
		} else if !flag && math.Abs(s-b) >= (math.Abs(c-d)/2.0) {
			return true
		} else if flag && math.Abs(b-c) < acc {
			return true
		} else if !flag && math.Abs(c-d) < acc {
			return true
		} else {
			return false
		}
	}

	for math.Abs(b-a) > acc {
		if a != c && b != c {
			s = rqi(f, a, b, c)
		} else {
			s = b - f(b)*(b-a)/(f(b)-f(a))
		}
		if check() {
			s = (a + b) / 2
			flag = true
		} else {
			flag = false
		}
		d, c = c, b
		if f(a)*f(s) < 0 {
			b = s
		} else {
			a = s
		}
		if fmag(a) < fmag(b) {
			a, b = b, a
		}
	}
	return s, true
}
