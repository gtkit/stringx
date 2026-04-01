package stringx

import (
	"slices"
	"testing"
)

func TestLenAndWidth(t *testing.T) {
	if got := Len("你好"); got != 2 {
		t.Fatalf("Len(%q) = %d, want 2", "你好", got)
	}

	if got := RuneWidth('你'); got != 2 {
		t.Fatalf("RuneWidth(%q) = %d, want 2", '你', got)
	}

	if got := Width("A你好"); got != 5 {
		t.Fatalf("Width(%q) = %d, want 5", "A你好", got)
	}
}

func TestWordCountAndSplit(t *testing.T) {
	if got := WordCount("hello, world"); got != 2 {
		t.Fatalf("WordCount(%q) = %d, want 2", "hello, world", got)
	}

	if got := WordCount("don't-stop now"); got != 2 {
		t.Fatalf("WordCount(%q) = %d, want 2", "don't-stop now", got)
	}

	want := []string{"hello", "world", "don't-stop"}
	if got := WordSplit("hello, world, don't-stop"); !slices.Equal(got, want) {
		t.Fatalf("WordSplit returned %#v, want %#v", got, want)
	}
}
