package stringx

import "github.com/gtkit/stringx/randomx"

var (
	// LowerCaseLettersCharset 表示全部小写英文字母字符集。
	LowerCaseLettersCharset = append([]byte(nil), randomx.LowerCaseLettersCharset...)
	// UpperCaseLettersCharset 表示全部大写英文字母字符集。
	UpperCaseLettersCharset = append([]byte(nil), randomx.UpperCaseLettersCharset...)
	// LettersCharset 表示大小写英文字母字符集。
	LettersCharset = append([]byte(nil), randomx.LettersCharset...)
	// NumbersCharset 表示十进制数字字符集。
	NumbersCharset = append([]byte(nil), randomx.NumbersCharset...)
	// AlphanumericCharset 表示英文字母与数字混合字符集。
	AlphanumericCharset = append([]byte(nil), randomx.AlphanumericCharset...)
	// SpecialCharset 表示常见特殊字符集。
	SpecialCharset = append([]byte(nil), randomx.SpecialCharset...)
	// AllCharset 表示当前包内置支持的全部字符集。
	AllCharset = append([]byte(nil), randomx.AllCharset...)
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

	return randomx.RandomFromCharset(length, charset)
}

// RandomFromCharset 使用指定字符集生成随机字符串。
func RandomFromCharset(length int, charset []byte) string {
	return randomx.RandomFromCharset(length, charset)
}

// SecRandom 使用加密安全随机源生成指定长度的随机字符串。
//
// 输出字符采用 URL 安全的 Base64 字符集，并保证返回值长度严格等于 length。
func SecRandom(length int) (string, error) {
	return randomx.SecRandom(length)
}

// RandomN 生成指定长度的随机数字字符串。
func RandomN(length int) string {
	return randomx.RandomFromCharset(length, NumbersCharset)
}

// RandStr 生成指定长度的随机小写字母字符串。
func RandStr(length int) string {
	return randomx.RandomFromCharset(length, LowerCaseLettersCharset)
}

// RandStrUpper 生成指定长度的随机大写字母字符串。
func RandStrUpper(length int) string {
	return randomx.RandomFromCharset(length, UpperCaseLettersCharset)
}

// RandId 生成一个固定 16 个十六进制字符的随机 ID。
//
// 该方法优先使用加密安全随机源；若安全随机源不可用，则退化为伪随机方案，但仍保证长度稳定。
func RandId() string {
	return randomx.RandId()
}

// Deprecated: Use RandomN instead.
func Randn(length int) string { return RandomN(length) }

// RandomEle 从切片中随机返回一个元素。
//
// 当切片为空时，返回元素类型的零值。
func RandomEle[T any](slice []T) T {
	return randomx.RandomEle(slice)
}
