package stringx

import (
	"testing"
)

func TestParseEmailAddress(t *testing.T) {
	addr, err := ParseEmailAddress(`Alice <alice@example.com>`)
	if err != nil {
		t.Fatalf("ParseEmailAddress returned error: %v", err)
	}

	if addr.Address != "alice@example.com" {
		t.Fatalf("ParseEmailAddress address = %q, want %q", addr.Address, "alice@example.com")
	}
}

func TestIsEmailStrict(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{name: "plain addr-spec", in: "user@example.com", want: true},
		{name: "display name not allowed", in: `Alice <alice@example.com>`, want: false},
		{name: "domain literal ipv4", in: "user@[127.0.0.1]", want: true},
		{name: "domain literal ipv6", in: "user@[IPv6:2001:db8::1]", want: true},
		{name: "missing domain", in: "user@", want: false},
		{name: "invalid domain", in: "user@-example.com", want: false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := IsEmailStrict(tc.in); got != tc.want {
				t.Fatalf("IsEmailStrict(%q) = %t, want %t", tc.in, got, tc.want)
			}
		})
	}
}

func TestStrictNetworkValidators(t *testing.T) {
	if !IsURL("https://example.com/path?q=1") {
		t.Fatal("IsURL should accept a valid absolute URL")
	}

	if IsURL("/relative/path") {
		t.Fatal("IsURL should reject a relative path")
	}

	if !IsHTTPURL("https://example.com") {
		t.Fatal("IsHTTPURL should accept https URL")
	}

	if IsHTTPURL("ftp://example.com") {
		t.Fatal("IsHTTPURL should reject non-http schemes")
	}

	if !IsHostname("example.com") {
		t.Fatal("IsHostname should accept a valid hostname")
	}

	if IsHostname("-bad.example.com") {
		t.Fatal("IsHostname should reject invalid hostname")
	}

	if !IsIP("2001:db8::1") {
		t.Fatal("IsIP should accept valid IP")
	}

	if !IsIPv4("127.0.0.1") {
		t.Fatal("IsIPv4 should accept IPv4")
	}

	if !IsIPv6("2001:db8::1") {
		t.Fatal("IsIPv6 should accept IPv6")
	}

	if !IsCIDR("10.0.0.0/24") {
		t.Fatal("IsCIDR should accept valid CIDR")
	}

	if !IsMAC("01:23:45:67:89:ab") {
		t.Fatal("IsMAC should accept valid MAC")
	}
}

func TestStrictIdentifiers(t *testing.T) {
	if !IsUUID("550e8400-e29b-41d4-a716-446655440000") {
		t.Fatal("IsUUID should accept valid UUID")
	}

	if !IsUUIDv4("550e8400-e29b-41d4-a716-446655440000") {
		t.Fatal("IsUUIDv4 should accept valid v4 UUID")
	}

	if IsUUIDv4("550e8400-e29b-11d4-a716-446655440000") {
		t.Fatal("IsUUIDv4 should reject non-v4 UUID")
	}

	if !IsE164("+8613800138000") {
		t.Fatal("IsE164 should accept valid E.164 number")
	}

	if IsE164("13800138000") {
		t.Fatal("IsE164 should reject number without plus sign")
	}
}

func TestNormalizeSpace(t *testing.T) {
	if got := NormalizeSpace("  hello \t world \n  golang  "); got != "hello world golang" {
		t.Fatalf("NormalizeSpace returned %q, want %q", got, "hello world golang")
	}
}

func TestIDNAValidators(t *testing.T) {
	if !IsHostname("例子.中国") {
		t.Fatal("IsHostname should accept valid IDNA hostname")
	}

	if !IsURL("https://例子.中国/路径?q=1") {
		t.Fatal("IsURL should accept URL with IDNA hostname")
	}

	if !IsHTTPURL("https://例子.中国") {
		t.Fatal("IsHTTPURL should accept https URL with IDNA hostname")
	}

	if !IsEmailStrict("user@例子.中国") {
		t.Fatal("IsEmailStrict should accept email with IDNA domain")
	}

	if IsHostname("bad..例子.中国") {
		t.Fatal("IsHostname should reject invalid IDNA hostname")
	}
}

func TestValidateStrictAPIs(t *testing.T) {
	if err := ValidateEmailStrict("user@example.com"); err != nil {
		t.Fatalf("ValidateEmailStrict returned error: %v", err)
	}

	if err := ValidateEmailStrict("bad@-example.com"); err == nil {
		t.Fatal("ValidateEmailStrict should reject invalid email")
	}

	if err := ValidateHTTPURL("https://例子.中国"); err != nil {
		t.Fatalf("ValidateHTTPURL returned error: %v", err)
	}

	if err := ValidateHTTPURL("ftp://example.com"); err == nil {
		t.Fatal("ValidateHTTPURL should reject non-http URL")
	}

	if err := ValidateHostname("例子.中国"); err != nil {
		t.Fatalf("ValidateHostname returned error: %v", err)
	}

	if err := ValidateIPv4("127.0.0.1"); err != nil {
		t.Fatalf("ValidateIPv4 returned error: %v", err)
	}

	if err := ValidateIPv4("2001:db8::1"); err == nil {
		t.Fatal("ValidateIPv4 should reject IPv6")
	}

	if err := ValidateUUIDv4("550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatalf("ValidateUUIDv4 returned error: %v", err)
	}

	if err := ValidateUUIDv4("550e8400-e29b-11d4-a716-446655440000"); err == nil {
		t.Fatal("ValidateUUIDv4 should reject non-v4 UUID")
	}

	if err := ValidateE164("+8613800138000"); err != nil {
		t.Fatalf("ValidateE164 returned error: %v", err)
	}

}
