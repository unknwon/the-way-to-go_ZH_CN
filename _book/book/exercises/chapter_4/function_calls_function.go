package main

var a string  // global scope

func main() {
	a = "G"
	print(a)
	f1()
}
func f1() {
	a := "O"   // new local variable a, only scoped within f1() !
	print(a)
	f2()
}
func f2() {
	print(a)   // global variable is taken
}
// GOG