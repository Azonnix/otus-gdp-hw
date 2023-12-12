package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m == 0 {
		return ErrErrorsLimitExceeded
	}

	var countErr int32
	var once sync.Once
	var runErr error
	wg := sync.WaitGroup{}
	taskc := make(chan Task, len(tasks))
	terminatedc := make(chan struct{}, 1)

	closeTerminatedFunc := func() {
		close(terminatedc)
	}

	for _, task := range tasks {
		taskc <- task
	}
	close(taskc)

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for task := range taskc {
				select {
				case <-terminatedc:
					return
				default:
					if err := task(); err != nil {
						atomic.AddInt32(&countErr, 1)
						if int(atomic.LoadInt32(&countErr)) == m {
							once.Do(closeTerminatedFunc)
							runErr = ErrErrorsLimitExceeded
						}
					}
				}
			}
		}()
	}

	wg.Wait()

	return runErr
}
