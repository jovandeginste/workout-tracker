package vec

import (
	"runtime"
	"sync"
)
/* Iterative Function->Slice Mapping
Maps a fuction onto an array on the values from
'start' to 'end' (typically '0' to 'len(arr)')
No concurrency
*/
func Smap(fm Mathop, arr []float64, start int, end int) {
	for i := start; i < end; i++ {
		arr[i] = fm(arr[i])
	}
}

/* Partitioned Funtion->Slice Mapping
Maps 'fm()' onto each member of 'arr' in-place
Uses 'NumCPU()' independent (non-load-balanced) goroutines
*/
func PPmap(fm Mathop, arr []float64) {
	NTHREADS := runtime.NumCPU()
	//test for nonsense
	if NTHREADS <= 0 {
		NTHREADS = 2
	}
	runtime.GOMAXPROCS(NTHREADS)
	l := len(arr)
	batch_size := l / NTHREADS
	rem := l % NTHREADS

	// Do simple mapping if fewer than NTHREADS items
	if batch_size == 0 {
		Smap(fm, arr, 0, len(arr))
		return
	} else {
		wg := new(sync.WaitGroup)
		wg.Add(NTHREADS)
		start := 0
		end := batch_size

		//Spawn (NTHREADS-1) goroutines with array of length 'batch_size'
		for i := 0; i < NTHREADS-1; i++ {
			go func(s int, e int) {
				Smap(fm, arr, s, e)
				wg.Done()
			}(start, end)
			start += batch_size
			end += batch_size
		}

		//Spawn last goroutine with array of length 'batch_size + rem'
		go func(s int, e int) {
			Smap(fm, arr, s, e)
			wg.Done()
		}(start, end+rem)

		wg.Wait()
		return
	}
}
