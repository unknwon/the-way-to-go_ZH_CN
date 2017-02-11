// httpfetch.go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Print("Give the url from which to read: ")
	iread := bufio.NewReader(os.Stdin)
	url, _ := iread.ReadString('\n')
	url = strings.Trim(url, " \n\r") // trimming space,etc.
	// fmt.Println("***", url,"***") // debugging
	res, err := http.Get(url)
	CheckError(err)
	data, err := ioutil.ReadAll(res.Body)
	CheckError(err)
	fmt.Printf("Got: %q", string(data))
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Get: %v", err)
	}
}
