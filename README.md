# safeslices

A small Go utility package that provides safe slice operations.

## 🚀 Features

- Safe maximum value for slices (`MaxSafe`)
- Prevents panic on empty slices
- Returns a boolean flag to indicate success

## 📦 Example

```go
package main

import (
    "fmt"
    "github.com/vinith-raj-16/safeslices"
)

func main() {
    max, ok := safeslices.MaxSafe([]int{1, 5, 3})

    if ok {
        fmt.Println("Max:", max)
    } else {
        fmt.Println("Empty slice")
    }
}