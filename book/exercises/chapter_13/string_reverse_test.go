// string_reverse_test.go
package strev

import "testing"
import "./strev"

type ReverseTest struct {
	in, out 	string
}

var ReverseTests = []ReverseTest {
	ReverseTest{"ABCD", "DCBA"},
	ReverseTest{"CVO-AZ", "ZA-OVC"},
	ReverseTest{"Hello 世界", "界世 olleH"},
}

func TestReverse(t *testing.T) {
/*
	in := "CVO-AZ"
	out := Reverse(in)
	exp := "ZA-OVC"
	if out != exp {
		t.Errorf("Reverse of %s expects %s, but got %s", in, exp, out)
	}
*/
// testing with a battery of testdata:
	for _, r := range ReverseTests {
		exp := strev.Reverse(r.in)
		if r.out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", r.in, exp, r.out)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "ABCD"
	for i:=0; i < b.N; i++ {
		strev.Reverse(s)
	}
}
