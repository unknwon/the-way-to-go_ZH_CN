package main

// WARNING! DEAD LOOP!

import "fmt"

type TT float64

//var i int

func (t TT) String() string {
print(".")
    return fmt.Sprintf("%v", t) // dead loop
//return "x"
}

func main(){
    fmt.Println("hi")
    //t := TT(1.2)
    var t TT = 1.2
    fmt.Println(t)
}
