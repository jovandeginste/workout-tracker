package vec

import (
	"math/rand"
)

type MHParams interface {
	Prob([]float64) float64 //proportional to desired stationary distribution
	Next([]float64) []float64 //proposition function (must be symmetric)
}

/*Creates a Markov chain of length N and width of len(model.Next())
Accepts steps in the chain based on the Metropolis-Hastings criteria.
*/
func MarkovChain(model MHParams, start []float64, N int) (out [][]float64) {
	out = make([][]float64, N)

	p0 := model.Prob(start)
	if len(model.Next(start)) != len(start) {
		panic("Next() is ill-formed: Returns the wrong number of floats.")
	}

	//Burn-in
	i:=0
	for i<1000 {
		next := model.Next(start)
		p1 := model.Prob(next)
		if p1/p0 > rand.Float64() {
			start = next
			p0 = p1
			i++
		}
	}

	//Run
	i = 0
	for i<N {
		next := model.Next(start)
		p1 := model.Prob(next)
		if p1/p0 > rand.Float64() {
			start = next
			p0 = p1
			out[i] = next
			i++
		}
	}
	return out
}
