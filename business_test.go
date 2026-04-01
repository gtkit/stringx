package stringx

import (
	"strings"
	"testing"
)

func TestChinaIDCard(t *testing.T) {
	valid := "11010519491231002X"
	if err := ValidateChinaIDCard(valid); err != nil {
		t.Fatalf("ValidateChinaIDCard returned error: %v", err)
	}

	if !IsChinaIDCard(valid) {
		t.Fatal("IsChinaIDCard should accept valid ID card")
	}

	if err := ValidateChinaIDCard("110105194912310021"); err == nil || !strings.Contains(err.Error(), "checksum") {
		t.Fatalf("ValidateChinaIDCard should report checksum error, got %v", err)
	}

	if err := ValidateChinaIDCard("99010519491231002X"); err == nil || !strings.Contains(err.Error(), "province") {
		t.Fatalf("ValidateChinaIDCard should report province error, got %v", err)
	}

	if err := ValidateChinaIDCard("11010519990230002X"); err == nil || !strings.Contains(err.Error(), "birth") {
		t.Fatalf("ValidateChinaIDCard should report birth date error, got %v", err)
	}
}

func TestChinaUSCC(t *testing.T) {
	valid := "91350211M000100Y46"
	if err := ValidateChinaUSCC(valid); err != nil {
		t.Fatalf("ValidateChinaUSCC returned error: %v", err)
	}

	if !IsChinaUSCC(valid) {
		t.Fatal("IsChinaUSCC should accept valid code")
	}

	if err := ValidateChinaUSCC("91350211M000100Y44"); err == nil || !strings.Contains(err.Error(), "checksum") {
		t.Fatalf("ValidateChinaUSCC should report checksum error, got %v", err)
	}

	if err := ValidateChinaUSCC("91350211M000100YZ!"); err == nil || !strings.Contains(err.Error(), "character") {
		t.Fatalf("ValidateChinaUSCC should report character error, got %v", err)
	}
}

func TestBankCard(t *testing.T) {
	valid := "4532015112830366"
	if err := ValidateBankCard(valid); err != nil {
		t.Fatalf("ValidateBankCard returned error: %v", err)
	}

	if !IsBankCard(valid) {
		t.Fatal("IsBankCard should accept valid bank card")
	}

	if err := ValidateBankCard("4532015112830367"); err == nil || !strings.Contains(err.Error(), "luhn") {
		t.Fatalf("ValidateBankCard should report Luhn error, got %v", err)
	}

	if err := ValidateBankCard("4532-0151-1283-0366"); err == nil || !strings.Contains(err.Error(), "digit") {
		t.Fatalf("ValidateBankCard should report digit error, got %v", err)
	}

	if got := MaskPhone("13800138000"); got != "138****8000" {
		t.Fatalf("MaskPhone(%q) = %q, want %q", "13800138000", got, "138****8000")
	}

	if got := MaskEmail("alice@example.com"); got != "a***e@example.com" {
		t.Fatalf("MaskEmail(%q) = %q, want %q", "alice@example.com", got, "a***e@example.com")
	}

	if got := MaskName("张三"); got != "张*" {
		t.Fatalf("MaskName(%q) = %q, want %q", "张三", got, "张*")
	}
}
