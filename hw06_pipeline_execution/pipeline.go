package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	out := make(Bi)
	go func() {
		defer close(out)
		for i := 0; i < len(stages); i++ {
			in = stages[i](broker(done, in))
		}
		for i := range in {
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
