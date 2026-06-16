package maskx

import "strings"

// MaskPhone 对手机号做脱敏，默认保留前 3 位和后 4 位。
func MaskPhone(text string) string {
	digits := digitsOnly(text)
	if len(digits) >= 13 && strings.HasPrefix(digits, "86") && digits[2] == '1' {
		digits = digits[2:]
	}

	return maskKeepEdges(digits, 3, 4)
}

// MaskEmail 对邮箱地址做脱敏，保留域名。
func MaskEmail(text string) string {
	local, domain, ok := strings.Cut(text, "@")
	if !ok || local == "" || domain == "" {
		return maskKeepEdgesByRune(text, 1, 1)
	}

	runes := []rune(local)
	switch len(runes) {
	case 0:
		return maskKeepEdgesByRune(text, 1, 1)
	case 1:
		return local + "@" + domain
	case 2:
		return string(runes[0]) + "*" + "@" + domain
	default:
		return string(runes[0]) + strings.Repeat("*", len(runes)-2) + string(runes[len(runes)-1]) + "@" + domain
	}
}

// MaskName 对姓名做脱敏。
func MaskName(text string) string {
	runes := []rune(text)
	switch len(runes) {
	case 0, 1:
		return text
	case 2:
		return string(runes[0]) + "*"
	default:
		return string(runes[0]) + strings.Repeat("*", len(runes)-1)
	}
}

// MaskChinaIDCard 对中国身份证号码做脱敏，保留前 6 位和后 4 位。
func MaskChinaIDCard(text string) string {
	return maskKeepEdges(text, 6, 4)
}

// MaskBankCard 对银行卡号做脱敏，保留前 6 位和后 4 位。
func MaskBankCard(text string) string {
	digits := digitsOnly(text)
	if digits == "" {
		return maskKeepEdges(text, 1, 1)
	}
	return maskKeepEdges(digits, 6, 4)
}

func digitsOnly(text string) string {
	var b strings.Builder
	b.Grow(len(text))
	for _, r := range text {
		if '0' <= r && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func maskKeepEdges(text string, prefix, suffix int) string {
	if text == "" {
		return text
	}

	if prefix < 0 {
		prefix = 0
	}
	if suffix < 0 {
		suffix = 0
	}

	if len(text) <= prefix+suffix {
		return maskKeepEdgesByRune(text, 1, 1)
	}

	return text[:prefix] + strings.Repeat("*", len(text)-prefix-suffix) + text[len(text)-suffix:]
}

func maskKeepEdgesByRune(text string, prefix, suffix int) string {
	runes := []rune(text)
	if len(runes) <= 1 {
		return text
	}

	if prefix < 0 {
		prefix = 0
	}
	if suffix < 0 {
		suffix = 0
	}

	if len(runes) <= prefix+suffix {
		prefix = 1
		suffix = 0
		if len(runes) > 2 {
			suffix = 1
		}
	}

	middle := len(runes) - prefix - suffix
	if middle <= 0 {
		middle = len(runes) - 1
		prefix = 1
		suffix = 0
	}

	return string(runes[:prefix]) + strings.Repeat("*", middle) + string(runes[len(runes)-suffix:])
}
