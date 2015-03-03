// main_stack_v2.go
package main

import (
    "./stack/collection"
    "fmt"
)

func main() {
    var s collection.Stack
    s.Push("world")
    s.Push("hello, ")
    for s.Size() > 0 {
        fmt.Print(s.Pop())
    }
    fmt.Println()
}