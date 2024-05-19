package hw06pipelineexecution

import (
	"sync"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	out := make([]<-chan interface{}, 0)
	dests := Split(done, in, 1) // 1 - if the initial order is needed
	var wg sync.WaitGroup
	for _, ch := range dests {
		wg.Add(1)
		go func(c In) {
			defer wg.Done()
			for _, stage := range stages {
				c = stage(c)
			}
			out = append(out, c)
		}(ch)
	}
	wg.Wait()
	answer := Funnel(out)
	return answer
}

func Split(done In, source <-chan interface{}, n int) []<-chan interface{} {
	if done != nil {
		return nil
	}
	dests := make([]<-chan interface{}, 0)
	for i := 0; i < n; i++ {
		ch := make((chan interface{}))
		dests = append(dests, ch)
		go func() {
			defer close(ch)
			for val := range source {
				ch <- val
			}
		}()
	}
	return dests
}

func Funnel(source []<-chan interface{}) Out {
	dest := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(source))
	for _, ch := range source {
		go func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				dest <- i
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(dest)
	}()
	return dest
}
