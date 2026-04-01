package maskx

import "testing"

func TestMaskHelpers(t *testing.T) {
	if got := MaskPhone("13800138000"); got != "138****8000" {
		t.Fatalf("MaskPhone returned %q, want %q", got, "138****8000")
	}

	if got := MaskPhone("+86 138 0013 8000"); got != "138****8000" {
		t.Fatalf("MaskPhone returned %q, want %q", got, "138****8000")
	}

	if got := MaskEmail("alice@example.com"); got != "a***e@example.com" {
		t.Fatalf("MaskEmail returned %q, want %q", got, "a***e@example.com")
	}

	if got := MaskEmail("alicelong"); got != "a*******g" {
		t.Fatalf("MaskEmail returned %q, want %q", got, "a*******g")
	}

	if got := MaskName("张三"); got != "张*" {
		t.Fatalf("MaskName returned %q, want %q", got, "张*")
	}

	if got := MaskName("欧阳娜娜"); got != "欧***" {
		t.Fatalf("MaskName returned %q, want %q", got, "欧***")
	}

	if got := MaskChinaIDCard("11010519491231002X"); got != "110105********002X" {
		t.Fatalf("MaskChinaIDCard returned %q, want %q", got, "110105********002X")
	}

	if got := MaskBankCard("4532015112830366"); got != "453201******0366" {
		t.Fatalf("MaskBankCard returned %q, want %q", got, "453201******0366")
	}
}
