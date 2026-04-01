package stringx

import "github.com/gtkit/stringx/textx"

// BuilderJoin 使用 strings.Builder 将字符串切片高效拼接为一个字符串。
func BuilderJoin(strs []string) string { return textx.BuilderJoin(strs) }

// BufferJoin 使用 bytes.Buffer 将字符串切片高效拼接为一个字符串。
func BufferJoin(strs []string) string { return textx.BufferJoin(strs) }
