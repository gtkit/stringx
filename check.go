package stringx

import "regexp"

var (
	emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phonePattern = regexp.MustCompile(`^1[3-9]\d{9}$`)
)

// IsEmail 判断字符串是否符合常见电子邮箱格式。
//
// 该方法适合常规表单校验，不追求完整覆盖 RFC 邮箱规范的所有边缘情况。
// 如果你需要更严格的协议语义校验，请使用 IsEmailStrict。
func IsEmail(email string) bool {
	return emailPattern.MatchString(email)
}

// IsPhone 判断字符串是否符合中国大陆手机号的轻量业务格式。
//
// 该方法基于正则表达式，适合表单级快速校验。
// 如果你需要支持 `+86`、空格、连字符等真实输入，并做更严格的本地规则校验，
// 请使用 `validatex` 子包中的相关能力。
func IsPhone(phone string) bool {
	return phonePattern.MatchString(phone)
}
