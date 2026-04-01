package stringx

import (
	"net/mail"

	"github.com/gtkit/stringx/validatex"
)

// ParseEmailAddress 使用 RFC 5322 语法解析邮箱地址。
func ParseEmailAddress(address string) (*mail.Address, error) {
	return validatex.ParseEmailAddress(address)
}

// IsEmailStrict 判断字符串是否为严格的邮箱地址格式。
func IsEmailStrict(address string) bool { return validatex.IsEmailStrict(address) }

// IsURL 判断字符串是否为带 scheme 和 host 的绝对 URL。
func IsURL(raw string) bool { return validatex.IsURL(raw) }

// IsHTTPURL 判断字符串是否为 http 或 https 绝对 URL。
func IsHTTPURL(raw string) bool { return validatex.IsHTTPURL(raw) }

// IsHostname 判断字符串是否为符合常见 RFC 约束的主机名。
func IsHostname(host string) bool { return validatex.IsHostname(host) }

// IsIP 判断字符串是否为合法 IP 地址。
func IsIP(text string) bool { return validatex.IsIP(text) }

// IsIPv4 判断字符串是否为合法 IPv4 地址。
func IsIPv4(text string) bool { return validatex.IsIPv4(text) }

// IsIPv6 判断字符串是否为合法 IPv6 地址。
func IsIPv6(text string) bool { return validatex.IsIPv6(text) }

// IsCIDR 判断字符串是否为合法 CIDR 表达式。
func IsCIDR(text string) bool { return validatex.IsCIDR(text) }

// IsMAC 判断字符串是否为合法 MAC 地址。
func IsMAC(text string) bool { return validatex.IsMAC(text) }

// IsUUID 判断字符串是否为合法 UUID 文本格式。
func IsUUID(text string) bool { return validatex.IsUUID(text) }

// IsUUIDv4 判断字符串是否为合法 UUID v4 文本格式。
func IsUUIDv4(text string) bool { return validatex.IsUUIDv4(text) }

// IsE164 判断字符串是否符合 E.164 国际电话号码语法格式。
func IsE164(text string) bool { return validatex.IsE164(text) }
