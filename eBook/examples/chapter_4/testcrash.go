// testcrash.go
// compiles , but crashes
package main

func main() {
	var p *int = nil
	*p = 0

}

// in Windows: stops only with: <exit code="-1073741819" msg="process crashed"/>
// runtime error: invalid memory address or nil pointer dereference
