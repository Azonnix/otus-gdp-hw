package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		out = stage(out)
	}

	result := make(Bi)

	go func() {
		defer close(result)

		for {
			select {
			case <-done:
				return
			case pipelineOut, ok := <-out:
				if !ok {
					return
				}

				result <- pipelineOut
			}
		}
	}()

	return result
}

// var wg sync.WaitGroup
// var cc []<-chan interface{}

// for i := range in {
// 	i := i
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		ic := make(chan interface{}, 1)
// 		ic <- i
// 		close(ic)
// 		inc := stages[0](ic)
// 		for _, stage := range stages[1:] {
// 			inc = stage(inc)
// 		}
// 		cc = append(cc, inc)
// 	}()
// }

// wg.Wait()

// return merge(cc...)
// }

// func merge(cs ...In) Out {
// 	var wg sync.WaitGroup
// 	out := make(chan interface{})

// 	output := func(c In) {
// 		for n := range c {
// 			out <- n
// 		}
// 		wg.Done()
// 	}

// 	wg.Add(len(cs))

// 	for _, c := range cs {
// 		go output(c)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	return out
// }
