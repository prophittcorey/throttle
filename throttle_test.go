package throttle

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	executor := New(1000)

	start := time.Now()

	var processed int32

	for i := 1; i <= 10_000; i++ {
		executor.Run(func() {
			defer func() {
				atomic.AddInt32(&processed, 1)
			}()
			time.Sleep(500 * time.Millisecond)
		})
	}

	executor.Wait()

	took := time.Since(start)

	/* ensure 10K functions were executed */
	if processed != 10_000 {
		t.Fatalf("failed to process 10,000 integers; %v", processed)
	}

	/* ensure timing is acceptable: (10,000 * 500ms) / 1000 = ~5s */
	if took < (5*time.Second) || took > (6*time.Second) {
		t.Fatalf("failed to execute in ~5 seconds; took %v", took)
	}
}
