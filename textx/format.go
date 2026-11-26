package textx

import (
	"strings"
	"unicode/utf8"
)

// RuneWidth 返回单个 rune 的显示宽度。
func RuneWidth(r rune) int {
	switch {
	case r == utf8.RuneError || r < '\x20':
		return 0
	case '\x20' <= r && r < '\u2000':
		return 1
	case '\u2000' <= r && r < '\uFF61':
		return 2
	case '\uFF61' <= r && r < '\uFFA0':
		return 1
	case '\uFFA0' <= r:
		return 2
	}
	return 0
}

// Width 返回字符串在等宽字体中的显示宽度。
func Width(str string) int {
	var width int
	for str != "" {
		r, size := utf8.DecodeRuneInString(str)
		width += RuneWidth(r)
		str = str[size:]
	}
	return width
}

// LeftJustify 在右侧填充字符串。
func LeftJustify(str string, length int, pad string) string {
	l := Len(str)
	if l >= length || pad == "" {
		return str
	}
	remains := length - l
	padLen := Len(pad)
	output := &strings.Builder{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	output.WriteString(str)
	writePadString(output, pad, padLen, remains)
	return output.String()
}

// RightJustify 在左侧填充字符串。
func RightJustify(str string, length int, pad string) string {
	l := Len(str)
	if l >= length || pad == "" {
		return str
	}
	remains := length - l
	padLen := Len(pad)
	output := &strings.Builder{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains)
	output.WriteString(str)
	return output.String()
}

// Center 在左右两侧填充字符串。
func Center(str string, length int, pad string) string {
	l := Len(str)
	if l >= length || pad == "" {
		return str
	}
	remains := length - l
	padLen := Len(pad)
	output := &strings.Builder{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains/2)
	output.WriteString(str)
	writePadString(output, pad, padLen, (remains+1)/2)
	return output.String()
}

func writePadString(output *strings.Builder, pad string, padLen, remains int) {
	repeats := remains / padLen
	for range repeats {
		output.WriteString(pad)
	}
	remains %= padLen
	for range remains {
		r, size := utf8.DecodeRuneInString(pad)
		output.WriteRune(r)
		pad = pad[size:]
	}
}
