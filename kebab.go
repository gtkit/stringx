package stringx

import "github.com/gtkit/stringx/caseconv"

// ToKebab 将字符串转换为 kebab-case 风格。
func ToKebab(str string) string {
	return caseconv.ToKebab(str)
}
