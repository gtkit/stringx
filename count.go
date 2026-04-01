package stringx

import "github.com/gtkit/stringx/textx"

// Len 返回字符串的 UTF-8 rune 数量，而不是底层字节长度。
func Len(str string) int { return textx.Len(str) }

// WordCount 返回字符串中的单词数量。
func WordCount(str string) int { return textx.WordCount(str) }

// Width 返回字符串在等宽字体中的显示宽度。
func Width(str string) int { return textx.Width(str) }

// RuneWidth 返回单个 rune 在等宽字体中的显示宽度。
func RuneWidth(r rune) int { return textx.RuneWidth(r) }
