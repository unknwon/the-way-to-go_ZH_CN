package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		// No connection could be made because the target machine actively refused it.
		fmt.Println("Error dialing", err.Error())
		return // terminate program
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	// fmt.Printf("CLIENTNAME %s",clientName)
	trimmedClient := strings.Trim(clientName, "\r\n") // "\r\n" on Windows, "\n" on Linux
		
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")   
		// fmt.Printf("input:--%s--",input)
		// fmt.Printf("trimmedInput:--%s--",trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
