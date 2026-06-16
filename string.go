package stringx

import (
	"github.com/gtkit/stringx/textx"
	"github.com/gtkit/stringx/unsafex"
)

// String2Bytes 在不发生额外内存拷贝的情况下，将字符串转换为字节切片。
//
// 返回的切片与原字符串共享底层数据，只适合只读场景，调用方不得修改返回切片。
// 如果结果需要被修改，请使用 CloneStringBytes。
func String2Bytes(s string) []byte { return unsafex.String2Bytes(s) }

// Bytes2String 在不发生额外内存拷贝的情况下，将字节切片转换为字符串。
//
// 返回的字符串与原字节切片共享底层数据；原切片后续修改会影响字符串视图。
// 如果需要稳定不可变快照，请使用 CloneBytesString。
func Bytes2String(b []byte) string { return unsafex.Bytes2String(b) }

// CloneStringBytes 将字符串复制为可安全修改的字节切片。
func CloneStringBytes(s string) []byte { return []byte(s) }

// CloneBytesString 将字节切片复制为不受后续切片修改影响的字符串。
func CloneBytesString(b []byte) string { return string(b) }

// LowerFirst 将字符串首个 rune 转换为小写，其余内容保持不变。
func LowerFirst(input string) string { return textx.LowerFirst(input) }

// UpperFirst 将字符串首个 rune 转换为大写，其余内容保持不变。
func UpperFirst(input string) string { return textx.UpperFirst(input) }

// FirstLowerCase 返回字符串首个 rune 的小写形式。
func FirstLowerCase(input string) string { return textx.FirstLowerCase(input) }

// FirstUpperCase 返回字符串首个 rune 的大写形式。
func FirstUpperCase(input string) string { return textx.FirstUpperCase(input) }

// IsEmpty 判断字符串是否为空字符串。
func IsEmpty(str string) bool { return textx.IsEmpty(str) }

// IsNotEmpty 判断字符串是否为非空字符串。
func IsNotEmpty(str string) bool { return textx.IsNotEmpty(str) }

// IsBlank 判断字符串在去除首尾空白字符后是否为空。
func IsBlank(str string) bool { return textx.IsBlank(str) }

// IsNotBlank 判断字符串在去除首尾空白字符后是否仍然非空。
func IsNotBlank(str string) bool { return textx.IsNotBlank(str) }

// IsAllBlank 判断给定的所有字符串是否全部为空白字符串。
func IsAllBlank(strs ...string) bool { return textx.IsAllBlank(strs...) }

// IsAllNotBlank 判断给定的所有字符串是否全部为非空白字符串。
func IsAllNotBlank(strs ...string) bool { return textx.IsAllNotBlank(strs...) }

// Ternary 根据 cond 的真假返回 a 或 b。
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

// Substr 按 rune 下标截取字符串。
func Substr(s string, start int, strlength ...int) string {
	return textx.Substr(s, start, strlength...)
}

// Slug 将字符串中的空格替换为指定分隔符。
func Slug(str, separator string) string { return textx.Slug(str, separator) }

// NormalizeSpace 去除首尾空白，并将中间连续空白折叠为单个 ASCII 空格。
func NormalizeSpace(str string) string { return textx.NormalizeSpace(str) }

// SubByte 按字节长度安全截取 UTF-8 字符串，避免返回截断的无效编码。
func SubByte(str string, length int) string { return textx.SubByte(str, length) }

// Char 将字符串拆分为按 rune 划分的字符串切片。
func Char(str string) []string { return textx.Char(str) }

// Escape 返回适合放入 Go 双引号字符串字面量中的转义结果，并移除外围引号。
func Escape(s string) string { return textx.Escape(s) }

// Reverse 按 rune 维度反转字符串。
func Reverse(s string) string { return textx.Reverse(s) }
