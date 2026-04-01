package textx

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const bufferMaxInitGrowSize = 2048
const minCJKCharacter = '\u3400'

type stringBuilder = strings.Builder

// LowerFirst 将字符串首个 rune 转换为小写。
func LowerFirst(input string) string {
	r, size := firstRuneInfo(input)
	if size == 0 {
		return ""
	}
	return string(unicode.ToLower(r)) + input[size:]
}

// UpperFirst 将字符串首个 rune 转换为大写。
func UpperFirst(input string) string {
	r, size := firstRuneInfo(input)
	if size == 0 {
		return input
	}
	return string(unicode.ToUpper(r)) + input[size:]
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
func IsBlank(str string) bool { return strings.TrimSpace(str) == "" }

// IsNotBlank 判断字符串在去除首尾空白后是否仍非空。
func IsNotBlank(str string) bool { return strings.TrimSpace(str) != "" }

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
	return strings.Join(strings.Fields(str), " ")
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

	bs := []byte(str)[:length]
	continuationBytes := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			continuationBytes++
		case bs[i] >= 192 && bs[i] <= 253:
			expectedLength := 0
			switch {
			case bs[i]&252 == 252:
				expectedLength = 6
			case bs[i]&248 == 248:
				expectedLength = 5
			case bs[i]&240 == 240:
				expectedLength = 4
			case bs[i]&224 == 224:
				expectedLength = 3
			default:
				expectedLength = 2
			}
			if continuationBytes+1 == expectedLength {
				return string(bs[:i+expectedLength])
			}
			return string(bs[:i])
		}
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
	quoted = strings.ReplaceAll(quoted, "'", "\\'")
	runes := []rune(quoted)
	return Substr(quoted, 1, len(runes)-2)
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
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		width += RuneWidth(r)
		str = str[size:]
	}
	return width
}

// WordCount 返回字符串中的单词数量。
func WordCount(str string) int {
	var count int
	inWord := false
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		switch {
		case isAlphabet(r):
			if !inWord {
				inWord = true
				count++
			}
		case inWord && (r == '\'' || r == '-'):
		default:
			inWord = false
		}
		str = str[size:]
	}
	return count
}

// WordSplit 将字符串拆分为单词切片。
func WordSplit(str string) []string {
	var word string
	var words []string
	var pos int
	inWord := false

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		switch {
		case isAlphabet(r):
			if !inWord {
				inWord = true
				word = str
				pos = 0
			}
		case inWord && (r == '\'' || r == '-'):
		default:
			if inWord {
				inWord = false
				words = append(words, word[:pos])
			}
		}
		pos += size
		str = str[size:]
	}
	if inWord {
		words = append(words, word[:pos])
	}
	return words
}

// Partition 按第一次匹配拆分字符串。
func Partition(str, sep string) (head, match, tail string) {
	index := strings.Index(str, sep)
	if index == -1 {
		head = str
		return
	}
	head = str[:index]
	match = str[index : index+len(sep)]
	tail = str[index+len(sep):]
	return
}

// LastPartition 按最后一次匹配拆分字符串。
func LastPartition(str, sep string) (head, match, tail string) {
	index := strings.LastIndex(str, sep)
	if index == -1 {
		tail = str
		return
	}
	head = str[:index]
	match = str[index : index+len(sep)]
	tail = str[index+len(sep):]
	return
}

// Insert 按 rune 索引插入字符串。
func Insert(dst, src string, index int) string {
	start, end := sliceRuneRange(dst, index, index)
	return dst[:start] + src + dst[end:]
}

// LeftJustify 在右侧填充字符串。
func LeftJustify(str string, length int, pad string) string {
	l := Len(str)
	if l >= length || pad == "" {
		return str
	}
	remains := length - l
	padLen := Len(pad)
	output := &stringBuilder{}
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
	output := &stringBuilder{}
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
	output := &stringBuilder{}
	output.Grow(len(str) + (remains/padLen+1)*len(pad))
	writePadString(output, pad, padLen, remains/2)
	output.WriteString(str)
	writePadString(output, pad, padLen, (remains+1)/2)
	return output.String()
}

type runeRangeMap struct {
	FromLo rune
	FromHi rune
	ToLo   rune
	ToHi   rune
}

type runeDict struct {
	Dict [unicode.MaxASCII + 1]rune
}

type runeMap map[rune]rune

// Translator 用于复用预编译字符映射规则。
type Translator struct {
	quickDict  *runeDict
	runeMap    runeMap
	ranges     []*runeRangeMap
	mappedRune rune
	reverted   bool
	hasPattern bool
}

