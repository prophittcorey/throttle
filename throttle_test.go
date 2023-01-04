package throttle

import (
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	executor := New(100)

	start := time.Now()

	for i := 1; i < 1000; i++ {
		executor.Run(func() {
			time.Sleep(1 * time.Second)
		})
	}

	executor.Wait()

	took := time.Since(start)

	if took < (10*time.Second) || took > (11*time.Second) {
		t.Fatalf("failed to execute in ~10 seconds; took %v", took)
	}
}
