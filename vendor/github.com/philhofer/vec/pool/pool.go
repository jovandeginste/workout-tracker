package pool

import "sync"

type Proc func()

type Pool struct {
	wg   *sync.WaitGroup
	pipe chan Proc
}

func listen(c chan Proc, wg *sync.WaitGroup) {
	defer wg.Done()
	for p := range c {
		p()
	}
}

func NewPool(Nthreads int, Nprocs int) *Pool {
	wg := new(sync.WaitGroup)
	pipe := make(chan Proc, Nprocs)
	//start N goroutines waiting for Procs
	for i := 0; i < Nthreads; i++ {
		wg.Add(1)
		go listen(pipe, wg)
	}
	return &Pool{wg, pipe}
}

func (p *Pool) Send(f Proc) {
	p.pipe <- f
}

func (p *Pool) WaitAll() {
	close(p.pipe)
	p.wg.Wait()
}
