package hw05parallelexecution

import (
	"errors"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m == 0 {
		return ErrErrorsLimitExceeded
	}

	taskChan := make(chan Task)
	doneChan := make(chan struct{})
	errChan := make(chan error)
	runErrChan := make(chan error)
	quitChan := make(chan struct{})
	quitWriteChan := make(chan struct{}, 1)
	var runErr error

	go RunThreads(taskChan, errChan, doneChan, n)
	go RunCountErr(runErrChan, errChan, quitChan, m)
	go RunWriteTasks(tasks, taskChan, quitWriteChan)

	select {
	case err := <-runErrChan:
		quitWriteChan <- struct{}{}
		<-doneChan
		quitChan <- struct{}{}
		runErr = err
	case <-doneChan:
		quitChan <- struct{}{}
		quitWriteChan <- struct{}{}
	}

	return runErr
}

func RunThreads(taskChan chan Task, errChan chan error, doneChan chan struct{}, n int) {
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func(taskChan <-chan Task, errChan chan error) {
			defer wg.Done()

			for task := range taskChan {
				err := task()
				if err != nil {
					errChan <- err
				}
				time.Sleep(1 * time.Nanosecond)
			}
		}(taskChan, errChan)
	}
	wg.Wait()
	doneChan <- struct{}{}
}

func RunCountErr(runErrChan chan error, errChan chan error, quitChan chan struct{}, m int) {
	countErr := 0

	for {
		select {
		case <-quitChan:
			return
		case <-errChan:
			countErr++

			if countErr == m {
				runErrChan <- ErrErrorsLimitExceeded
			}
		}
	}
}

func RunWriteTasks(tasks []Task, taskChan chan Task, quitWriteChan chan struct{}) {
	for _, task := range tasks {
		select {
		case <-quitWriteChan:
			close(taskChan)
			return
		default:
			taskChan <- task
		}
	}
	close(taskChan)
}

// package hw05parallelexecution

// import (
// 	"errors"
// 	"fmt"
// 	"sync"
// )

// var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

// type Task func() error

// // Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
// func Run(tasks []Task, n, m int) error {
// 	taskChan := make(chan Task)
// 	doneChan := make(chan struct{})
// 	errChan := make(chan error)
// 	runErrChan := make(chan error)
// 	quitChan := make(chan struct{}, 1)
// 	quitTaskChan := make(chan struct{}, 1)
// 	var runErr error

// 	go RunThreads(taskChan, errChan, doneChan, quitTaskChan, n)
// 	go RunCountErr(runErrChan, errChan, quitChan, m)

// 	for _, task := range tasks {
// 		taskChan <- task
// 	}
// 	close(taskChan)

// 	// for {
// 	// 	select {
// 	// 	case err := <-runErrChan:
// 	// 		runErr = err
// 	// 		quitChan <- struct{}{}
// 	// 		close(taskChan)
// 	// 		quitTaskChan <- struct{}{}
// 	// 		<-doneChan
// 	// 		return runErr
// 	// 	default:
// 	// 		if len(tasks) == 0 {
// 	// 			quitChan <- struct{}{}
// 	// 			close(taskChan)
// 	// 			<-doneChan
// 	// 			return nil
// 	// 		}
// 	// 		task := tasks[0]
// 	// 		tasks = tasks[1:]
// 	// 		taskChan <- task
// 	// 	}
// 	// }

// 	select {
// 	case err := <-runErrChan:
// 		quitChan <- struct{}{}
// 		quitTaskChan <- struct{}{}
// 		runErr = err
// 		fmt.Println("noooooooooooooooo")
// 	case <-doneChan:
// 		quitChan <- struct{}{}
// 		fmt.Println("eeeeeeeeeeeeeee")
// 	}

// 	return runErr
// }

// func RunThreads(taskChan chan Task, errChan chan error, doneChan chan struct{}, quitChan chan struct{}, n int) {
// 	wg := sync.WaitGroup{}

// 	for i := 0; i < n; i++ {
// 		wg.Add(1)

// 		go func(taskChan <-chan Task, errChan chan error) {
// 			defer wg.Done()

// 			for task := range taskChan {
// 				select {
// 				case <-quitChan:
// 					fmt.Println("quit")
// 					return
// 				default:
// 					err := task()
// 					if err != nil {
// 						errChan <- err
// 					}
// 				}

// 				// err := task()
// 				// if err != nil {
// 				// 	errChan <- err
// 				// }
// 			}
// 		}(taskChan, errChan)
// 	}

// 	wg.Wait()
// 	doneChan <- struct{}{}
// }

// func RunCountErr(runErrChan chan error, errChan chan error, quitChan chan struct{}, m int) {
// 	countErr := 0

// 	for {
// 		select {
// 		case <-quitChan:
// 			fmt.Println("quit err")
// 			return
// 		case <-errChan:
// 			countErr++
// 			fmt.Println(countErr)

// 			if countErr == m {
// 				fmt.Println("wow")
// 				runErrChan <- ErrErrorsLimitExceeded
// 			}
// 		}
// 	}
// }
