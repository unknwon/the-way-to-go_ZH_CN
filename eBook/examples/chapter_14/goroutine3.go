package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
//	getData(ch)
	getData2(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}
// Washington Tripoli London Beijing Tokio 

func getData2(ch chan string) {
//	for input, open := <-ch; open ; { // doesn't work, infinite loop, seems open was checked by "for"
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}
