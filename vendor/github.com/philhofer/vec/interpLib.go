package vec

import "math"

/*
Return a slice of length N
with equally-spaced values from 'start' to 'stop'
(non-inclusive)
*/

func notRat(a float64) bool {
	if math.IsNaN(a) || math.IsInf(a, 0) {
		return true
	}
	return false
}

func Arange(start float64, stop float64, N int) []float64 {
	if notRat(start) || notRat(stop) {
		return []float64{}
	}
	if N == 0 {
		return []float64{}
	}
	f := make([]float64, N)
	h := (stop - start) / float64(N)
	for i := range f {
		f[i] = start + (float64(i) * h)
	}
	return f
}

/*
CubicSplineInterpolation type def
points to a BiVariateData container and contains
an array of coefficients used for evaluating
the interpolation, its derivatives, and its
definite integral
*/
type CubicSplineInterpolation struct {
	data   *BiVariateData
	coeffs []float64
}

/*
Default constructor for CubicSplineInterpolation
Returns a fully-formed CubicSplineInterpolation from
the data containe in 'd'

*Note: this is a 'natural' cubic spline (meaning that
points evaluated outside the data endpoints are extrapolated
linearly)

Feeds values to getCubicSplinCoeffs() that are equivalent to
the right-hand side of the matrix equation (18) referenced in:
http://mathworld.wolfram.com/CubicSpline.html
*/
func CubicSpline(d *BiVariateData) *CubicSplineInterpolation {
	spline := &CubicSplineInterpolation{data: d}
	N := len(d.Ys)

	//Edge case - len(Ys) is zero
	//Return default contructor
	if N == 0 {
		return spline
	}

	fs := make([]float64, N)
	fs[0] = 3 * (d.Ys[1] - d.Ys[0])
	for i := 1; i < (N - 1); i++ {
		fs[i] = 3 * (d.Ys[i+1] - d.Ys[i-1])
	}
	fs[N-1] = 3 * (d.Ys[N-1] - d.Ys[N-2])

	getCubicSplinCoeffs(fs)
	spline.coeffs = fs

	return spline
}

/*
Returns the interpolated value of 'x'
on the data pointed to by 's'
*/
func (s *CubicSplineInterpolation) F(x float64) float64 {
	//Edge case - x is not rational
	if notRat(x) {
		return math.NaN()
	}

	i, i1 := s.data.findXBounds(x)
	t := (x - s.data.Xs[i]) / (s.data.Xs[i1] - s.data.Xs[i])
	yi, yi1 := s.data.Ys[i], s.data.Ys[i1]
	Di, Di1 := s.coeffs[i], s.coeffs[i+1]
	a := yi
	b := Di
	c := 3*(yi1-yi) - 2*Di - Di1
	d := 2*(yi-yi1) + Di + Di1

	return a + b*t + c*t*t + d*t*t*t
}

/*
First derivative of interpolated data
evaluated at 'x'

*/
func (s *CubicSplineInterpolation) DF(x float64) float64 {
	//Edge case - x is not rational
	if notRat(x) {
		return math.NaN()
	}

	i, i1 := s.data.findXBounds(x)
	fin := s.data.Xs[i1]
	start := s.data.Xs[i]
	t := (x - start) / (fin - start)
	yi, yi1 := s.data.Ys[i], s.data.Ys[i1]
	Di, Di1 := s.coeffs[i], s.coeffs[i+1]
	b := Di
	c := 3*(yi1-yi) - 2*Di - Di1
	d := 2*(yi-yi1) + Di + Di1

	return (b + 2*c*t + 3*d*t*t) / (fin - start)
}

/*
Second derivative evaluated at 'x'
*/
func (s *CubicSplineInterpolation) DDF(x float64) float64 {

	if notRat(x) {
		return math.NaN()
	}

	i, i1 := s.data.findXBounds(x)
	fin := s.data.Xs[i1]
	start := s.data.Xs[i]
	t := (x - start) / (fin - start)
	yi, yi1 := s.data.Ys[i], s.data.Ys[i1]
	Di, Di1 := s.coeffs[i], s.coeffs[i+1]
	c := 3*(yi1-yi) - 2*Di - Di1
	d := 2*(yi-yi1) + Di + Di1

	return (2*c + 6*d*t) / ((fin - start) * (fin - start))
}

