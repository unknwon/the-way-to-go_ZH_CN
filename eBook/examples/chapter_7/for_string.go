package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
)

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}

	c :=
		//		[]byte(s)
		//		[]int(s)
		[]rune(s)
	//	print(c)
	for i, v := range c {
		fmt.Printf("%d=%s", i, v)
	}

	s = "hello"
	b := []byte(s)
	println(&b)
	b[0] = 'c'
	println(&b)
	s2 := string(b) // s2 == "cello"
	print(s2)
	print(&s)
	println(&s2)

	println(Compare([]byte(s), []byte(s2)))

	is := []int{7, 3, 5, 3, 4, -4, 2}
	fmt.Printf("%s\n", is)
	sort.Ints(is)
	fmt.Printf("%s\n", is)
	var idx int
	idx = sort.SearchInts(is, 3)
	println(idx)
	idx = sort.SearchInts(is, 3)
	println(idx)

	fmt.Printf("%c\n", FindDigits("/etc/passwd"))

	var x = 3
	println(x / 2)
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 数组相等
}

var digitRegexp = regexp.MustCompile("[0-9]+")

//func FindDigits(filename string) []byte {
//	b, _ := ioutil.ReadFile(filename)
//	return digitRegexp.Find(b)
//}

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// prints: 0:ÿ 2:界
