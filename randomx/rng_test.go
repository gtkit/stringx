package randomx_test

import (
	"slices"
	"testing"

	"github.com/gtkit/stringx/randomx"
)

func TestRandIntHandlerOrder(t *testing.T) {
	var indices []int
	randomx.RandIntHandler(1, 3, func(num, i int) {
		if num != 0 {
			t.Errorf("RandIntHandler produced %d, want 0", num)
		}
		indices = append(indices, i)
	})
	if want := []int{2, 1, 0}; !slices.Equal(indices, want) {
		t.Fatalf("RandIntHandler callback indices = %v, want %v", indices, want)
	}
}

func TestRNGBounds(t *testing.T) {
	var rng randomx.RNG
	if got := rng.Uint32n(0); got != 0 {
		t.Fatalf("Uint32n(0) = %d, want 0", got)
	}
	if got := rng.Uint32n(1); got != 0 {
		t.Fatalf("Uint32n(1) = %d, want 0", got)
	}
	if got := rng.Uint64n(0); got != 0 {
		t.Fatalf("Uint64n(0) = %d, want 0", got)
	}
}
