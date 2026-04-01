package stringx

import "testing"

func TestIsEmail(t *testing.T) {
	if !IsEmail("user@example.com") {
		t.Fatal("IsEmail should accept a valid email address")
	}

	if IsEmail("not-an-email") {
		t.Fatal("IsEmail should reject an invalid email address")
	}
}

func TestIsPhone(t *testing.T) {
	if !IsPhone("13800138000") {
		t.Fatal("IsPhone should accept a valid mainland China mobile number")
	}

	if IsPhone("12345") {
		t.Fatal("IsPhone should reject an invalid mobile number")
	}
}
