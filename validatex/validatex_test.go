package validatex

import (
	"strings"
	"testing"
	"time"
)

func TestValidateAndParse(t *testing.T) {
	if err := ValidateEmailStrict("user@example.com"); err != nil {
		t.Fatalf("ValidateEmailStrict returned error: %v", err)
	}

	if err := ValidateEmailStrict(strings.Repeat("a", 65) + "@example.com"); err == nil {
		t.Fatal("ValidateEmailStrict should reject local parts longer than 64 characters")
	}

	if err := ValidateEmailStrict(strings.Repeat("a", 250) + "@example.com"); err == nil {
		t.Fatal("ValidateEmailStrict should reject addresses longer than protocol limits")
	}

	if !IsEmailStrict("user@例子.中国") {
		t.Fatal("IsEmailStrict should accept IDNA domain")
	}

	if err := ValidateHTTPURL("https://例子.中国"); err != nil {
		t.Fatalf("ValidateHTTPURL returned error: %v", err)
	}

	if !IsHostname("例子.中国") {
		t.Fatal("IsHostname should accept IDNA hostname")
	}

	if err := ValidateUUIDv4("550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatalf("ValidateUUIDv4 returned error: %v", err)
	}

	if err := ValidateE164("+8613800138000"); err != nil {
		t.Fatalf("ValidateE164 returned error: %v", err)
	}

	if err := ValidateChinaMobilePhone("13800138000"); err != nil {
		t.Fatalf("ValidateChinaMobilePhone returned error: %v", err)
	}

	if err := ValidateChinaMobilePhone(strings.Repeat("1", 65)); err == nil {
		t.Fatal("ValidateChinaMobilePhone should reject excessively long input")
	}

	if err := ValidateChinaIDCard("11010519491231002X"); err != nil {
		t.Fatalf("ValidateChinaIDCard returned error: %v", err)
	}

	if err := ValidateChinaUSCC("91350211M000100Y46"); err != nil {
		t.Fatalf("ValidateChinaUSCC returned error: %v", err)
	}

	if err := ValidateBankCard("4532015112830366"); err != nil {
		t.Fatalf("ValidateBankCard returned error: %v", err)
	}

	info, err := ParseChinaIDCard("11010519491231002X", time.Date(2026, 3, 28, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("ParseChinaIDCard returned error: %v", err)
	}

	if info.BirthDate.Format("2006-01-02") != "1949-12-31" {
		t.Fatalf("BirthDate = %s, want 1949-12-31", info.BirthDate.Format("2006-01-02"))
	}

	if info.Gender != GenderFemale {
		t.Fatalf("Gender = %v, want %v", info.Gender, GenderFemale)
	}

	if info.Age != 76 {
		t.Fatalf("Age = %d, want 76", info.Age)
	}
}

func TestValidateChinaMobilePhoneFormattingRules(t *testing.T) {
	validCases := []string{
		"13800138000",
		"+86 138 0013 8000",
		"86-138-0013-8000",
		"(+86) 138-0013-8000",
	}

	for _, input := range validCases {
		if err := ValidateChinaMobilePhone(input); err != nil {
			t.Fatalf("ValidateChinaMobilePhone(%q) returned error: %v", input, err)
		}
	}

	invalidCases := []string{
		"138.0013.8000",
		"138/0013/8000",
		"+1 650-253-0000",
		"12600138000",
	}

	for _, input := range invalidCases {
		if err := ValidateChinaMobilePhone(input); err == nil {
			t.Fatalf("ValidateChinaMobilePhone(%q) should reject unsupported input", input)
		}
	}
}
