package textx

import (
	"unicode"
	"unicode/utf8"
)

const minCJKCharacter = '\u3400'

// WordCount 返回字符串中的单词数量。
func WordCount(str string) int {
	var count int
	inWord := false
	for str != "" {
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

	for str != "" {
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
