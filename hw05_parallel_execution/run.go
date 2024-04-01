package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	lt := len(tasks)
	if lt < n {
		n = lt
	}
	done := make(chan interface{})
	c := tasksCh(done, tasks)

	errors := workers(done, c, n)
	answer := calculator(done, errors, m)

	if answer <= 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func tasksCh(done chan interface{}, tasks []Task) <-chan Task {
	c := make(chan Task)
	go func() {
		defer close(c)
		for _, t := range tasks {
			select {
			case c <- t:
			case <-done:
				return
			}
		}
	}()
	return c
}

func workers(done chan interface{}, in <-chan Task, n int) <-chan error {
	c := make(chan error)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range in {
				select {
				case c <- t():
				case <-done:
					return
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func calculator(done chan interface{}, err <-chan error, m int) int {
	for e := range err {
		if e != nil && m > 0 {
			m--
			if m <= 0 {
				close(done)
			}
		}
	}
	return m
}
