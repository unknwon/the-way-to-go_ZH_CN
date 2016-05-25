// string_reverse.go
package strev

func Reverse(s string) string {
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

/*
func main() {
	s := "My Test String!"
	fmt.Println(s, " --> ", Reverse(s))
}
*/