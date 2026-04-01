package stringx

import (
	"bytes"
	"slices"
	"testing"
)

func TestFirstCaseHelpers(t *testing.T) {
	if got := LowerFirst("BigCamelCase"); got != "bigCamelCase" {
		t.Fatalf("LowerFirst(%q) = %q, want %q", "BigCamelCase", got, "bigCamelCase")
	}

	if got := UpperFirst("smallCamelCase"); got != "SmallCamelCase" {
		t.Fatalf("UpperFirst(%q) = %q, want %q", "smallCamelCase", got, "SmallCamelCase")
	}

	if got := LowerFirst("Éclair"); got != "éclair" {
		t.Fatalf("LowerFirst(%q) = %q, want %q", "Éclair", got, "éclair")
	}

	if got := UpperFirst("éclair"); got != "Éclair" {
		t.Fatalf("UpperFirst(%q) = %q, want %q", "éclair", got, "Éclair")
	}

	if got := FirstLowerCase("FirstLowerCase"); got != "f" {
		t.Fatalf("FirstLowerCase(%q) = %q, want %q", "FirstLowerCase", got, "f")
	}

	if got := FirstUpperCase("firstUpperCase"); got != "F" {
		t.Fatalf("FirstUpperCase(%q) = %q, want %q", "firstUpperCase", got, "F")
	}
}

func TestBlankHelpers(t *testing.T) {
	if !IsEmpty("") {
		t.Fatal("IsEmpty should return true for empty string")
	}

	if IsEmpty("x") {
		t.Fatal("IsEmpty should return false for non-empty string")
	}

	if !IsNotEmpty("x") {
		t.Fatal("IsNotEmpty should return true for non-empty string")
	}

	if !IsBlank(" \t\n") {
		t.Fatal("IsBlank should return true for whitespace-only string")
	}

	if IsBlank("go") {
		t.Fatal("IsBlank should return false for content string")
	}

	if !IsNotBlank("go") {
		t.Fatal("IsNotBlank should return true for content string")
	}

	if !IsAllBlank("", " ", "\n") {
		t.Fatal("IsAllBlank should return true when all strings are blank")
	}

	if IsAllBlank("", "go") {
		t.Fatal("IsAllBlank should return false when one string is not blank")
	}

	if !IsAllNotBlank("go", "lang") {
		t.Fatal("IsAllNotBlank should return true when all strings are not blank")
	}

	if IsAllNotBlank("go", " ") {
		t.Fatal("IsAllNotBlank should return false when one string is blank")
	}
}

func TestTernary(t *testing.T) {
	if got := Ternary(true, 1, 2); got != 1 {
		t.Fatalf("Ternary(true, 1, 2) = %d, want 1", got)
	}

	if got := Ternary(false, 1, 2); got != 2 {
		t.Fatalf("Ternary(false, 1, 2) = %d, want 2", got)
	}
}

func TestSubstr(t *testing.T) {
	if got := Substr("hello world", 0, 5); got != "hello" {
		t.Fatalf("Substr(%q, 0, 5) = %q, want %q", "hello world", got, "hello")
	}

	if got := Substr("abcdef", -3, 2); got != "de" {
		t.Fatalf("Substr(%q, -3, 2) = %q, want %q", "abcdef", got, "de")
	}

	if got := Substr("你好世界", 1, 2); got != "好世" {
		t.Fatalf("Substr(%q, 1, 2) = %q, want %q", "你好世界", got, "好世")
	}
}

func TestSubByte(t *testing.T) {
	if got := SubByte("你好世界", 3); got != "你" {
		t.Fatalf("SubByte(%q, 3) = %q, want %q", "你好世界", got, "你")
	}

	if got := SubByte("你好", 10); got != "你好" {
		t.Fatalf("SubByte(%q, 10) = %q, want %q", "你好", got, "你好")
	}

	if got := SubByte("你好", 0); got != "" {
		t.Fatalf("SubByte(%q, 0) = %q, want empty string", "你好", got)
	}
}

func TestStringHelpers(t *testing.T) {
	if got := Slug("Laravel 5 Framework", "-"); got != "Laravel-5-Framework" {
		t.Fatalf("Slug(%q, %q) = %q, want %q", "Laravel 5 Framework", "-", got, "Laravel-5-Framework")
	}

	if got := Escape("a\"b'c"); got != "a\\\"b\\'c" {
		t.Fatalf("Escape(%q) = %q, want %q", "a\"b'c", got, "a\\\"b\\'c")
	}

	if got := Reverse("hello 世界"); got != "界世 olleh" {
		t.Fatalf("Reverse(%q) = %q, want %q", "hello 世界", got, "界世 olleh")
	}

	if got := Char("Go语言"); !slices.Equal(got, []string{"G", "o", "语", "言"}) {
		t.Fatalf("Char(%q) = %#v, want %#v", "Go语言", got, []string{"G", "o", "语", "言"})
	}
}

func TestJoinHelpers(t *testing.T) {
	strs := []string{"hello", " ", "world", "!"}
	if got := BuilderJoin(strs); got != "hello world!" {
		t.Fatalf("BuilderJoin(%#v) = %q, want %q", strs, got, "hello world!")
	}

	if got := BufferJoin(strs); got != "hello world!" {
		t.Fatalf("BufferJoin(%#v) = %q, want %q", strs, got, "hello world!")
	}
}

func TestUnsafeConverters(t *testing.T) {
	if got := Bytes2String([]byte("你好")); got != "你好" {
		t.Fatalf("Bytes2String returned %q, want %q", got, "你好")
	}

	if got := String2Bytes("hello"); !bytes.Equal(got, []byte("hello")) {
		t.Fatalf("String2Bytes returned %q, want %q", got, []byte("hello"))
	}
}
