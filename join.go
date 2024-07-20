package stringx

import (
	"bytes"
	"strings"
)

// BuilderJoin joins a list of strings into a single string using a strings.Builder.
func BuilderJoin(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

// BufferJoin joins a list of strings into a single string using a bytes.Buffer.
func BufferJoin(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}
