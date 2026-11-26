package validatex

import (
	"errors"
	"net/mail"
	"strings"
)

// ParseEmailAddress 使用 RFC 5322 语法解析邮箱地址。
func ParseEmailAddress(address string) (*mail.Address, error) {
	return mail.ParseAddress(address)
}

// IsEmailStrict 判断字符串是否为严格邮箱格式。
func IsEmailStrict(address string) bool {
	return ValidateEmailStrict(address) == nil
}

// ValidateEmailStrict 验证字符串是否为严格邮箱格式。
func ValidateEmailStrict(address string) error {
	if address == "" {
		return errors.New("email is empty")
	}
	if len(address) > 254 {
		return errors.New("email exceeds maximum length")
	}
	if strings.TrimSpace(address) != address {
		return errors.New("email contains surrounding whitespace")
	}

	local, domain, found := strings.CutLast(address, "@")
	if !found || local == "" || domain == "" {
		return errors.New("email must contain local part and domain")
	}

	if len(local) > 64 {
		return errors.New("email local part exceeds maximum length")
	}
	if !isEmailLocalPart(local) {
		return errors.New("email local part is invalid")
	}
	if !isEmailDomain(domain) {
		return errors.New("email domain is invalid")
	}
	return nil
}

func isEmailLocalPart(local string) bool {
	if local == "" {
		return false
	}
	synthetic := local + "@example.com"
	addr, err := mail.ParseAddress(synthetic)
	return err == nil && addr.Name == "" && addr.Address == synthetic
}

func isEmailDomain(domain string) bool {
	if strings.HasPrefix(domain, "[") && strings.HasSuffix(domain, "]") {
		literal := domain[1 : len(domain)-1]
		if ip, found := strings.CutPrefix(literal, "IPv6:"); found {
			return isIPv6(ip)
		}
		return isIP(literal)
	}
	return isHostname(domain)
}
