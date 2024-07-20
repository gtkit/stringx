package test

import (
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
