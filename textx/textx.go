package textx

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// LowerFirst 将字符串首个 rune 转换为小写。
func LowerFirst(input string) string {
	r, size := firstRuneInfo(input)
	if size == 0 {
		return ""
	}
	lower := unicode.ToLower(r)
	if lower == r {
		return input
	}
	return string(lower) + input[size:]
}

// UpperFirst 将字符串首个 rune 转换为大写。
func UpperFirst(input string) string {
	r, size := firstRuneInfo(input)
	if size == 0 {
		return input
	}
	upper := unicode.ToUpper(r)
	if upper == r {
		return input
	}
	return string(upper) + input[size:]
}

// FirstLowerCase 返回字符串首个 rune 的小写形式。
func FirstLowerCase(input string) string {
	r, size := firstRuneInfo(input)
	if size == 0 {
		return input
	}
	return string(unicode.ToLower(r))
}

// FirstUpperCase 返回字符串首个 rune 的大写形式。
func FirstUpperCase(input string) string {
	r, size := firstRuneInfo(input)
	if size == 0 {
		return input
	}
	return string(unicode.ToUpper(r))
}

// IsEmpty 判断字符串是否为空。
func IsEmpty(str string) bool { return str == "" }

// IsNotEmpty 判断字符串是否非空。
func IsNotEmpty(str string) bool { return str != "" }

// IsBlank 判断字符串在去除首尾空白后是否为空。
func IsBlank(str string) bool {
	for _, r := range str {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// IsNotBlank 判断字符串在去除首尾空白后是否仍非空。
func IsNotBlank(str string) bool { return !IsBlank(str) }

// IsAllBlank 判断所有字符串是否全为空白。
func IsAllBlank(strs ...string) bool {
	for _, str := range strs {
		if !IsBlank(str) {
			return false
		}
	}
	return true
}

// IsAllNotBlank 判断所有字符串是否全为非空白。
func IsAllNotBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return false
		}
	}
	return true
}

// NormalizeSpace 去除首尾空白，并将连续空白折叠为单个空格。
func NormalizeSpace(str string) string {
	if str == "" || isNormalizedSpace(str) {
		return str
	}
	if isASCII(str) {
		return normalizeASCIISpace(str)
	}

	var builder strings.Builder
	builder.Grow(len(str))
	pendingSpace := false
	for _, r := range str {
		if unicode.IsSpace(r) {
			if builder.Len() > 0 {
				pendingSpace = true
			}
			continue
		}
		if pendingSpace {
			builder.WriteByte(' ')
			pendingSpace = false
		}
		builder.WriteRune(r)
	}
	return builder.String()
}

// Substr 按 rune 下标截取字符串。
func Substr(s string, start int, length ...int) string {
	charlist := []rune(s)
	l := len(charlist)

	size := l
	if len(length) != 0 {
		size = length[0]
	}
	if start < 0 {
		start = l + start
	}
	end := start + size
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > l {
		start = l
	}
	if end < 0 {
		end = 0
	}
	if end > l {
		end = l
	}
	return string(charlist[start:end])
}

// Slug 将空格替换为指定分隔符。
func Slug(str, separator string) string {
	return strings.ReplaceAll(str, " ", separator)
}

// SubByte 按字节长度安全截取 UTF-8 字符串。
func SubByte(str string, length int) string {
	switch {
	case length <= 0:
		return ""
	case length >= len(str):
		return str
	}

	for length > 0 {
		r, size := utf8.DecodeLastRuneInString(str[:length])
		if r != utf8.RuneError || size != 1 {
			return str[:length]
		}
		length--
	}
	return ""
}

// Char 将字符串拆分为按 rune 划分的切片。
func Char(str string) []string {
	result := make([]string, 0, Len(str))
	for _, r := range str {
		result = append(result, string(r))
	}
	return result
}

// Escape 返回适合放入 Go 双引号字面量中的转义结果，并去掉外围引号。
func Escape(s string) string {
	quoted := strconv.Quote(s)
	return strings.ReplaceAll(quoted[1:len(quoted)-1], "'", "\\'")
}

// Reverse 按 rune 反转字符串。
func Reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// BuilderJoin 使用 strings.Builder 拼接字符串。
func BuilderJoin(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var builder strings.Builder
	total := 0
	for _, str := range strs {
		total += len(str)
	}
	builder.Grow(total)
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

// BufferJoin 使用 bytes.Buffer 拼接字符串。
func BufferJoin(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	total := 0
	for _, str := range strs {
		total += len(str)
	}
	buffer.Grow(total)
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}

// Len 返回字符串的 UTF-8 rune 数量。
func Len(str string) int { return utf8.RuneCountInString(str) }

// Partition 按第一次匹配拆分字符串。
func Partition(str, sep string) (head, match, tail string) {
	head, tail, found := strings.Cut(str, sep)
	if !found {
		return str, "", ""
	}
	return head, str[len(head) : len(head)+len(sep)], tail
}

// LastPartition 按最后一次匹配拆分字符串。
func LastPartition(str, sep string) (head, match, tail string) {
	head, tail, found := strings.CutLast(str, sep)
	if !found {
		return "", "", str
	}
	return head, str[len(head) : len(head)+len(sep)], tail
}

// Insert 按 rune 索引插入字符串。
func Insert(dst, src string, index int) string {
	start, end := sliceRuneRange(dst, index, index)
	return dst[:start] + src + dst[end:]
}

func firstRuneInfo(s string) (r rune, size int) {
	if s == "" {
		return utf8.RuneError, 0
	}
	r, size = utf8.DecodeRuneInString(s)
	return r, size
}

func isNormalizedSpace(str string) bool {
	seenNonSpace := false
	prevSpace := false
	for _, r := range str {
		if unicode.IsSpace(r) {
			if !seenNonSpace || prevSpace || r != ' ' {
				return false
			}
			prevSpace = true
			continue
		}
		seenNonSpace = true
		prevSpace = false
	}
	return !prevSpace
}

func isASCII(str string) bool {
	for i := range len(str) {
		if str[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

func normalizeASCIISpace(str string) string {
	var builder strings.Builder
	builder.Grow(len(str))
	pendingSpace := false
	for i := range len(str) {
		c := str[i]
		if isASCIISpace(c) {
			if builder.Len() > 0 {
				pendingSpace = true
			}
			continue
		}
		if pendingSpace {
			builder.WriteByte(' ')
			pendingSpace = false
		}
		builder.WriteByte(c)
	}
	return builder.String()
}

func isASCIISpace(c byte) bool {
	return c == ' ' || '\t' <= c && c <= '\r'
}

func sliceRuneRange(str string, start, end int) (startPos, endPos int) {
	if start < 0 {
		start = 0
	}
	if end < start {
		end = start
	}
	var size, runeIndex int
	for str != "" {
		if runeIndex == start {
			startPos = endPos
		}
		if runeIndex == end {
			return startPos, endPos
		}
		_, size = utf8.DecodeRuneInString(str)
		str = str[size:]
		endPos += size
		runeIndex++
	}
	if start > runeIndex {
		startPos = endPos
	}
	return startPos, endPos
}
