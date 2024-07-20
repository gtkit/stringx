package stringx

import (
	"strings"
	"unicode"
)

// CamelToSnake 驼峰命名转成蛇形命名，如果不是驼峰命名，则转成对应小写字符串
// UserAgent → user_agent
// User → user.
func CamelToSnake(camelCase string) string {
	if len(camelCase) == 0 {
		return ""
	}
	var result strings.Builder
	for i, c := range camelCase {
		if unicode.IsUpper(c) && i > 0 {
			result.WriteByte('_')
		}
		result.WriteRune(unicode.ToLower(c))
	}
	return result.String()
}

// SnakeToCamel 蛇形命名转成驼峰命名，如果不是蛇形命名，则转成对应小写字符串
// user_agent → UserAgent
// user → user.
func SnakeToCamel(snakeCase string) string {
	if len(snakeCase) == 0 {
		return ""
	}
	var (
		result strings.Builder
		num    int
	)
	for i, c := range snakeCase {
		if c == '_' {
			num = i + 1
			continue
		}
		if i == 0 || i == num {
			result.WriteRune(unicode.ToUpper(c))
		} else {
			result.WriteRune(unicode.ToLower(c))
		}
	}
	return result.String()
}

// LowerFirstCase 将字符串的第一个字母转换为小写
// UserAgent → userAgent.
func LowerFirstCase(input string) string {
	if len(input) == 0 {
		return ""
	}

	return strings.ToLower(string(input[0])) + input[1:]
}

// UpperFirstCase 将首字母转换为大写
// user → User.
func UpperFirstCase(input string) string {
	if len(input) == 0 {
		return input
	}

	return strings.ToUpper(string(input[0])) + input[1:]
}

// FirstLowerCase 获取字符串的首字母小写.
// user → u.
// User → u.
func FirstLowerCase(input string) string {
	if len(input) == 0 {
		return input
	}

	firstChar := []rune(input)[:1]

	return strings.ToLower(string(firstChar[0]))
}

// FirstUpperCase 获取字符串的首字母大写.
// user → U.
// User → U.
func FirstUpperCase(input string) string {
	if len(input) == 0 {
		return input
	}

	firstChar := []rune(input)[:1]

	return strings.ToUpper(string(firstChar[0]))
}

// IsEmpty 判断字符串是否为空.
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty 判断字符串是否不为空.
func IsNotEmpty(str string) bool {
	return len(str) > 0
}

// IsBlank 判断字符串是否为空白.
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// IsNotBlank 判断字符串是否不为空白.
func IsNotBlank(str string) bool {
	return len(strings.TrimSpace(str)) > 0
}

// IsAllBlank 判断字符串是否全部为空白.
func IsAllBlank(strs ...string) bool {
	for _, str := range strs {
		if !IsBlank(str) {
			return false
		}
	}
	return true
}

// IsAllNotBlank 判断字符串是否全部不为空白.
func IsAllNotBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return false
		}
	}
	return true
}
