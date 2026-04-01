package stringx

import "github.com/gtkit/stringx/validatex"

// ValidateEmailStrict 验证字符串是否为严格的邮箱地址格式。
func ValidateEmailStrict(address string) error { return validatex.ValidateEmailStrict(address) }

// ValidateURL 验证字符串是否为带 scheme 和 host 的绝对 URL。
func ValidateURL(raw string) error { return validatex.ValidateURL(raw) }

// ValidateHTTPURL 验证字符串是否为 http 或 https 绝对 URL。
func ValidateHTTPURL(raw string) error { return validatex.ValidateHTTPURL(raw) }

// ValidateHostname 验证字符串是否为合法主机名。
func ValidateHostname(host string) error { return validatex.ValidateHostname(host) }

// ValidateIP 验证字符串是否为合法 IP 地址。
func ValidateIP(text string) error { return validatex.ValidateIP(text) }

// ValidateIPv4 验证字符串是否为合法 IPv4 地址。
func ValidateIPv4(text string) error { return validatex.ValidateIPv4(text) }

// ValidateIPv6 验证字符串是否为合法 IPv6 地址。
func ValidateIPv6(text string) error { return validatex.ValidateIPv6(text) }

// ValidateCIDR 验证字符串是否为合法 CIDR。
func ValidateCIDR(text string) error { return validatex.ValidateCIDR(text) }

// ValidateMAC 验证字符串是否为合法 MAC 地址。
func ValidateMAC(text string) error { return validatex.ValidateMAC(text) }

// ValidateUUID 验证字符串是否为合法 UUID。
func ValidateUUID(text string) error { return validatex.ValidateUUID(text) }

// ValidateUUIDv4 验证字符串是否为合法 UUID v4。
func ValidateUUIDv4(text string) error { return validatex.ValidateUUIDv4(text) }

// ValidateE164 验证字符串是否符合 E.164 语法格式。
func ValidateE164(text string) error { return validatex.ValidateE164(text) }
