package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"strings"
)

func main() {
	var conn net.Conn
	var error error
	var inputReader *bufio.Reader
	var input string
	var clientName string

	// maak connectie met de server:
	conn, error = net.Dial("tcp", "localhost:50000")
	checkError(error)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ = inputReader.ReadString('\n')
	// fmt.Printf("CLIENTNAME %s",clientName)
	trimmedClient := strings.Trim(clientName, "\r\n") // "\r\n" voor Windows, "\n" voor Linux
		
	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to shutdown server.")
		input, _ = inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")   
		// fmt.Printf("input:--%s--",input)
		// fmt.Printf("trimmedInput:--%s--",trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, error = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		checkError(error)
	}
}

func checkError(error error) {
	if error != nil {
		panic("Error: " + error.Error())  // terminate program
	}	
}