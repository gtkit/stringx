package test_test

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"sync"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/gtkit/stringx"
)

func TestCamelToSnake(t *testing.T) {
	str := "CamelToSnake"
	// str := "camelToSnake"
	newStr := stringx.CamelToSnake(str)
	t.Log("newStr:", newStr)
}

func TestSnakeToCamel(t *testing.T) {
	// str := "snake_to_camel_case"
	str := "Snake"
	newStr := stringx.SnakeToCamel(str, false)
	t.Log("newStr:", newStr)
	snakeStr := strings.ToTitle(str)
	t.Log("snakeStr:", snakeStr)
}

func TestSnakeToSmallCamel(t *testing.T) {
	str := "Snake_to_camel_case"
	// str := "Snake"
	newStr := stringx.SnakeToSmallCamel(str)
	t.Log("newStr:", newStr)
	// snakeStr := strings.ToTitle()
}

func TestLowerFirstCase(t *testing.T) {
	str := "BigCamelCaseCase"
	newStr1 := stringx.LowerFirst(str)
	newStr2 := stringx.Lfirst(str)

	t.Log("newStr 1:", newStr1)
	t.Log("newStr 2:", newStr2)
}

func TestUpperFirstCase(t *testing.T) {
	str := "upperFirstCase"
	newStr1 := stringx.UpperFirst(str)
	newStr2 := stringx.Ufirst(str)

	t.Log("newStr 1:", newStr1)
	t.Log("newStr 2:", newStr2)
}

func TestFirstLowerCase(t *testing.T) {
	str := "FirstLowerCase"
	newStr := stringx.FirstLowerCase(str)

	t.Log("newStr:", newStr)
}

func TestFirstUpperCase(t *testing.T) {
	str := "firstUpperCase"
	newStr := stringx.FirstUpperCase(str)

	t.Log("newStr:", newStr)
}

func TestIsBlank(t *testing.T) {
	str := " "
	if stringx.IsBlank(str) {
		t.Log("str is blank")
	} else {
		t.Log("str is not blank")
	}
}

func TestIsNotBlank(t *testing.T) {
	str := "not blank"
	if stringx.IsNotBlank(str) {
		t.Log("str is not blank")
	} else {
		t.Log("str is blank")
	}
}

func TestBuilderJoin(t *testing.T) {
	strs := []string{"hello", " ", "world ", "! ", " 你好"}
	builder := stringx.BuilderJoin(strs)
	t.Log("builder:", builder)
}

func TestBufferJoin(t *testing.T) {
	strs := []string{"hello", " ", "world ", "! ", " 你好"}
	buffer := stringx.BufferJoin(strs)
	t.Log("buffer:", buffer)
}

func TestRandStr(t *testing.T) {
	// str, err := SecureRandomString(10)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		str := stringx.RandomString(5)
		t.Log("---str1: ", str)
	}()
	go func() {
		defer wg.Done()
		str := stringx.RandomString(0)
		t.Log("---str2: ", str)
	}()

	go func() {
		defer wg.Done()
		str := stringx.RandomString(10)
		t.Log("---str3: ", str)
	}()
	// str := stringx.RandomNumber(5)
	// if err != nil {
	// 	t.Error(err)
	// }
	wg.Wait()
}

func SecureRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// 注意：如果你需要确切的字符串长度，请根据base64编码的特性,
	// 调整b的长度，因为base64编码会增加输出长度。
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}

func TestSlug(t *testing.T) {
	str := "Laravel 5 Framework"

	slug := stringx.Slug(str, "-")
	t.Log("slug:", slug)
}

func toLower(s string) string {
	r := []rune(s) // 将字符串转换为rune切片

	for i := 0; i < len(r); i++ {
		if 'A' <= r[i] && r[i] <= 'Z' { // 判断是否为大写字母
			r[i] += 'a' - 'A' // 转换为小写

		}
	}
	return string(r) // 将rune切片转换回字符串
}

func TestToLower(t *testing.T) {
	str := "HELLO WORLD"
	newStr := toLower(str)
	t.Log("newStr:", newStr)
}

func toUpper(s string) string {
	r := []rune(s) // 将字符串转换为rune切片
	for i := 0; i < len(r); i++ {
		if 'a' <= r[i] && r[i] <= 'z' { // 判断是否为小写字母
			r[i] -= 'a' - 'A' // 转换为大写
		}
	}
	return string(r) // 将rune切片转换回字符串
}

func TestSubByte(t *testing.T) {
	str := "hello world"
	newStr := stringx.SubByte(str, 3)
	t.Log("newStr:", newStr)
}

