package business

import (
	"errors"
	"strings"
)

// ValidateChinaUSCC 验证统一社会信用代码。
func ValidateChinaUSCC(text string) error {
	if len(text) != 18 {
		return errors.New("USCC length must be 18")
	}

	text = strings.ToUpper(text)
	sum := 0
	for i := range 17 {
		value, ok := chinaUSCCValueMap[text[i]]
		if !ok {
			return errors.New("USCC contains invalid character")
		}
		sum += value * chinaUSCCWeights[i]
	}
	lastValue, ok := chinaUSCCValueMap[text[17]]
	if !ok {
		return errors.New("USCC contains invalid character")
	}
	check := (31 - sum%31) % 31
	if lastValue != check {
		return errors.New("USCC checksum is invalid")
	}
	return nil
}

// IsChinaUSCC 判断字符串是否为合法统一社会信用代码。
func IsChinaUSCC(text string) bool {
	return ValidateChinaUSCC(text) == nil
}

// ValidateBankCard 使用 Luhn 算法验证银行卡号。
func ValidateBankCard(text string) error {
	if len(text) < 12 || len(text) > 19 {
		return errors.New("bank card length must be between 12 and 19")
	}
	sum := 0
	double := false
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] < '0' || text[i] > '9' {
			return errors.New("bank card must contain digits only")
		}
		digit := int(text[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	if sum%10 != 0 {
		return errors.New("bank card luhn checksum is invalid")
	}
	return nil
}

// IsBankCard 判断字符串是否为通过 Luhn 校验的银行卡号。
func IsBankCard(text string) bool {
	return ValidateBankCard(text) == nil
}

const chinaUSCCCharset = "0123456789ABCDEFGHJKLMNPQRTUWXY"

var chinaUSCCValueMap = func() map[byte]int {
	m := make(map[byte]int, len(chinaUSCCCharset))
	for i := range len(chinaUSCCCharset) {
		m[chinaUSCCCharset[i]] = i
	}
	return m
}()

var chinaUSCCWeights = [...]int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}
