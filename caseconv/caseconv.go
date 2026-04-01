package caseconv

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// ToCamel 将字符串转换为大驼峰格式。
func ToCamel(s string) string {
	return toConversionCamel(s, true)
}

// ToLowerCamel 将字符串转换为小驼峰格式。
func ToLowerCamel(s string) string {
	return toConversionCamel(s, false)
}

// ToSnake 将字符串转换为 snake_case。
func ToSnake(s string) string {
	return toDelimitedCase(s, '_', false)
}

// ToUpperSnake 将字符串转换为 UPPER_SNAKE_CASE。
func ToUpperSnake(s string) string {
	return toDelimitedCase(s, '_', true)
}

// ToKebab 将字符串转换为 kebab-case。
func ToKebab(s string) string {
	return toDelimitedCase(s, '-', false)
}

func toConversionCamel(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	var result strings.Builder
	result.Grow(len(s))

	capNext := initCase
	prevIsCap := false

	for i, r := range s {
		rIsCap := unicode.IsUpper(r)
		rIsLow := unicode.IsLower(r)

		if capNext {
			if rIsLow {
				r = unicode.ToUpper(r)
			}
		} else if i == 0 {
			if rIsCap {
				r = unicode.ToLower(r)
			}
		} else if prevIsCap && rIsCap {
			r = unicode.ToLower(r)
		}

		prevIsCap = rIsCap

		switch {
		case rIsCap || rIsLow:
			result.WriteRune(r)
			capNext = false
		case unicode.IsDigit(r):
			result.WriteRune(r)
			capNext = true
		default:
			capNext = isCaseSeparator(r)
		}
	}

	return result.String()
}

type runeClass uint8

const (
	runeClassUnknown runeClass = iota
	runeClassLower
	runeClassUpper
	runeClassDigit
	runeClassLetter
)

func classifyRune(r rune) runeClass {
	switch {
	case unicode.IsLower(r):
		return runeClassLower
	case unicode.IsUpper(r):
		return runeClassUpper
	case unicode.IsDigit(r):
		return runeClassDigit
	case unicode.IsLetter(r):
		return runeClassLetter
	default:
		return runeClassUnknown
	}
}

func isCaseSeparator(r rune) bool {
	return r == '_' || r == '-' || r == '.' || unicode.IsSpace(r)
}

func shouldInsertDelimiter(prevClass, currClass, nextClass runeClass) bool {
	switch {
	case prevClass == runeClassLower && (currClass == runeClassUpper || currClass == runeClassDigit):
		return true
	case prevClass == runeClassUpper && currClass == runeClassDigit:
		return true
	case prevClass == runeClassUpper && currClass == runeClassUpper && nextClass == runeClassLower:
		return true
	case prevClass == runeClassDigit && (currClass == runeClassLower || currClass == runeClassUpper || currClass == runeClassLetter):
		return true
	case prevClass == runeClassLetter && currClass == runeClassDigit:
		return true
	}

	return false
}

func toDelimitedCase(s string, separator rune, upper bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(s) + len(s)/2)

	var prevClass runeClass
	pendingSeparator := false
	lastWasSeparator := false

	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if isCaseSeparator(r) {
			if builder.Len() > 0 {
				pendingSeparator = true
			}
			lastWasSeparator = builder.Len() == 0 || lastWasSeparator
			s = s[size:]
			continue
		}

		currClass := classifyRune(r)
		nextClass := nextRuneClass(s[size:])

		if builder.Len() > 0 && (pendingSeparator || shouldInsertDelimiter(prevClass, currClass, nextClass)) && !lastWasSeparator {
			builder.WriteRune(separator)
			lastWasSeparator = true
		}

		if upper {
			builder.WriteRune(unicode.ToUpper(r))
		} else {
			builder.WriteRune(unicode.ToLower(r))
		}

		prevClass = currClass
		pendingSeparator = false
		lastWasSeparator = false
		s = s[size:]
	}

	return builder.String()
}

func nextRuneClass(s string) runeClass {
	if s == "" {
		return runeClassUnknown
	}

	r, _ := utf8.DecodeRuneInString(s)
	return classifyRune(r)
}
