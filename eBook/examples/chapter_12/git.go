package main

// fake git exe to intercept args to real git.exe

import (
	"flag" // command line option parser
	"log"
	"os"
)

var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.Parse() // Scans the arg list and sets up flags
	setLogFile()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
	log.Println(s)
}

func setLogFile() {
	f, err := os.OpenFile("t:/git.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		print("error opening file: ", err)
	}
	//	defer f.Close()

	log.SetOutput(f)
	log.Println("This is a test log entry")
}
