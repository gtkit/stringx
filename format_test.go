package stringx

import "testing"

func TestJustifyAndCenter(t *testing.T) {
	if got := LeftJustify("hello", 10, " "); got != "hello     " {
		t.Fatalf("LeftJustify(%q, 10, %q) = %q, want %q", "hello", " ", got, "hello     ")
	}

	if got := RightJustify("hello", 10, " "); got != "     hello" {
		t.Fatalf("RightJustify(%q, 10, %q) = %q, want %q", "hello", " ", got, "     hello")
	}

	if got := Center("hello", 10, " "); got != "  hello   " {
		t.Fatalf("Center(%q, 10, %q) = %q, want %q", "hello", " ", got, "  hello   ")
	}
}
