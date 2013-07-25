package main

import (
    "fmt"
    "errors"
    "math"
)


func MySqrt1(a float64) (Sqrt float64, Err error) {
    if a < 0 {
        Err = errors.New("Fuck!")
        Sqrt = 0
        return
    }
    Sqrt = math.Sqrt(a)
    return
}

func MySqrt2(a float64) (float64, error) {
    var Sqrt float64
    var Err error
    if a < 0 {
        Err = errors.New("Fuck!")
        Sqrt = 0
        return Sqrt, Err
    }
    Sqrt = math.Sqrt(a)
    return Sqrt, Err
}

func main() {
    var a float64 = 99.99
    var b float64 = -99.99
    Sqrt11, Err11 := MySqrt1(a)
    Sqrt12, Err12 := MySqrt1(b)
    Sqrt21, Err21 := MySqrt1(a)
    Sqrt22, Err22 := MySqrt1(b)

    fmt.Println(Sqrt11, Err11)
    fmt.Println(Sqrt12, Err12)
    fmt.Println(Sqrt21, Err21)
    fmt.Println(Sqrt22, Err22)
}