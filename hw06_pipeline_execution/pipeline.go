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
	// Place your code here.
	return finish(done, stg3(done, stg2(done, stg1(done, stg0(done, gen(done, in), stages[0]), stages[1]), stages[2]), stages[3]))
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

func stg0(done In, in In, stage Stage) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for i := range stage(in) {
			select {
			case <-done:
				return
			default:
			}
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

func stg1(done In, in In, stage Stage) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for i := range stage(in) {
			select {
			case <-done:
				return
			default:
			}
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

func stg2(done In, in In, stage Stage) Out {
	out := make(Bi)
	go func() {
		defer close(out)

		for i := range stage(in) {
			select {
			case <-done:
				return
			default:
			}
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

func stg3(done In, in In, stage Stage) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for i := range stage(in) {
			select {
			case <-done:
				return
			default:
			}
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}

func finish(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-done:
				return
			default:
			}
			select {
			case <-done:
				return
			case out <- i:
			}
		}
	}()
	return out
}
