// polartocartesian.go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type polar struct {
	radius float64
	Θ      float64
}

type cartesian struct {
	x float64
	y float64
}

const result = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			Θ := polarCoord.Θ * math.Pi / 180.0 // degrees to radians
			x := polarCoord.radius * math.Cos(Θ)
			y := polarCoord.radius * math.Sin(Θ)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1] // chop of newline character
		if numbers := strings.Fields(line); len(numbers) == 2 {
			polars, err := floatsForStrings(numbers)
			if err != nil {
				fmt.Fprintln(os.Stderr, "invalid number")
				continue
			}
			questions <- polar{polars[0], polars[1]}
			coord := <-answers
			fmt.Printf(result, polars[0], polars[1], coord.x, coord.y)
		} else {
			fmt.Fprintln(os.Stderr, "invalid input")
		}
	}
	fmt.Println()
}

func floatsForStrings(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}
	return floats, nil
}

/* Output:
Enter a radius and an angle (in degrees), e.g., 12.5 90, or Ctrl+Z, Enter to qui
t.
Radius and angle: 12.5 90
Polar: radius=12.50 angle=90.00 degrees -- Cartesian: x=0.00 y=12.50
Radius and angle: ^Z
*/
