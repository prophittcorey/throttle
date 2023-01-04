// Package throttle implements a data structure to help limit concurrency
// around a fixed resource limit without excessive blocking.
package throttle

import (
	"log"

	"sync"
)

// Executor is the monitor for executing functions. The executor will ensure
// only a fixed number of functions are fun concurrently and will block as
// needed.
type Executor struct {
	sync.WaitGroup
	sem chan struct{}
}

// Run executes a function and ensures concurrency is throttled.
func (q *Executor) Run(f func()) {
	q.sem <- struct{}{}
	q.Add(1)

	go (func() {
		defer (func() {
			if err := recover(); err != nil {
				log.Printf("recovered from a panic; %s\n", err)
			}

			q.Done()
			<-q.sem
		})()

		f()
	})()
}

// New creates a new Executor that is ready to be used.
func New(resourcelimit int) *Executor {
	if resourcelimit < 1 {
		resourcelimit = 1
	}

	return &Executor{
		sem: make(chan struct{}, resourcelimit),
	}
}
