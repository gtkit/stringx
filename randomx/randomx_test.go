package randomx

import (
	"strings"
	"testing"
	"unicode"
)

func TestRandomHelpers(t *testing.T) {
	if got := Random(8); len(got) != 8 {
		t.Fatalf("Random length = %d, want 8", len(got))
	}

	if got := RandomN(6); len(got) != 6 {
		t.Fatalf("RandomN length = %d, want 6", len(got))
	}

	if got := RandStr(6); len(got) != 6 {
		t.Fatalf("RandStr length = %d, want 6", len(got))
	}

	if got := RandStrUpper(6); len(got) != 6 {
		t.Fatalf("RandStrUpper length = %d, want 6", len(got))
	}

	if got := RandId(); len(got) != 16 {
		t.Fatalf("RandId length = %d, want 16", len(got))
	}

	if got, err := SecRandom(12); err != nil || len(got) != 12 {
		t.Fatalf("SecRandom = %q, %v", got, err)
	}

	if got := RandomFromCharset(8, []byte("xy")); len(got) != 8 {
		t.Fatalf("RandomFromCharset length = %d, want 8", len(got))
	}

	values := []string{"apple", "banana", "orange"}
	got := RandomEle(values)
	if got != "apple" && got != "banana" && got != "orange" {
		t.Fatalf("RandomEle returned unexpected value %q", got)
	}

	for _, r := range Random(16, "n") {
		if !unicode.IsDigit(r) {
			t.Fatalf("Random(n) returned non-digit rune %q", r)
		}
	}

	for _, r := range RandId() {
		if !strings.ContainsRune("0123456789abcdef", r) {
			t.Fatalf("RandId returned non-hex rune %q", r)
		}
	}
}
