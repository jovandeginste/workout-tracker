package vec

import "sort"

/*
Discreet Data Operations
*/

//Struct for bivariate data (xs, ys)
type BiVariateData struct {
	Xs       []float64
	Ys       []float64
	isSorted bool
}

/*
Convert separate 'xs' and 'ys' arrays into a BiVariateData object

-- 'xs' and 'ys' should be the same length
-- sorts data by x-value automatically

NOTE: sorting is done on slices, which means that 'xs' and 'ys' may be changed
by calling this function
*/
func MakeBiVariateData(xs []float64, ys []float64) *BiVariateData {
	out := BiVariateData{xs, ys, false}
	out.Sort()
	return &out
}

//Len for Sort interface
func (b *BiVariateData) Len() int {
	return len(b.Xs)
}

//Less for Sort interface
func (b *BiVariateData) Less(i, j int) bool {
	if b.Xs[i] < b.Xs[j] {
		return true
	} else {
		return false
	}
}

//Swap for Sort interface
func (b *BiVariateData) Swap(i, j int) {
	b.Xs[i], b.Xs[j] = b.Xs[j], b.Xs[i]
	b.Ys[i], b.Ys[j] = b.Ys[j], b.Ys[i]
	return
}

//Sorts b by ascending x value
func (b *BiVariateData) Sort() {
	//edge case 1
	if b.isSorted {
		return
	}

	//edge case 2
	if sort.Float64sAreSorted(b.Xs) {
		b.isSorted = true
		return
	}

	sort.Sort(b)

	b.isSorted = true
}

//Implement XYer interface for plotinum.plotter
func (b *BiVariateData) XY(n int) (x, y float64) {
	if n >= b.Len() || n < 0 {
		panic("plotinum.plotter.XYer indexed beyond bounds")
	}
	x = b.Xs[n]
	y = b.Ys[n]
	return
}

func (b *BiVariateData) findXBounds(x float64) (int, int) {
	if !b.isSorted {
		b.Sort()
	}
	l := len(b.Xs)

	//edge cases
	if x >= b.Xs[l-1] {
		return l - 2, l - 1
	}
	if x <= b.Xs[0] {
		return 0, 1
	}

	out := 0
	for b.Xs[out] < x && out < l-2 {
		out++
	}

	return out, out + 1

}

/* TODO:
func DiscreetConvolve(dat []float64, conv []float64) []float64 {}

Returns a discreet convolution of 'conv' on 'dat'
*/
