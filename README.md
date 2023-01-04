# Throttle

A golang package that gives you the ability to throttle concurrent functions
around a limited resource.

## Package Usage

```golang

import (
  "fmt"
  "time"

  "github.com/prophittcorey/throttle"
)

func main() {
	executor := throttle.New(runtime.NumCPU())

	start := time.Now()

	for i := 1; i < 100; i++ {
		executor.Run(func() {
			time.Sleep(1 * time.Second)
		})
	}

	executor.Wait()

	fmt.Printf("Finished in %s\n", time.Since(start))
}

```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
