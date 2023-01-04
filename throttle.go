package throttle

import (
	"log"

	"sync"
)

type Executor struct {
	sync.WaitGroup
	sem chan struct{}
}

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

func New(resourcelimit int) *Executor {
	return &Executor{
		sem: make(chan struct{}, resourcelimit),
	}
}

// func main() {
// 	executor := throttle.New(runtime.NumCPU())

// 	start := time.Now()

// 	for i := 1; i < 100; i++ {
// 		executor.Run(func() {
// 			time.Sleep(1 * time.Second)
// 		})
// 	}

// 	executor.Wait()

// 	fmt.Printf("Finished in %s\n", time.Since(start))
// }
