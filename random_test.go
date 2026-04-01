package stringx

import (
	"slices"
	"strings"
	"sync"
	"testing"
	"unicode"
)

func TestRandom(t *testing.T) {
	if got := Random(0); got != "" {
		t.Fatalf("Random(0) = %q, want empty string", got)
	}

	cases := []struct {
		name     string
		length   int
		chartype string
		check    func(rune) bool
	}{
		{
			name:     "lowercase",
			length:   32,
			chartype: "l",
			check:    unicode.IsLower,
		},
		{
			name:     "uppercase",
			length:   32,
			chartype: "u",
			check:    unicode.IsUpper,
		},
		{
			name:     "numeric",
			length:   32,
			chartype: "n",
			check:    unicode.IsDigit,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Random(tc.length, tc.chartype)
			if len(got) != tc.length {
				t.Fatalf("Random(%d, %q) length = %d, want %d", tc.length, tc.chartype, len(got), tc.length)
			}

			for _, r := range got {
				if !tc.check(r) {
					t.Fatalf("Random(%d, %q) produced unexpected rune %q", tc.length, tc.chartype, r)
				}
			}
		})
	}
}

func TestSecRandom(t *testing.T) {
	for _, length := range []int{0, 1, 8, 32} {
		got, err := SecRandom(length)
		if err != nil {
			t.Fatalf("SecRandom(%d) returned error: %v", length, err)
		}

		if len(got) != length {
			t.Fatalf("SecRandom(%d) length = %d, want %d", length, len(got), length)
		}
	}
}

func TestRandomNAndRandn(t *testing.T) {
	if got := RandomN(24); len(got) != 24 {
		t.Fatalf("RandomN(24) length = %d, want 24", len(got))
	}

	if got := Randn(12); len(got) != 12 {
		t.Fatalf("Randn(12) length = %d, want 12", len(got))
	}

	for _, got := range []string{RandomN(24), Randn(12)} {
		for _, r := range got {
			if !unicode.IsDigit(r) {
				t.Fatalf("numeric random function produced non-digit rune %q", r)
			}
		}
	}
}

func TestRandStrHelpers(t *testing.T) {
	if got := RandStr(16); len(got) != 16 {
		t.Fatalf("RandStr(16) length = %d, want 16", len(got))
	}

	if got := RandStrUpper(16); len(got) != 16 {
		t.Fatalf("RandStrUpper(16) length = %d, want 16", len(got))
	}
}

func TestRandID(t *testing.T) {
	got := RandId()
	if len(got) != 16 {
		t.Fatalf("RandId() length = %d, want 16", len(got))
	}

	for _, r := range got {
		if !strings.ContainsRune("0123456789abcdef", r) {
			t.Fatalf("RandId() produced non-hex rune %q", r)
		}
	}
}

func TestRandomEle(t *testing.T) {
	if got := RandomEle([]string(nil)); got != "" {
		t.Fatalf("RandomEle(nil) = %q, want empty string", got)
	}

	values := []string{"apple", "banana", "orange"}
	got := RandomEle(values)
	if !slices.Contains(values, got) {
		t.Fatalf("RandomEle(%#v) = %q, want one of the input values", values, got)
	}
}

func TestRandomFunctionsConcurrent(t *testing.T) {
	var wg sync.WaitGroup

	for range 32 {
		wg.Go(func() {
			_ = Random(32)
			_ = RandomN(32)
			_ = RandStr(32)
			_ = RandStrUpper(32)
			_ = RandId()
		})
	}

	wg.Wait()
}
