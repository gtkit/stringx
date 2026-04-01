package stringx

import "github.com/gtkit/stringx/caseconv"

// ToSnake 将字符串转换为 snake_case 风格。
func ToSnake(str string) string {
	return caseconv.ToSnake(str)
}

// ToUpperSnake 将字符串转换为 UPPER_SNAKE_CASE 风格。
func ToUpperSnake(str string) string {
	return caseconv.ToUpperSnake(str)
}
