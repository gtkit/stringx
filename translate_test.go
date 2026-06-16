package stringx

import "testing"

func TestTranslate(t *testing.T) {
	if got := Translate("hello", "aeiou", "12345"); got != "h2ll4" {
		t.Fatalf("Translate(%q, %q, %q) = %q, want %q", "hello", "aeiou", "12345", got, "h2ll4")
	}

	if got := Translate("a b", "a", "\x00"); got != "\x00 b" {
		t.Fatalf("Translate should support ASCII to NUL mapping, got %q", got)
	}
}

func TestDelete(t *testing.T) {
	if got := Delete("hello", "aeiou"); got != "hll" {
		t.Fatalf("Delete(%q, %q) = %q, want %q", "hello", "aeiou", got, "hll")
	}
}

func TestCount(t *testing.T) {
	if got := Count("hello", "aeiou"); got != 2 {
		t.Fatalf("Count(%q, %q) = %d, want 2", "hello", "aeiou", got)
	}
}

func TestSqueeze(t *testing.T) {
	if got := Squeeze("hello   world", " "); got != "hello world" {
		t.Fatalf("Squeeze(%q, %q) = %q, want %q", "hello   world", " ", got, "hello world")
	}
}
