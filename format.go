package stringx

import "github.com/gtkit/stringx/textx"

// LeftJustify 在字符串右侧补齐 pad，直到 rune 长度达到 length。
func LeftJustify(str string, length int, pad string) string {
	return textx.LeftJustify(str, length, pad)
}

// RightJustify 在字符串左侧补齐 pad，直到 rune 长度达到 length。
func RightJustify(str string, length int, pad string) string {
	return textx.RightJustify(str, length, pad)
}

// Center 在字符串左右两侧补齐 pad，直到 rune 长度达到 length。
func Center(str string, length int, pad string) string { return textx.Center(str, length, pad) }
