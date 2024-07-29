package test

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"testing"

	"github.com/gtkit/stringx"
)

func TestCamelToSnake(t *testing.T) {
	str := "CamelToSnake"
	newStr := stringx.CamelToSnake(str)
	t.Log("newStr:", newStr)
}

func TestSnakeToCamel(t *testing.T) {
	str := "snake_to_camel_case"
	newStr := stringx.SnakeToCamel(str)
	t.Log("newStr:", newStr)
	// snakeStr := strings.ToTitle()
}

func TestLowerFirstCase(t *testing.T) {
	str := "BigCamelCaseCase"
	newStr := stringx.LowerFirstCase(str)

	t.Log("newStr:", newStr)
}

func TestUpperFirstCase(t *testing.T) {
	str := "upperFirstCase"
	newStr := stringx.UpperFirstCase(str)

	t.Log("newStr:", newStr)
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
