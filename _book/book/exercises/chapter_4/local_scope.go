package main

var a = "G"   // global (package) scope

func main() {
	n()
	m()
	n()
}
func n() {
	print(a)
}
func m() {
	a := "O"  // new local variable a is declared
	print(a)
}
// GOG