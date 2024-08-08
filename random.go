package stringx

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	rand2 "math/rand/v2"
)

var (
	LowerCaseLettersCharset = []byte("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []byte("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []byte("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)
)

/**
 * 生成随机字符串
 * @param length 字符串长度
 * @param chartype 字符串类型，可选值："l"、"u"、"lu"、"n"、"lun"、"sc"、"all"，默认值为"lun"
 * @return 随机字符串
 * l 小写字母
 * u 大写字母
 * lu 小写字母+大写字母
 * n 数字
 * lun 小写字母+大写字母+数字
 * sc 特殊字符
 * all 所有字符.
 */
func Random(length int, chartype ...string) string {
	var charset []byte
	if len(chartype) == 0 {
		chartype = []string{""}
	}

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
	default:
		charset = AlphanumericCharset
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand2.N(len(charset))]
	}
	return string(b)
}

// SecRandom 生成长度为 length 的安全随机字符串.
func SecRandom(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// 注意：如果你需要确切的字符串长度，请根据base64编码的特性,
	// 调整b的长度，因为base64编码会增加输出长度。
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

// RandomN 生成长度为 length 随机数字字符串.
func RandomN(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
