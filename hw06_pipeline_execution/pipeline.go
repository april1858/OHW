package hw06pipelineexecution

import (
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Bi)

	tasks := gen(done, in)

	go func() {
		defer close(out)
		for i := 0; i < len(stages); i++ {
			tasks = stages[i](broker(done, tasks))
		}
		for i := range tasks {
			out <- i
		}
	}()
	return out
}

func broker(done In, in In) Out {
	c := make(Bi)
	go func() {
		defer close(c)
		for i := range in {
			select {
			case <-done:
				return
			default:
				c <- i
			}
		}
	}()

	return c
}

func gen(done In, in In) Out {
	time.Sleep(2 * time.Millisecond) // !???
	out := make(Bi)
	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}
