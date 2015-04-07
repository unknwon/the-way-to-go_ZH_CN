package main
<<<<<<< HEAD

=======
>>>>>>> 0364d180eb1d1b8067d1248ace49ca46c816e541
import (
	"fmt"
	"os"
)
<<<<<<< HEAD

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")  
    fmt.Printf("Path is %s\n", path)
}
=======
func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
>>>>>>> 0364d180eb1d1b8067d1248ace49ca46c816e541
