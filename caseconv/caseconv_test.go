package caseconv

import (
	"strings"
	"testing"
)

func TestCaseConversion(t *testing.T) {
	if got := ToCamel("user_profile"); got != "UserProfile" {
		t.Fatalf("ToCamel returned %q, want %q", got, "UserProfile")
	}

	if got := ToLowerCamel("user_profile"); got != "userProfile" {
		t.Fatalf("ToLowerCamel returned %q, want %q", got, "userProfile")
	}

	if got := ToSnake("HTTPServer"); got != "http_server" {
		t.Fatalf("ToSnake returned %q, want %q", got, "http_server")
	}

	if got := ToUpperSnake("userProfile"); got != "USER_PROFILE" {
		t.Fatalf("ToUpperSnake returned %q, want %q", got, "USER_PROFILE")
	}

	if got := ToKebab("Version2Build3"); got != "version-2-build-3" {
		t.Fatalf("ToKebab returned %q, want %q", got, "version-2-build-3")
	}
}

func TestDelimitedCaseLongASCIIAllocations(t *testing.T) {
	longSnakeInput := strings.Repeat("HTTPServerFieldName", 8)
	if allocs := testing.AllocsPerRun(1000, func() {
		_ = ToSnake(longSnakeInput)
	}); allocs > 1 {
		t.Fatalf("ToSnake long ASCII allocations = %.0f, want <= 1", allocs)
	}

	longKebabInput := strings.Repeat("Version2Build3", 8)
	if allocs := testing.AllocsPerRun(1000, func() {
		_ = ToKebab(longKebabInput)
	}); allocs > 1 {
		t.Fatalf("ToKebab long ASCII allocations = %.0f, want <= 1", allocs)
	}
}
