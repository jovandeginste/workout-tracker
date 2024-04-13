package vec

import (
	"runtime"
	"sync"
)

/* Simple Vector Operation

*/
func simpleVO(f BiMathop, arrOne []float64, arrTwo []float64, outVec []float64, start int, end int) {
	for i := start; i <= end; i++ {
		outVec[i] = f(arrOne[i], arrTwo[i])
	}
	return
}

/* Parallel Vector Operation
- Creates a new slice from two other slices according to 'f()'
- Uses NumCPU() goroutines for workload partioning (defaults to 2)
*/
func PVecOperation(f BiMathop, arrOne []float64, arrTwo []float64) []float64 {
	l := len(arrOne)
	if l != len(arrTwo) {
		panic("PVecOperation must be performed on slices of identical length.")
	}
	NTHREADS := runtime.NumCPU()

	//Just in case...
	if NTHREADS == 0 {
		NTHREADS = 2
	}

	batch_size := l / NTHREADS
	rem := l % NTHREADS
	outVec := make([]float64, l)

	//Simple Case
	if batch_size == 0 {
		simpleVO(f, arrOne, arrTwo, outVec, 0, l-1)
		return outVec
	} else {
		//Parallel Case
		start := 0
		end := batch_size - 1
		wg := new(sync.WaitGroup)
		wg.Add(NTHREADS)

		for i := 0; i < NTHREADS-1; i++ {
			go func(s int, e int) {
				simpleVO(f, arrOne, arrTwo, outVec, s, e)
				wg.Done()
			}(start, end)
			start += batch_size
			end += batch_size
		}
		go func(s int, e int) {
			simpleVO(f, arrOne, arrTwo, outVec, s, e)
			wg.Done()
		}(start, end+rem)

		wg.Wait()
		return outVec
	}

}
