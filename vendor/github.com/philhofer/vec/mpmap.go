package vec

import (
	"github.com/philhofer/vec/pool"
	"runtime"
)
/* Load-balanced Parallel Function-Slice Mapping
Creates a pool of workers that map 'fm' onto 'arr'
*/
func MPmap(fm Mathop, arr []float64) {
	p := pool.NewPool(runtime.NumCPU(), len(arr))
	each := func(i int) pool.Proc {
		return func() {
			arr[i] = fm(arr[i])
	
		}
	}
	for i := range arr {
		p.Send(each(i))
	}
	p.WaitAll()
}
