package stringx

import "github.com/gtkit/stringx/textx"

// Partition 按照 sep 第一次出现的位置将字符串拆分为三段。
func Partition(str, sep string) (head, match, tail string) { return textx.Partition(str, sep) }

// LastPartition 按照 sep 最后一次出现的位置将字符串拆分为三段。
func LastPartition(str, sep string) (head, match, tail string) { return textx.LastPartition(str, sep) }

// Insert 按 rune 索引将 src 插入到 dst 中。
func Insert(dst, src string, index int) string { return textx.Insert(dst, src, index) }

// WordSplit 将字符串拆分为单词切片。
func WordSplit(str string) []string { return textx.WordSplit(str) }
