package main

// START OMIT
import (
	"testing"
	"unicode/utf8"

	"github.com/edsrzf/quick" // HL
)

func TestValidUTF8(t *testing.T) {
	f := func(s string) bool {
		return utf8.ValidString(reverse(s))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
// END OMIT

func reverse(s string) string {
	b := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b[len(s) - i - 1] = s[i]
	}
	return string(b)
}

func main() {
	match := func(string, string) (bool, error) { return true, nil }
	tests := []testing.InternalTest{{"TestValidUTF8", TestValidUTF8}}
	testing.Main(match, tests, nil, nil)
}
