package stringx

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	rand2 "math/rand/v2"
)

const randIDBytes = 8

var (
	// LowerCaseLettersCharset 表示全部小写英文字母字符集。
	LowerCaseLettersCharset = []byte("abcdefghijklmnopqrstuvwxyz")
	// UpperCaseLettersCharset 表示全部大写英文字母字符集。
	UpperCaseLettersCharset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// LettersCharset 表示大小写英文字母字符集。
	LettersCharset = append(append([]byte(nil), LowerCaseLettersCharset...), UpperCaseLettersCharset...)
	// NumbersCharset 表示十进制数字字符集。
	NumbersCharset = []byte("0123456789")
	// AlphanumericCharset 表示英文字母与数字混合字符集。
	AlphanumericCharset = append(append([]byte(nil), LettersCharset...), NumbersCharset...)
	// SpecialCharset 表示常见特殊字符集。
	SpecialCharset = []byte("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	// AllCharset 表示当前包内置支持的全部字符集。
	AllCharset = append(append([]byte(nil), AlphanumericCharset...), SpecialCharset...)
	hexCharset = []byte("0123456789abcdef")
)

// Random 按指定长度和字符集类型生成随机字符串。
//
// chartype 支持以下取值：
//
//   - `l`：小写字母
//   - `u`：大写字母
//   - `lu`：大小写字母
//   - `n`：数字
//   - `lun`：大小写字母加数字
//   - `sc`：特殊字符
//   - `all`：全部字符
//
// 当 chartype 未提供或取值未知时，默认使用 `lun`。
func Random(length int, chartype ...string) string {
	if length <= 0 {
		return ""
	}

	charset := AlphanumericCharset
	if len(chartype) > 0 {
		switch chartype[0] {
		case "l":
			charset = LowerCaseLettersCharset
		case "u":
			charset = UpperCaseLettersCharset
		case "lu":
			charset = LettersCharset
		case "n":
			charset = NumbersCharset
		case "lun":
			charset = AlphanumericCharset
		case "sc":
			charset = SpecialCharset
		case "all":
			charset = AllCharset
		}
	}

	return randomFromCharset(length, charset)
}

// SecRandom 使用加密安全随机源生成指定长度的随机字符串。
//
// 输出字符采用 URL 安全的 Base64 字符集，并保证返回值长度严格等于 length。
func SecRandom(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	byteLen := (length*3 + 3) / 4
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

	return randomFromCharset(length, NumbersCharset)
}

// RandStr 生成指定长度的随机小写字母字符串。
func RandStr(length int) string {
	return randomFromCharset(length, LowerCaseLettersCharset)
}

// RandStrUpper 生成指定长度的随机大写字母字符串。
func RandStrUpper(length int) string {
	return randomFromCharset(length, UpperCaseLettersCharset)
}

// RandId 生成一个固定 16 个十六进制字符的随机 ID。
//
// 该方法优先使用加密安全随机源；若安全随机源不可用，则退化为伪随机方案，但仍保证长度稳定。
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
//
// 当切片为空时，返回元素类型的零值。
func RandomEle[T any](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(slice))))
	if err != nil {
		return slice[rand2.IntN(len(slice))]
	}

	return slice[index.Int64()]
}

// randomFromCharset 根据给定字符集生成指定长度的随机字符串。
func randomFromCharset(length int, charset []byte) string {
	if length <= 0 || len(charset) == 0 {
		return ""
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand2.IntN(len(charset))]
	}

	return string(b)
}
