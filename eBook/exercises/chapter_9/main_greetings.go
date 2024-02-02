package main

import (
	"./greetings"
	"fmt"
)

func main() {
	name := "James"
	fmt.Println(greetings.GoodDay(name))
	fmt.Println(greetings.GoodNight(name))

	if greetings.IsAM() {
		fmt.Println("Good morning", name)
	} else if greetings.IsAfternoon() {
		fmt.Println("Good afternoon", name)
	} else if greetings.IsEvening() {
		fmt.Println("Good evening", name)
	} else {
		fmt.Println("Good night", name)
	}
}