// NewTranslator 创建字符转换器。
func NewTranslator(from, to string) *Translator {
	tr := &Translator{}
	if from == "" {
		return tr
	}
	reverted := from[0] == '^'
	deletion := len(to) == 0
	if reverted {
		from = from[1:]
	}

	var fromStart, fromEnd, fromRangeStep rune
	var toStart, toEnd, toRangeStep rune
	var fromRangeSize, toRangeSize rune
	var singleRunes []rune

	updateRange := func() {
		if toEnd == utf8.RuneError {
			return
		}
		if toRangeStep == 0 {
			to, toStart, toEnd, toRangeStep = nextRuneRange(to, toEnd)
			return
		}
		if toStart != toEnd {
			toStart += toRangeStep
			return
		}
		if to == "" {
			toEnd = utf8.RuneError
			return
		}
		to, toStart, toEnd, toRangeStep = nextRuneRange(to, utf8.RuneError)
	}

	if deletion {
		toStart = utf8.RuneError
		toEnd = utf8.RuneError
	} else if reverted {
		var size int
		for len(to) > 0 {
			toStart, size = utf8.DecodeRuneInString(to)
			to = to[size:]
		}
		toEnd = utf8.RuneError
	} else {
		to, toStart, toEnd, toRangeStep = nextRuneRange(to, utf8.RuneError)
	}

	fromEnd = utf8.RuneError
	for len(from) > 0 {
		from, fromStart, fromEnd, fromRangeStep = nextRuneRange(from, fromEnd)
		if fromRangeStep == 0 {
			singleRunes = tr.addRune(fromStart, toStart, singleRunes)
			updateRange()
			continue
		}
		for toEnd != utf8.RuneError && fromStart != fromEnd {
			if toRangeStep == 0 {
				singleRunes = tr.addRune(fromStart, toStart, singleRunes)
				updateRange()
				fromStart += fromRangeStep
				continue
			}
			fromRangeSize = (fromEnd - fromStart) * fromRangeStep
			toRangeSize = (toEnd - toStart) * toRangeStep
			if fromRangeSize > toRangeSize {
				fromStart, toStart = tr.addRuneRange(fromStart, fromStart+toRangeSize*fromRangeStep, toStart, toEnd, singleRunes)
				fromStart += fromRangeStep
				updateRange()
				if fromStart == fromEnd {
					singleRunes = tr.addRune(fromStart, toStart, singleRunes)
					updateRange()
				}
				continue
			}
			fromStart, toStart = tr.addRuneRange(fromStart, fromEnd, toStart, toStart+fromRangeSize*toRangeStep, singleRunes)
			updateRange()
			break
		}
		if fromStart == fromEnd {
			fromEnd = utf8.RuneError
			continue
		}
		_, toStart = tr.addRuneRange(fromStart, fromEnd, toStart, toStart, singleRunes)
		fromEnd = utf8.RuneError
	}
	if fromEnd != utf8.RuneError {
		tr.addRune(fromEnd, toStart, singleRunes)
	}
	tr.reverted = reverted
	tr.mappedRune = -1
	tr.hasPattern = true
	if deletion || reverted {
		tr.mappedRune = toStart
	}
	return tr
}

// Translate 按规则转换字符串。
func (tr *Translator) Translate(str string) string {
	if !tr.hasPattern || str == "" {
		return str
	}
	orig := str
	var output *stringBuilder
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		r, needTranslate := tr.TranslateRune(r)
		if needTranslate && output == nil {
			output = allocBuffer(orig, str)
		}
		if r != utf8.RuneError && output != nil {
			output.WriteRune(r)
		}
		str = str[size:]
	}
	if output == nil {
		return orig
	}
	return output.String()
}

// TranslateRune 转换单个 rune。
func (tr *Translator) TranslateRune(r rune) (result rune, translated bool) {
	switch {
	case tr.quickDict != nil:
		if r <= unicode.MaxASCII {
			result = tr.quickDict.Dict[r]
			if result != 0 {
				translated = true
				if tr.mappedRune >= 0 {
					result = tr.mappedRune
				}
				break
			}
		}
		fallthrough
	case tr.runeMap != nil:
		var ok bool
		if result, ok = tr.runeMap[r]; ok {
			translated = true
			if tr.mappedRune >= 0 {
				result = tr.mappedRune
			}
			break
		}
		fallthrough
	default:
		for i := len(tr.ranges) - 1; i >= 0; i-- {
			rrm := tr.ranges[i]
			if rrm.FromLo <= r && r <= rrm.FromHi {
				translated = true
				if tr.mappedRune >= 0 {
					result = tr.mappedRune
					break
				}
				switch {
				case rrm.ToLo < rrm.ToHi:
					result = rrm.ToLo + r - rrm.FromLo
				case rrm.ToLo > rrm.ToHi:
					result = rrm.ToLo - r + rrm.FromLo
				default:
					result = rrm.ToLo
				}
				break
			}
		}
	}
	if tr.reverted {
		if !translated {
			result = tr.mappedRune
		}
		translated = !translated
	}
	if !translated {
		result = r
	}
	return
}

// HasPattern 返回是否存在有效模式。
func (tr *Translator) HasPattern() bool { return tr.hasPattern }

// Translate 按模式转换字符串。
func Translate(str, from, to string) string {
	tr := NewTranslator(from, to)
	return tr.Translate(str)
}

