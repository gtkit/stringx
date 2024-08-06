package stringx

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	rand2 "math/rand/v2"
	"strconv"
	"strings"
	"time"
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
func CamelToUpperSnake(camelCase string) string {
	if len(camelCase) == 0 {
		return ""
	}
	var result strings.Builder
	for i, c := range camelCase {
		if unicode.IsUpper(c) && i > 0 {
			result.WriteByte('_')
		}
		result.WriteRune(unicode.ToUpper(c))
	}
	return result.String()
}

// SnakeToCamel 蛇形命名转成驼峰命名.
// firstUpper 是否首字母大写，默认为true, 转为大驼峰命名.
func SnakeToCamel(snakeCase string, firstUpper ...bool) string {
	if len(snakeCase) == 0 {
		return ""
	}
	if len(firstUpper) == 0 || firstUpper[0] {
		return SnakeToGreatCamel(snakeCase)
	}
	return SnakeToSmallCamel(snakeCase)
}

// SnakeToGreatCamel 蛇形命名转成驼峰命名，如果不是蛇形命名，则转成对应首字母大写字符串
// user_agent → UserAgent
// user → User.
func SnakeToGreatCamel(snakeCase string) string {
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

// SnakeToSmallCamel 将蛇形命名转成小驼峰命名，如果不是蛇形命名，则转成对应小写字符串
// user_agent → userAgent.
// user → user.
func SnakeToSmallCamel(snakeCase string) string {
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

		if i == num && num > 0 {
			result.WriteRune(unicode.ToUpper(c))
		} else {
			result.WriteRune(unicode.ToLower(c))
		}
	}
	return result.String()
}

// LowerFirstCase 将字符串的第一个字母转换为小写
// UserAgent → userAgent.
func LowerFirst(input string) string {
	if len(input) == 0 {
		return ""
	}

	return strings.ToLower(string(input[0])) + input[1:]
}

// UpperFirstCase 将首字母转换为大写
// user → User.
func UpperFirst(input string) string {
	if len(input) == 0 {
		return input
	}

	return strings.ToUpper(string(input[0])) + input[1:]
}

// Ufirst 首字母大写.
func Ufirst(s string) string {
	r := []rune(s)
	if len(s) > 0 && unicode.IsLetter(r[0]) && unicode.IsLower(r[0]) {
		r[0] -= 'a' - 'A'
		return string(r)
	}
	return s
}

// Lfirst 首字母小写.
func Lfirst(s string) string {
	r := []rune(s)
	if len(s) > 0 && unicode.IsLetter(r[0]) && unicode.IsUpper(r[0]) {
		r[0] += 'a' - 'A'
		return string(r)
	}
	return s
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
	return str == ""
}

// IsNotEmpty 判断字符串是否不为空.
func IsNotEmpty(str string) bool {
	return str != ""
}

// IsBlank 判断字符串是否为空白.
func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

// IsNotBlank 判断字符串是否不为空白.
func IsNotBlank(str string) bool {
	return strings.TrimSpace(str) != ""
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

// MillTime 获取当前时间戳(毫秒)字符串.
func MillTime() string {
	milltime := time.Now().UnixMilli()
	return strconv.FormatInt(milltime, 10)
}

// UnixTime 获取当前时间戳(秒)字符串.
func UnixTime() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

// Ternary 工具, 类似于 PHP 的三元运算符.
/**
 * Example:
 	func main() {
	  fmt.Println(Ter(true, 1, 2)) // 1
	  fmt.Println(Ter(false, 1, 2)) // 2
	}.
*/
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}

	return b
}

// RandomString 生成长度为 length 的随机字符串.
func RandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand2.N(len(letters))]
	}
	return string(b)
}

// SecRandomString 生成长度为 length 的安全随机字符串.
func SecRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// 注意：如果你需要确切的字符串长度，请根据base64编码的特性,
	// 调整b的长度，因为base64编码会增加输出长度。
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

// RandomNumber 生成长度为 length 随机数字字符串.
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// Substr 截取字符串.
func Substr(s string, start int, strlength ...int) string {
	charlist := []rune(s)
	l := len(charlist)
	length := 0
	end := 0

	if len(strlength) == 0 {
		length = l
	} else {
		length = strlength[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + length

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

// Slug 空格转换为指定分隔符.
func Slug(str, separator string) string {
	return strings.ReplaceAll(str, " ", separator)
}

// SubByte 截取指定长度的字节.
func SubByte(str string, length int) string {
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

// Char returns a char slice
func Char(str string) []string {
	c := make([]string, 0)
	for _, v := range str {
		c = append(c, string(v))
	}
	return c
}

// Escape 转义字符串.
func Escape(s string) string {
	str := strconv.Quote(s)
	str = strings.Replace(str, "'", "\\'", -1)
	strlist := []rune(str)
	l := len(strlist)
	return Substr(str, 1, l-2)
}
