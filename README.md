# Slice

Exploring generics in Go.

```go
package main

import (
	"fmt"
	"github.com/nwjlyons/slice"
)

func main() {
	max := slice.Max([]int{6, 4, 8, 2, 1, 9, 4, 7, 5})
	fmt.Printf("Maximum: %v\n", max)
}
```

```shell
go run main.go
Maximum: 9
```