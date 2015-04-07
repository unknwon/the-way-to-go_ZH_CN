// interface_poly3.go
package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Shape struct {}
func (sh Shape) Area() float32 {
	return -1 // the shape is indetermined, so we return something impossible
}

type Square struct {
	side float32
	Shape
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
   length, width float32
   Shape
}

func (r *Rectangle) Area() float32 {
   return r.length * r.width
}

type Circle struct {
	radius float32
	Shape
}

func (c *Circle) Area() float32 {
   return math.Pi * c.radius * c.radius 
}

func main() {
   s := Shape{}
   r := &Rectangle{5, 3, s} // Area() of Rectangle needs a value
   q := &Square{5, s}      // Area() of Square needs a pointer
   c := &Circle{2.5, s}
   shapes := []Shaper{r, q, c, s}
   fmt.Println("Looping through shapes for area ...")
   for n, _ := range shapes {
       fmt.Println("Shape details: ", shapes[n])
       fmt.Println("Area of this shape is: ", shapes[n].Area())
   }
}
/* Output:
Looping through shapes for area ...
Shape details:  {5 3}
Area of this shape is:  15
Shape details:  &{5}
Area of this shape is:  25
Shape details:  &{2.5}
Area of this shape is:  19.634954
Shape details:  {}
Area of this shape is:  -1
*/

