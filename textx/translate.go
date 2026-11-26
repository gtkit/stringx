package textx

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const bufferMaxInitGrowSize = 2048

type runeRangeMap struct {
	FromLo rune
	FromHi rune
	ToLo   rune
	ToHi   rune
}

type runeDict struct {
	Dict  [unicode.MaxASCII + 1]rune
	Match [unicode.MaxASCII + 1]bool
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
	deletion := to == ""
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
		for to != "" {
			toStart, size = utf8.DecodeRuneInString(to)
			to = to[size:]
		}
		toEnd = utf8.RuneError
	} else {
		to, toStart, toEnd, toRangeStep = nextRuneRange(to, utf8.RuneError)
	}

	fromEnd = utf8.RuneError
	for from != "" {
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
	var output *strings.Builder
	for str != "" {
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
			if tr.quickDict.Match[r] {
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
	return result, translated
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
	for str != "" {
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
	var output *strings.Builder
	orig := str
	if pattern != "" {
		tr = NewTranslator(pattern, "")
	}
	for str != "" {
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

func allocBuffer(orig, cur string) *strings.Builder {
	output := &strings.Builder{}
	maxSize := len(orig) * 4
	if maxSize > bufferMaxInitGrowSize {
		maxSize = bufferMaxInitGrowSize
	}
	output.Grow(maxSize)
	output.WriteString(orig[:len(orig)-len(cur)])
	return output
}

func (tr *Translator) addRune(from, to rune, singleRunes []rune) []rune {
	if from <= unicode.MaxASCII {
		if tr.quickDict == nil {
			tr.quickDict = &runeDict{}
		}
		tr.quickDict.Dict[from] = to
		tr.quickDict.Match[from] = true
	} else {
		if tr.runeMap == nil {
			tr.runeMap = make(runeMap)
		}
		tr.runeMap[from] = to
	}
	return append(singleRunes, from)
}

func (tr *Translator) addRuneRange(fromLo, fromHi, toLo, toHi rune, singleRunes []rune) (from, to rune) {
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
				tr.quickDict.Match[r] = false
			} else {
				delete(tr.runeMap, r)
			}
		}
	}
	tr.ranges = append(tr.ranges, rrm)
	return fromHi, toHi
}

func nextRuneRange(str string, last rune) (remaining string, start, end, rangeStep rune) {
	remaining = str
	escaping := false
	isRange := false
	for remaining != "" {
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
			return remaining, start, end, rangeStep
		}
		last = r
	}
	start = last
	end = utf8.RuneError
	return remaining, start, end, rangeStep
}