/*
Returns the definite integral of a cubic spline
evaluated from 'a' to 'b'

*Note: 'a' and 'b' must be rational
(can't be +/-Inf -- will return NaN)
*/
func (s *CubicSplineInterpolation) Integral(a float64, b float64) float64 {

	//Does not support +/- Inf bounds
	if notRat(a) || notRat(b) {
		return math.NaN()
	}

	ia, ia1 := s.data.findXBounds(a)
	ib, ib1 := s.data.findXBounds(b)

	//t1 - t up to nearest datapoint
	t1 := (a - s.data.Xs[ia]) / (s.data.Xs[ia1] - s.data.Xs[ia])
	//t2 - t after last datapoint
	t2 := (b - s.data.Xs[ib]) / (s.data.Xs[ib1] - s.data.Xs[ib])
	out := 0.0

	//context before first datapoint in integral range
	for {
		yi, yi1 := s.data.Ys[ia], s.data.Ys[ia1]
		Di, Di1 := s.coeffs[ia], s.coeffs[ia1]
		a := yi
		b := Di
		c := 3*(yi1-yi) - 2*Di - Di1
		d := 2*(yi-yi1) + Di + Di1

		out += (s.data.Xs[ia1] - s.data.Xs[ia]) * (a + (b / 2.0) + (c / 3.0) + (d / 4.0) - a*t1 - (b*t1*t1)/2.0 - (c*t1*t1*t1)/3.0 - (d*t1*t1*t1*t1)/4.0)
		break
	}

	//middle contexts (within datapoint & integral range)
	for i := ia + 1; i < ib; i++ {
		yi, yi1 := s.data.Ys[i], s.data.Ys[i+1]
		Di, Di1 := s.coeffs[i], s.coeffs[i+1]
		a := yi
		b := Di
		c := 3*(yi1-yi) - 2*Di - Di1
		d := 2*(yi-yi1) + Di + Di1

		out += (s.data.Xs[i+1] - s.data.Xs[i]) * (a + (b / 2.0) + (c / 3.0) + (d / 4.0))
	}

	//context after last datapoint in integral range
	for {
		yi, yi1 := s.data.Ys[ib], s.data.Ys[ib1]
		Di, Di1 := s.coeffs[ib], s.coeffs[ib1]
		a := yi
		b := Di
		c := 3*(yi1-yi) - 2*Di - Di1
		d := 2*(yi-yi1) + Di + Di1

		out += (a*t2 + (b*t2*t2)/2.0 + (c*t2*t2*t2)/3.0 + (d*t2*t2*t2*t2)/4.0) * (s.data.Xs[ib1] - s.data.Xs[ib])
		break
	}

	return out
}

func makeConstVec(a float64, N int) []float64 {
	out := make([]float64, N)
	for i := range out {
		x := a
		out[i] = x
	}
	return out
}

/*
Computes the coefficients for
cubic spline interpolation from an array of
values by means of solving a tridiagonal matrix
using gaussian elimination and back-substitution

See:
http://mathworld.wolfram.com/CubicSpline.html
Eqn (18)
*/
func getCubicSplinCoeffs(x []float64) {
	l := len(x)

	//subdiagonal
	a := makeConstVec(1.0, l)

	//main diag
	b := makeConstVec(4.0, l)
	b[0] = 2.0
	b[l-1] = 2.0

	//superdiagonal
	c := makeConstVec(1.0, l)

	c[0] = (c[0] / b[0])
	x[0] = (x[0] / b[0])

	//gaussian elimination
	for i := 1; i < l; i++ {
		m := 1.0 / (b[i] - (a[i] * c[i-1]))
		c[i] = c[i] * m
		x[i] = (x[i] - (a[i] * x[i-1])) * m
	}

	//backsubstitution
	for i := l - 2; i >= 0; i-- {
		x[i] = x[i] - c[i]*x[i+1]
	}

}
