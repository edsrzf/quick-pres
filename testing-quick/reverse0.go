package main

import (
	"testing"
	"testing/quick"
)

// START OMIT
func reverse(s string) string {
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[len(s)-i-1] = s[i]
	}
	return string(b)
}

func TestDoubleReverse(t *testing.T) {
	f := func(s string) bool { return reverse(reverse(s)) == s }
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestReverseConcat(t *testing.T) {
	f := func(x, y string) bool {
		return reverse(x)+reverse(y) == reverse(y+x)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// END OMIT

func main() {
	match := func(string, string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{{"TestDoubleReverse", TestDoubleReverse}, {"TestReverseConcat", TestReverseConcat}}
	testing.Main(match, tests, nil, nil)
}
