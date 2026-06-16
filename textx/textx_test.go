package textx

import "testing"

func TestTextHelpers(t *testing.T) {
	if got := LowerFirst("BigCamelCase"); got != "bigCamelCase" {
		t.Fatalf("LowerFirst returned %q, want %q", got, "bigCamelCase")
	}

	if got := UpperFirst("smallCamelCase"); got != "SmallCamelCase" {
		t.Fatalf("UpperFirst returned %q, want %q", got, "SmallCamelCase")
	}

	if !IsBlank(" \t\n") {
		t.Fatal("IsBlank should accept whitespace-only string")
	}

	if got := NormalizeSpace("  hello \t world \n  golang  "); got != "hello world golang" {
		t.Fatalf("NormalizeSpace returned %q, want %q", got, "hello world golang")
	}

	if got := Substr("你好世界", 1, 2); got != "好世" {
		t.Fatalf("Substr returned %q, want %q", got, "好世")
	}

	if got := SubByte("你好世界", 3); got != "你" {
		t.Fatalf("SubByte returned %q, want %q", got, "你")
	}

	if got := Reverse("hello 世界"); got != "界世 olleh" {
		t.Fatalf("Reverse returned %q, want %q", got, "界世 olleh")
	}

	if got := Escape("a\"b'c"); got != "a\\\"b\\'c" {
		t.Fatalf("Escape returned %q, want %q", got, "a\\\"b\\'c")
	}

	if got := Slug("Laravel 5 Framework", "-"); got != "Laravel-5-Framework" {
		t.Fatalf("Slug returned %q, want %q", got, "Laravel-5-Framework")
	}

	head, match, tail := Partition("hello", "l")
	if head != "he" || match != "l" || tail != "lo" {
		t.Fatalf("Partition returned (%q, %q, %q)", head, match, tail)
	}

	head, match, tail = LastPartition("hello", "l")
	if head != "hel" || match != "l" || tail != "o" {
		t.Fatalf("LastPartition returned (%q, %q, %q)", head, match, tail)
	}

	if got := Insert("你好世界", "Go", 2); got != "你好Go世界" {
		t.Fatalf("Insert returned %q, want %q", got, "你好Go世界")
	}

	if got := Len("你好"); got != 2 {
		t.Fatalf("Len returned %d, want 2", got)
	}

	if got := Width("A你好"); got != 5 {
		t.Fatalf("Width returned %d, want 5", got)
	}

	if got := WordCount("hello, world"); got != 2 {
		t.Fatalf("WordCount returned %d, want 2", got)
	}

	if got := LeftJustify("go", 4, "."); got != "go.." {
		t.Fatalf("LeftJustify returned %q, want %q", got, "go..")
	}

	if got := RightJustify("go", 4, "."); got != "..go" {
		t.Fatalf("RightJustify returned %q, want %q", got, "..go")
	}

	if got := Center("go", 4, "."); got != ".go." {
		t.Fatalf("Center returned %q, want %q", got, ".go.")
	}

	if got := Translate("hello", "aeiou", "12345"); got != "h2ll4" {
		t.Fatalf("Translate returned %q, want %q", got, "h2ll4")
	}

	if got := Delete("hello", "aeiou"); got != "hll" {
		t.Fatalf("Delete returned %q, want %q", got, "hll")
	}

	if got := Count("hello", "aeiou"); got != 2 {
		t.Fatalf("Count returned %d, want 2", got)
	}

	if got := Squeeze("hello   world", " "); got != "hello world" {
		t.Fatalf("Squeeze returned %q, want %q", got, "hello world")
	}
}

func TestNormalizeSpaceFastPathAndUnicode(t *testing.T) {
	if got := NormalizeSpace("hello world"); got != "hello world" {
		t.Fatalf("NormalizeSpace returned %q, want %q", got, "hello world")
	}

	if allocs := testing.AllocsPerRun(1000, func() {
		_ = NormalizeSpace("hello world")
	}); allocs != 0 {
		t.Fatalf("NormalizeSpace normalized fast path allocations = %.0f, want 0", allocs)
	}

	if got := NormalizeSpace("\u3000hello\t\nworld\u00a0"); got != "hello world" {
		t.Fatalf("NormalizeSpace unicode whitespace returned %q, want %q", got, "hello world")
	}
}

func TestSubByteEmojiBoundary(t *testing.T) {
	input := "a🙂b"
	if got := SubByte(input, 2); got != "a" {
		t.Fatalf("SubByte(%q, 2) = %q, want %q", input, got, "a")
	}
	if got := SubByte(input, 5); got != "a🙂" {
		t.Fatalf("SubByte(%q, 5) = %q, want %q", input, got, "a🙂")
	}
}
