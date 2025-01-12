package caches

import (
	"sync"
)

func ease(t task, queue *sync.Map) task {
	eq := &eased{
		task: t,
		wg:   &sync.WaitGroup{},
	}
	eq.wg.Add(1)
	defer eq.wg.Done()

	runner, ok := queue.LoadOrStore(t.GetId(), eq)
	if ok {
		et := runner.(*eased)
		et.wg.Wait()

		return et.task
	}

	eq.task.Run()
	queue.Delete(t.GetId())
	return eq.task
}

type eased struct {
	task task
	wg   *sync.WaitGroup
}
