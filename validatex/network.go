package validatex

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/idna"
)

// IsURL 判断字符串是否为带 scheme 和 host 的绝对 URL。
func IsURL(raw string) bool {
	return ValidateURL(raw) == nil
}

// ValidateURL 验证字符串是否为带 scheme 和 host 的绝对 URL。
func ValidateURL(raw string) error {
	u, err := parseAbsoluteURL(raw)
	if err != nil {
		return err
	}
	return validateURLHost(u)
}

// IsHTTPURL 判断字符串是否为 http 或 https 绝对 URL。
func IsHTTPURL(raw string) bool {
	return ValidateHTTPURL(raw) == nil
}

// ValidateHTTPURL 验证字符串是否为 http 或 https 绝对 URL。
func ValidateHTTPURL(raw string) error {
	u, err := parseAbsoluteURL(raw)
	if err != nil {
		return err
	}
	if !strings.EqualFold(u.Scheme, "http") && !strings.EqualFold(u.Scheme, "https") {
		return fmt.Errorf("unsupported URL scheme %q", u.Scheme)
	}
	return validateURLHost(u)
}

// IsHostname 判断字符串是否为合法主机名。
func IsHostname(host string) bool {
	return ValidateHostname(host) == nil
}

// ValidateHostname 验证字符串是否为合法主机名。
func ValidateHostname(host string) error {
	return normalizeHostnameASCII(host)
}

// IsIP 判断字符串是否为合法 IP 地址。
func IsIP(text string) bool {
	return ValidateIP(text) == nil
}

// ValidateIP 验证字符串是否为合法 IP 地址。
func ValidateIP(text string) error {
	if _, err := netip.ParseAddr(text); err != nil {
		return fmt.Errorf("invalid IP address: %w", err)
	}
	return nil
}

// IsIPv4 判断字符串是否为合法 IPv4 地址。
func IsIPv4(text string) bool {
	return ValidateIPv4(text) == nil
}

// ValidateIPv4 验证字符串是否为合法 IPv4 地址。
func ValidateIPv4(text string) error {
	addr, err := netip.ParseAddr(text)
	if err != nil {
		return fmt.Errorf("invalid IPv4 address: %w", err)
	}
	if !addr.Is4() {
		return errors.New("address is not IPv4")
	}
	return nil
}

// IsIPv6 判断字符串是否为合法 IPv6 地址。
func IsIPv6(text string) bool {
	return ValidateIPv6(text) == nil
}

// ValidateIPv6 验证字符串是否为合法 IPv6 地址。
func ValidateIPv6(text string) error {
	addr, err := netip.ParseAddr(text)
	if err != nil {
		return fmt.Errorf("invalid IPv6 address: %w", err)
	}
	if !addr.Is6() {
		return errors.New("address is not IPv6")
	}
	return nil
}

// IsCIDR 判断字符串是否为合法 CIDR。
func IsCIDR(text string) bool {
	return ValidateCIDR(text) == nil
}

// ValidateCIDR 验证字符串是否为合法 CIDR。
func ValidateCIDR(text string) error {
	if _, err := netip.ParsePrefix(text); err != nil {
		return fmt.Errorf("invalid CIDR: %w", err)
	}
	return nil
}

// IsMAC 判断字符串是否为合法 MAC 地址。
func IsMAC(text string) bool {
	return ValidateMAC(text) == nil
}

// ValidateMAC 验证字符串是否为合法 MAC 地址。
func ValidateMAC(text string) error {
	if _, err := net.ParseMAC(text); err != nil {
		return fmt.Errorf("invalid MAC address: %w", err)
	}
	return nil
}

func parseAbsoluteURL(raw string) (*url.URL, error) {
	if raw == "" || strings.TrimSpace(raw) != raw {
		return nil, errors.New("URL is empty or contains surrounding whitespace")
	}
	u, err := url.ParseRequestURI(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %w", err)
		}
		return nil, errors.New("URL must be absolute and contain scheme and host")
	}
	return u, nil
}

func validateURLHost(u *url.URL) error {
	if err := validateURLPort(u); err != nil {
		return err
	}

	host := u.Hostname()
	if host == "" {
		return errors.New("URL host is empty")
	}
	if ValidateHostname(host) == nil || ValidateIP(host) == nil {
		return nil
	}
	return errors.New("URL host is invalid")
}

func validateURLPort(u *url.URL) error {
	host := u.Host
	if strings.HasPrefix(host, "[") {
		_, rest, found := strings.CutLast(host, "]")
		if !found {
			return errors.New("URL IPv6 host is missing closing bracket")
		}
		if rest == "" {
			return nil
		}
		if rest[0] != ':' {
			return errors.New("URL host is invalid")
		}
		return validatePort(rest[1:])
	}

	if strings.Count(host, ":") != 1 {
		return nil
	}
	_, port, _ := strings.Cut(host, ":")
	return validatePort(port)
}

func validatePort(port string) error {
	if port == "" {
		return errors.New("URL port is empty")
	}
	for i := range len(port) {
		if port[i] < '0' || port[i] > '9' {
			return errors.New("URL port must contain digits only")
		}
	}
	value, err := strconv.Atoi(port)
	if err != nil || value > 65535 {
		return errors.New("URL port is out of range")
	}
	return nil
}

func isHostname(host string) bool {
	return normalizeHostnameASCII(host) == nil
}

func normalizeHostnameASCII(host string) error {
	if host == "" {
		return errors.New("hostname is empty")
	}
	host = strings.TrimSuffix(host, ".")
	if host == "" {
		return errors.New("hostname is empty after trimming trailing dot")
	}
	if _, err := netip.ParseAddr(host); err == nil {
		return errors.New("hostname cannot be a raw IP address")
	}
	asciiHost, err := idna.Lookup.ToASCII(host)
	if err != nil {
		return fmt.Errorf("hostname IDNA conversion failed: %w", err)
	}
	if asciiHost == "" || len(asciiHost) > 253 {
		return errors.New("hostname length is invalid")
	}
	for label := range strings.SplitSeq(asciiHost, ".") {
		if label == "" || len(label) > 63 {
			return errors.New("hostname label length is invalid")
		}
		if label[0] == '-' || label[len(label)-1] == '-' {
			return errors.New("hostname label cannot start or end with hyphen")
		}
		for i := range len(label) {
			b := label[i]
			if !isASCIIAlphaNum(b) && b != '-' {
				return errors.New("hostname contains invalid character")
			}
		}
	}
	return nil
}

func isASCIIAlphaNum(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9'
}

func isIP(text string) bool {
	_, err := netip.ParseAddr(text)
	return err == nil
}

func isIPv6(text string) bool {
	addr, err := netip.ParseAddr(text)
	return err == nil && addr.Is6()
}
