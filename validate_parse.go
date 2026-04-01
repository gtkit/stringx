package stringx

import (
	"time"

	"github.com/gtkit/stringx/maskx"
	"github.com/gtkit/stringx/validatex"
)

// Gender 表示身份证解析出的性别信息。
type Gender = validatex.Gender

const (
	GenderUnknown = validatex.GenderUnknown
	GenderMale    = validatex.GenderMale
	GenderFemale  = validatex.GenderFemale
)

// ChinaIDCardInfo 表示中国身份证结构化解析结果。
type ChinaIDCardInfo = validatex.ChinaIDCardInfo

// ParseChinaIDCard 解析中国身份证的结构化信息。
func ParseChinaIDCard(text string, now time.Time) (*ChinaIDCardInfo, error) {
	return validatex.ParseChinaIDCard(text, now)
}

// MaskChinaIDCard 对中国身份证号码做脱敏处理。
func MaskChinaIDCard(text string) string {
	return maskx.MaskChinaIDCard(text)
}

// MaskBankCard 对银行卡号做脱敏处理。
func MaskBankCard(text string) string {
	return maskx.MaskBankCard(text)
}

// MaskPhone 对手机号做脱敏处理。
func MaskPhone(text string) string {
	return maskx.MaskPhone(text)
}

// MaskEmail 对邮箱地址做脱敏处理。
func MaskEmail(text string) string {
	return maskx.MaskEmail(text)
}

// MaskName 对姓名做脱敏处理。
func MaskName(text string) string {
	return maskx.MaskName(text)
}
