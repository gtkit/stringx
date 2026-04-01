package stringx

import (
	"slices"
	"testing"
)

func TestPartition(t *testing.T) {
	head, match, tail := Partition("hello", "l")
	if head != "he" || match != "l" || tail != "lo" {
		t.Fatalf("Partition(%q, %q) = (%q, %q, %q)", "hello", "l", head, match, tail)
	}

	head, match, tail = Partition("hello", "x")
	if head != "hello" || match != "" || tail != "" {
		t.Fatalf("Partition(%q, %q) = (%q, %q, %q)", "hello", "x", head, match, tail)
	}
}

func TestLastPartition(t *testing.T) {
	head, match, tail := LastPartition("hello", "l")
	if head != "hel" || match != "l" || tail != "o" {
		t.Fatalf("LastPartition(%q, %q) = (%q, %q, %q)", "hello", "l", head, match, tail)
	}

	head, match, tail = LastPartition("hello", "x")
	if head != "" || match != "" || tail != "hello" {
		t.Fatalf("LastPartition(%q, %q) = (%q, %q, %q)", "hello", "x", head, match, tail)
	}
}

func TestInsert(t *testing.T) {
	if got := Insert("你好世界", "Go", 2); got != "你好Go世界" {
		t.Fatalf("Insert(%q, %q, 2) = %q, want %q", "你好世界", "Go", got, "你好Go世界")
	}
}

func TestWordSplitNil(t *testing.T) {
	if got := WordSplit("12345"); got != nil {
		t.Fatalf("WordSplit(%q) = %#v, want nil", "12345", got)
	}

	want := []string{"alpha", "beta"}
	if got := WordSplit("alpha beta"); !slices.Equal(got, want) {
		t.Fatalf("WordSplit(%q) = %#v, want %#v", "alpha beta", got, want)
	}
}
