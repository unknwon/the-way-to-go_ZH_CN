package main
import "fmt"

func main() {
    slice1 := make([]int, 4)

    slice1[0] = 1
    slice1[0] = 2
    slice1[0] = 3
    slice1[0] = 4

    for ix, value := range slice1 {
        fmt.Printf("Slice at %d is: %d\n", ix, value)
    }
}
