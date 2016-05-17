// pattern.go
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
)

func main() {

	inf := new(Info)
	//	println(inf)
	fmt.Printf("info %v\n", inf)
	Update(inf)
	fmt.Printf("info %v\n", inf)

	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" // string to search
	pat := "[0-9]+.[0-9]+"                                      // pattern search for in searchIn

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		fmt.Printf("v %f\n", v)
		return strconv.FormatFloat(v*2, 'f', 3, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}

	re2 := regexp.MustCompile(pat)
	println(re2)

	re, err := regexp.Compile(pat)
	fmt.Printf("err %v\n", err)
	//	println(err)
	err = nil
	str := re.ReplaceAllString(searchIn, "##.#") // replace pat with "##.#"
	fmt.Println(str)
	// using a function :
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

}

type Info struct {
	mu sync.Mutex
	// ... other fields, e.g.:
	Str string
}

func Update(info *Info) {
	info.mu.Lock()
	// critical section:
	info.Str = "ab" // new value
	// end critical section
	info.mu.Unlock()
}

/* Output:
Match found!
John: ##.# William: ##.# Steve: ##.#
John: 5156.68 William: 9134.46 Steve: 11264.36
*/
