package uc

import "testing"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	{"abc", "ABC"},
	{"cvo-az", "CVO-AZ"},
	{"Antwerp", "ANTWERP"},
}

func TestUC(t *testing.T) {
	for _, ut := range ucTests {
		uc := UpperCase(ut.in)
		if uc != ut.out {
			t.Errorf("UpperCase(%s) = %s, must be %s", ut.in, uc,
				ut.out)
		}
	}
}
