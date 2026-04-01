package stringx

import "github.com/gtkit/stringx/textx"

type Translator = textx.Translator

// NewTranslator 根据 from 和 to 模式创建一个可复用的字符转换器。
func NewTranslator(from, to string) *Translator { return textx.NewTranslator(from, to) }

// Translate 根据 from 和 to 定义的模式转换字符串中的字符。
func Translate(str, from, to string) string { return textx.Translate(str, from, to) }

// Delete 删除字符串中所有命中 pattern 模式的字符。
func Delete(str, pattern string) string { return textx.Delete(str, pattern) }

// Count 返回字符串中命中 pattern 模式的字符数量。
func Count(str, pattern string) int { return textx.Count(str, pattern) }

// Squeeze 压缩字符串中连续重复的字符。
func Squeeze(str, pattern string) string { return textx.Squeeze(str, pattern) }
