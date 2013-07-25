package main

import "fmt"

func MultReturn1(a, b int) (Sum, Mul, Dif int) {
    Sum = a + b
    Mul = a * b
    Dif = a - b
    return
}

func MultReturn2(a, b int) (int, int, int) {
    Sum := a + b
    Mul := a * b
    Dif := a - b
    return Sum, Mul, Dif

}

func main() {
    a := 100
    b := 78
    Sum1, Mul1, Dif1 := MultReturn1(a, b)
    Sum2, Mul2, Dif2 := MultReturn2(a, b)
    fmt.Println("Named:", Sum1, Mul1, Dif1)
    fmt.Println("Unnamed:", Sum2, Mul2, Dif2)
}
