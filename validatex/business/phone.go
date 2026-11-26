package business

import (
	"errors"
	"strconv"
	"strings"
)

// ValidateChinaMobilePhone 使用本地规则严格验证中国手机号码。
//
// 支持可选的 `+86` 或 `86` 前缀，并允许空格、连字符和括号作为格式化字符。
func ValidateChinaMobilePhone(text string) error {
	number, err := normalizeChinaMobilePhone(text)
	if err != nil {
		return err
	}
	if _, ok := chinaMobilePrefixes[number[:3]]; !ok {
		return errors.New("phone number prefix is invalid")
	}
	return nil
}

func normalizeChinaMobilePhone(text string) (string, error) {
	if len(text) > 64 {
		return "", errors.New("phone input is too long")
	}

	var compact strings.Builder
	compact.Grow(len(text))

	for _, r := range text {
		switch r {
		case ' ', '-':
			continue
		case '+', '(', ')':
		default:
			if r < '0' || r > '9' {
				return "", errors.New("phone input contains unsupported characters")
			}
		}
		compact.WriteRune(r)
	}

	normalized := compact.String()
	for _, prefix := range [...]string{"(+86)", "(86)", "+86", "86"} {
		if number, found := strings.CutPrefix(normalized, prefix); found {
			normalized = number
			break
		}
	}

	if strings.ContainsAny(normalized, "+()") {
		return "", errors.New("phone input format is invalid")
	}
	if len(normalized) != 11 {
		return "", errors.New("phone number must contain 11 digits")
	}
	if normalized[0] != '1' {
		return "", errors.New("phone number must start with 1")
	}
	return normalized, nil
}

var chinaMobilePrefixes = func() map[string]struct{} {
	prefixes := make(map[string]struct{}, 55)
	addRange := func(start, end int) {
		for prefix := start; prefix <= end; prefix++ {
			prefixes[strconv.Itoa(prefix)] = struct{}{}
		}
	}

	addRange(130, 139)
	addRange(145, 149)
	addRange(150, 159)
	for _, prefix := range []int{162, 165, 166, 167} {
		prefixes[strconv.Itoa(prefix)] = struct{}{}
	}
	addRange(170, 179)
	addRange(180, 189)
	addRange(190, 199)
	return prefixes
}()
