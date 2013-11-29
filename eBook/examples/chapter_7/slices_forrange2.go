package main
import "fmt"

func main() {
    seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
    for ix, season := range seasons {
        fmt.Printf("Season %d is: %s\n", ix, season)
    }

    var season string
    for _, season = range seasons {
        fmt.Printf("%s\n", season)
    }
}