// Delete 删除命中 pattern 的字符。
func Delete(str, pattern string) string {
	tr := NewTranslator(pattern, "")
	return tr.Translate(str)
}

// Count 返回命中 pattern 的字符数量。
func Count(str, pattern string) int {
	if pattern == "" || str == "" {
		return 0
	}
	tr := NewTranslator(pattern, "")
	count := 0
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		str = str[size:]
		if _, matched := tr.TranslateRune(r); matched {
			count++
		}
	}
	return count
}

// Squeeze 压缩连续重复字符。
func Squeeze(str, pattern string) string {
	var last rune = -1
	var skipSqueeze bool
	var tr *Translator
	var output *stringBuilder
	orig := str
	if pattern != "" {
		tr = NewTranslator(pattern, "")
	}
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		if last == r && !skipSqueeze {
			if tr != nil {
				if _, matched := tr.TranslateRune(r); !matched {
					skipSqueeze = true
				}
			}
			if output == nil {
				output = allocBuffer(orig, str)
			}
			if skipSqueeze {
				output.WriteRune(r)
			}
		} else {
			if output != nil {
				output.WriteRune(r)
			}
			last = r
			skipSqueeze = false
		}
		str = str[size:]
	}
	if output == nil {
		return orig
	}
	return output.String()
}

func allocBuffer(orig, cur string) *stringBuilder {
	output := &stringBuilder{}
	maxSize := len(orig) * 4
	if maxSize > bufferMaxInitGrowSize {
		maxSize = bufferMaxInitGrowSize
	}
	output.Grow(maxSize)
	output.WriteString(orig[:len(orig)-len(cur)])
	return output
}

func firstRuneInfo(s string) (rune, int) {
	if s == "" {
		return utf8.RuneError, 0
	}
	r, size := utf8.DecodeRuneInString(s)
	return r, size
}

func isAlphabet(r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}
	switch {
	case r < minCJKCharacter:
		return true
	case r >= '\u4E00' && r <= '\u9FCC':
		return false
	case r >= '\u3400' && r <= '\u4D85':
		return false
	case r >= '\U00020000' && r <= '\U0002B81D':
		return false
	}
	return true
}

func sliceRuneRange(str string, start, end int) (startPos, endPos int) {
	if start < 0 {
		start = 0
	}
	if end < start {
		end = start
	}
	var size, runeIndex int
	for len(str) > 0 {
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

func writePadString(output *stringBuilder, pad string, padLen, remains int) {
	repeats := remains / padLen
	for i := 0; i < repeats; i++ {
		output.WriteString(pad)
	}
	remains %= padLen
	for i := 0; i < remains; i++ {
		r, size := utf8.DecodeRuneInString(pad)
		output.WriteRune(r)
		pad = pad[size:]
	}
}

func (tr *Translator) addRune(from, to rune, singleRunes []rune) []rune {
	if from <= unicode.MaxASCII {
		if tr.quickDict == nil {
			tr.quickDict = &runeDict{}
		}
		tr.quickDict.Dict[from] = to
	} else {
		if tr.runeMap == nil {
			tr.runeMap = make(runeMap)
		}
		tr.runeMap[from] = to
	}
	return append(singleRunes, from)
}

func (tr *Translator) addRuneRange(fromLo, fromHi, toLo, toHi rune, singleRunes []rune) (rune, rune) {
	var rrm *runeRangeMap
	if fromLo < fromHi {
		rrm = &runeRangeMap{FromLo: fromLo, FromHi: fromHi, ToLo: toLo, ToHi: toHi}
	} else {
		rrm = &runeRangeMap{FromLo: fromHi, FromHi: fromLo, ToLo: toHi, ToHi: toLo}
	}
	for _, r := range singleRunes {
		if rrm.FromLo <= r && r <= rrm.FromHi {
			if r <= unicode.MaxASCII {
				tr.quickDict.Dict[r] = 0
			} else {
				delete(tr.runeMap, r)
			}
		}
	}
	tr.ranges = append(tr.ranges, rrm)
	return fromHi, toHi
}

func nextRuneRange(str string, last rune) (remaining string, start, end rune, rangeStep rune) {
	remaining = str
	escaping := false
	isRange := false
	for len(remaining) > 0 {
		r, size := utf8.DecodeRuneInString(remaining)
		remaining = remaining[size:]
		if !escaping {
			switch r {
			case '\\':
				escaping = true
				continue
			case '-':
				if last == utf8.RuneError {
					continue
				}
				start = last
				isRange = true
				continue
			}
		}
		escaping = false
		if last != utf8.RuneError {
			if isRange && last == r {
				isRange = false
				continue
			}
			start = last
			end = r
			if isRange {
				if start < end {
					rangeStep = 1
				} else {
					rangeStep = -1
				}
			}
			return
		}
		last = r
	}
	start = last
	end = utf8.RuneError
	return
}
