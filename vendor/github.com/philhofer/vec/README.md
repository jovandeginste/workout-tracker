Vec for go [![Build Status](https://travis-ci.org/philhofer/vec.png?branch=master)](https://travis-ci.org/philhofer/vec)
====================
Vec is currently in the development stage. Source is covered by the Mozilla Public License v. 2.0.

Objectives
------------

* Make a Go library for higher math and science (a la GSL / numpy /
  scipy), written in Go
* Functionality including, but not limited to, integration, differentiation, interpolation, matrix operations, local max/minimization, global max/minization, optimization

Example - Current Functionality
-----------
Here were going to define a function, map it onto a set of points, create a BiVariateData object from those points, and then create a CubicSplineInterpolation object with which we can evaluate derivatives and integrals.

```
import "github.com/philhofer/vec"
import "math"

//our function
func myFunc(x float64) float64 {
     return 3*x*x - 2*x + x - 1
}

//create a slice of 100 x-points from 0 to 5
//create another identical slice that we'll map myFunc() onto
xs := vec.Arange(0, 5, 100)
ys := vec.Arange(0, 5, 100)

//map 'myFunc()' onto 'ys' (very fast; uses NumCPU() parallel goroutines)
vec.PPmap(myFunc, ys)

//create a BiVariateData object from our data (returns a pointer)
bvd := vec.MakeBiVariateData(xs, ys)

//make a cubic spline from our data
spl := vec.CubicSpline(*bvd)

//evaluate spline at 3.2678
x := spl.F(3.2678)

//evaluate the first and second  derivatives of the spline at 3.2678
dx := spl.DF(3.2678)
ddx := spl.DDF(3.2678)

//evaluate the integral of myFunc() from 1 to 4
Ix := spl.Integral(1, 4)

//we can also evaluate infinite bounds!
func myinfFunc(x float64) float64 {
     return math.Exp(-5.0*x)
}

num := vec.Integral(myinfFunc, 0, math.Inf(1))

```

Contributing
-------------

* Read the documentation before commiting code
* Use types already defined in the package if at all possible
* Check out the feature requests and future functionality on github.com/philhofer/go-vec.git
