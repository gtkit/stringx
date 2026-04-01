package stringx

import "github.com/gtkit/stringx/caseconv"

// ToCamel 将字符串转换为大驼峰命名风格。
func ToCamel(s string) string {
	return caseconv.ToCamel(s)
}

// ToLowerCamel 将字符串转换为小驼峰命名风格。
func ToLowerCamel(s string) string {
	return caseconv.ToLowerCamel(s)
}
