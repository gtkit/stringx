package randomx

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	rand2 "math/rand/v2"
)

const (
	randIDBytes = 8

	lowerCaseLettersCharset = "abcdefghijklmnopqrstuvwxyz"
	upperCaseLettersCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lettersCharset          = lowerCaseLettersCharset + upperCaseLettersCharset
	numbersCharset          = "0123456789"
	alphanumericCharset     = lettersCharset + numbersCharset
	specialCharset          = "!@#$%^&*()_+-=[]{}|;':\",./<>?"
	allCharset              = alphanumericCharset + specialCharset
	hexCharset              = "0123456789abcdef"
)

var (
	// LowerCaseLettersCharset 表示全部小写英文字母字符集。
	LowerCaseLettersCharset = []byte(lowerCaseLettersCharset)
	// UpperCaseLettersCharset 表示全部大写英文字母字符集。
	UpperCaseLettersCharset = []byte(upperCaseLettersCharset)
	// LettersCharset 表示大小写英文字母字符集。
	LettersCharset = []byte(lettersCharset)
	// NumbersCharset 表示十进制数字字符集。
	NumbersCharset = []byte(numbersCharset)
	// AlphanumericCharset 表示英文字母与数字混合字符集。
	AlphanumericCharset = []byte(alphanumericCharset)
	// SpecialCharset 表示常见特殊字符集。
	SpecialCharset = []byte(specialCharset)
	// AllCharset 表示当前包支持的全部字符集。
	AllCharset = []byte(allCharset)
)

// Random 按指定长度和字符集类型生成随机字符串。
func Random(length int, chartype ...string) string {
	if length <= 0 {
		return ""
	}

	charset := alphanumericCharset
	if len(chartype) > 0 {
		switch chartype[0] {
		case "l":
			charset = lowerCaseLettersCharset
		case "u":
			charset = upperCaseLettersCharset
		case "lu":
			charset = lettersCharset
		case "n":
			charset = numbersCharset
		case "lun":
			charset = alphanumericCharset
		case "sc":
			charset = specialCharset
		case "all":
			charset = allCharset
		}
	}

	return randomFromCharset(length, charset)
}

// RandomFromCharset 使用指定字符集生成随机字符串。
func RandomFromCharset(length int, charset []byte) string {
	if length <= 0 || len(charset) == 0 {
		return ""
	}
	return randomFromCharset(length, string(charset))
}

// SecRandom 使用加密安全随机源生成指定长度的随机字符串。
func SecRandom(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	byteLen := secRandomByteLen(length)
	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b)[:length], nil
}

// RandomN 生成指定长度的随机数字字符串。
func RandomN(length int) string {
	if length <= 0 {
		return ""
	}
	return randomFromCharset(length, numbersCharset)
}

// RandStr 生成指定长度的随机小写字母字符串。
func RandStr(length int) string {
	return randomFromCharset(length, lowerCaseLettersCharset)
}

// RandStrUpper 生成指定长度的随机大写字母字符串。
func RandStrUpper(length int) string {
	return randomFromCharset(length, upperCaseLettersCharset)
}

// RandId 生成固定 16 个十六进制字符的随机 ID。
func RandId() string {
	b := make([]byte, randIDBytes)
	if _, err := rand.Read(b); err == nil {
		return hex.EncodeToString(b)
	}

	return randomFromCharset(randIDBytes*2, hexCharset)
}

// Deprecated: Use RandomN instead.
func Randn(length int) string { return RandomN(length) }

// RandomEle 从切片中随机返回一个元素。
func RandomEle[T any](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(slice))))
	if err != nil {
		return slice[rand2.IntN(len(slice))]
	}

	return slice[index.Int64()]
}

func randomFromCharset(length int, charset string) string {
	if length <= 0 || charset == "" {
		return ""
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand2.IntN(len(charset))]
	}

	return string(b)
}

func secRandomByteLen(length int) int {
	byteLen := length / 4 * 3
	if rem := length % 4; rem > 0 {
		byteLen += rem
	}
	return byteLen
}
