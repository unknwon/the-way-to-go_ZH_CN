// float64 is necessary as input to math.Sqrt()
package main

import (
	"fmt"
	"math"
)

type Magnitude interface {
	Abs() float64
}

var m Magnitude

type Point struct {
	X, Y float64
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

type Point3 struct {
	X, Y, Z float64
}

func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

type Polar struct {
	R, Theta float64
}

func (p Polar) Abs() float64 { return p.R }

func main() {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	m = p1 // p1 is type *Point, has method Abs()
	fmt.Printf("The length of the vector p1 is: %f\n", m.Abs())

	p2 := &Point{4, 5}
	m = p2
	fmt.Printf("The length of the vector p2 is: %f\n", m.Abs())

	p1.Scale(5)
	m = p1
	fmt.Printf("The length of the vector p1 after scaling is: %f\n", m.Abs())
	fmt.Printf("Point p1 after scaling has the following coordinates: X %f - Y %f\n", p1.X, p1.Y)

	mag := m.Abs()
	m = &Point3{3, 4, 5}
	mag += m.Abs()
	m = Polar{2.0, math.Pi / 2}
	mag += m.Abs()
	fmt.Printf("The float64 mag is now: %f", mag)
}

/* Output:
The length of the vector p1 is: 5.000000
The length of the vector p2 is: 6.403124
The length of the vector p1 after scaling is: 25.000000
Point p1 after scaling has the following coordinates: X 15.000000 - Y 20.000000
The float64 mag is now: 34.071068

-- instead of:
func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
m = p1
we can write:
func (p Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
m = p1
and instead of:
func (p Polar) Abs() float64 { return p.R }
m = Polar{2.0, math.Pi / 2}
we can write:
func (p *Polar) Abs() float64 { return p.R }
m = &Polar{2.0, math.Pi / 2}
with the same output
*/
