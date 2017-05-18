package main

var a = "G" // global scope

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O" // simple assignment: global a gets a new value
	print(a)
}

// GOO
