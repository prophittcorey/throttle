# Throttle

[![Go Reference](https://pkg.go.dev/badge/github.com/prophittcorey/throttle.svg)](https://pkg.go.dev/github.com/prophittcorey/throttle)

A golang package that gives you the ability to throttle arbitrary concurrent
functions around a limited resource with minimal blocking.

## Package Usage

Here's a complete program that gives an overview of how to use the throttle.

```golang
package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/prophittcorey/throttle"
)

/* imagine something CPU or IO intensive like network or database requests */
func expensivetask() {
	time.Sleep(500 * time.Millisecond)
}

func main() {
	executor := throttle.New(runtime.NumCPU())

	start := time.Now()

	for i := 0; i <= 50; i++ {
		executor.Run(expensivetask)
	}

	executor.Wait()

	/* example: 12 cores x 50 tasks at 500ms = ~2.5s */
	fmt.Printf("Finished in %s\n", time.Since(start))
}
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
