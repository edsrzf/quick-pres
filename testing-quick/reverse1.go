package main

import (
	"testing"
	"testing/quick"
)

// START OMIT
func reverseCompare(x, y string) bool {
	for i := 0; i < len(x); i++ {
		if x[len(x)-i-1] != y[i] {
			return false
		}
	}
	return true
}

func TestReverseCompare(t *testing.T) {
	f := func(s string) bool {
		return reverseCompare(reverse(s), s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// END OMIT

func reverse(s string) string {
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[len(s)-i-1] = s[i]
	}
	return string(b)
}

func main() {
	match := func(string, string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{{"TestReverseCompare", TestReverseCompare}}
	testing.Main(match, tests, nil, nil)
}