func TestUnicode(t *testing.T) {
	rune1 := ' '
	rune2 := 'h'

	t.Logf("[%c] 是空格? %t\n", rune1, unicode.IsSpace(rune1))
	// 输出 [ ] 是空格? true
	t.Logf("[%c] 是空格? %t\n\n", rune2, unicode.IsSpace(rune2))
	// 输出 [h] 是空格? false

	d1 := '1'
	d2 := 'w'
	t.Logf("[%c] 是十进制数? %t\n", d1, unicode.IsDigit(d1))
	// 输出: [1] 是十进制数? true
	t.Logf("[%c] 是十进制数? %t\n\n", d2, unicode.IsDigit(d2))
	// 输出: [w] 是十进制数? false

	t.Logf("[%c] 是数字? %t\n", d1, unicode.IsNumber(d1))
	// 输出: [1] 是数字? true
	t.Logf("[%c] 是数字? %t\n\n", d2, unicode.IsNumber(d2))
	// 输出: [w] 是数字? false

	str1 := '刘'
	str2 := 'l'
	str3 := 'W'
	str4 := '!'
	str5 := '_'
	str6 := '-'

	t.Logf("[%c] 是字母? %t\n", str1, unicode.IsLetter(str1))
	// [刘] 是字母? true
	t.Logf("[%c] 是字母? %t\n", str2, unicode.IsLetter(str2))
	// [l] 是字母? true
	t.Logf("[%c] 是字母? %t\n", str3, unicode.IsLetter(str3))
	// [W] 是字母? true
	t.Logf("[%c] 是字母? %t\n\n", str4, unicode.IsLetter(str4))
	// [!] 是字母? false

	t.Logf("[%c] 是标点符号? %t\n", str1, unicode.IsPunct(str1))
	// [刘] 是标点符号? false
	t.Logf("[%c] 是标点符号? %t\n", str2, unicode.IsPunct(str2))
	// [l] 是标点符号? false
	t.Logf("[%c] 是标点符号? %t\n", str3, unicode.IsPunct(str3))
	// [W] 是标点符号? false
	t.Logf("[%c] 是标点符号? %t\n", str4, unicode.IsPunct(str4))
	// [!] 是标点符号? true
	t.Logf("[%c] 是标点符号? %t\n", str5, unicode.IsPunct(str5))
	// [_] 是标点符号? true
	t.Logf("[%c] 是标点符号? %t\n\n", str6, unicode.IsPunct(str6))
	// [-] 是标点符号? true

	t.Logf("[%c] 是小写字母? %t\n", str1, unicode.IsLower(str1))
	// [刘] 是小写字母? false
	t.Logf("[%c] 是小写字母? %t\n", str2, unicode.IsLower(str2))
	// [l] 是小写字母? true
	t.Logf("[%c] 是小写字母? %t\n", str3, unicode.IsLower(str3))
	// [W] 是小写字母? false
	t.Logf("[%c] 是小写字母? %t\n\n", str4, unicode.IsLower(str4))
	// [!] 是小写字母? false

	t.Logf("[%c] 是大写字母? %t\n", str1, unicode.IsUpper(str1))
	// [刘] 是大写字母? false
	t.Logf("[%c] 是大写字母? %t\n", str2, unicode.IsUpper(str2))
	// [l] 是大写字母? false
	t.Logf("[%c] 是大写字母? %t\n", str3, unicode.IsUpper(str3))
	// [W] 是大写字母? true
	t.Logf("[%c] 是大写字母? %t\n\n", str4, unicode.IsUpper(str4))
	// [!] 是大写字母? false

	t.Logf("[%c] 是汉字? %t\n", str1, unicode.Is(unicode.Scripts["Han"], str1))
	// [刘] 是汉字? true
	t.Logf("[%c] 是汉字? %t\n", str2, unicode.Is(unicode.Scripts["Han"], str2))
	// [l] 是汉字? false
	t.Logf("[%c] 是汉字? %t\n\n", str3, unicode.Is(unicode.Scripts["Han"], str3))
	// [!] 是汉字? false

	str1 = 'W'
	t.Logf("[%c] 转成小写: %c \n\n", str1, unicode.ToLower(str1))
	// [W] 转成小写: w

	str2 = 'a'
	t.Logf("[%c] 转成大写: %c \n\n", str2, unicode.ToUpper(str2))
	// [a] 转成大写: A

	var char1 rune = 'A'
	var char2 rune = '中'

	t.Logf("Character 1: %c\n", char1)
	t.Logf("Character 2: %c\n", char2)

	{
		s := "hello"
		t.Logf("len(%s) = %d", s, len(s))
	}
	{
		s := "你好"
		t.Logf("len(%s) = %d", s, len(s))

	}

	{
		s := "hello"
		t.Logf("unicodeLen(%s) = %d", s, utf8.RuneCountInString(s))
	}
	{
		s := "你好"
		t.Logf("unicodeLen(%s) = %d", s, utf8.RuneCountInString(s))
	}
	{
		s := " "
		t.Logf("unicodeLen(%s) = %d", s, utf8.RuneCountInString(s))
	}

	{
		s := '福'
		t.Logf("runelen(%c) = %d", s, utf8.RuneLen(s))
	}
}
