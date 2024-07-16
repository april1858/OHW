package hw06pipelineexecution

import (
	"sort"
	"sync"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

type Data struct {
	val  interface{}
	turn int
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	sd := make([]Data, 0)
	data := Splitter(done, in, stages...)
	for d := range data {
		sd = append(sd, d)
	}
	sort.Slice(sd, func(i, j int) bool {
		return sd[i].turn < sd[j].turn
	})

	out := genAnswer(sd)

	return out
}

func genAnswer(sd []Data) In {
	out := make(Bi)
	go func() {
		for _, v := range sd {
			out <- v.val
		}
		close(out)
	}()

	return out
}

func Splitter(done In, in In, stages ...Stage) <-chan Data {
	out := make(chan Data)
	turn := -1
	var wg sync.WaitGroup
	for i := range in {
		turn++
		wg.Add(1)
		go func(done In, i interface{}, q int) {
			msg := Data{turn: q}
			defer wg.Done()
			in := g(i)
			for i := 0; i < len(stages); i++ {
				select {
				case <-done:
					return
				default:
					in = stages[i](in)
					time.Sleep(time.Millisecond * 75)
				}
			}
			msg.val = <-in

			out <- msg
		}(done, i, turn)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func g(in interface{}) In {
	c := make(Bi)
	go func() {
		c <- in
	}()
	return c
}
