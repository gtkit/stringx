package validatex

import (
	"time"

	"github.com/gtkit/stringx/validatex/business"
)

// Gender 表示身份证解析后的性别信息。
type Gender string

const (
	GenderUnknown = Gender(business.GenderUnknown)
	GenderMale    = Gender(business.GenderMale)
	GenderFemale  = Gender(business.GenderFemale)
)

// ChinaIDCardInfo 是中国身份证号码的解析结果。
type ChinaIDCardInfo struct {
	Number       string
	ProvinceCode string
	BirthDate    time.Time
	Age          int
	Gender       Gender
}

// ValidateChinaMobilePhone 使用本地规则严格验证中国手机号码。
//
// 支持可选的 +86 或 86 前缀，并允许空格、连字符和括号作为格式化字符。
func ValidateChinaMobilePhone(text string) error { return business.ValidateChinaMobilePhone(text) }

// ValidateChinaIDCard 验证 18 位中国居民身份证。
func ValidateChinaIDCard(text string) error { return business.ValidateChinaIDCard(text) }

// IsChinaIDCard 判断字符串是否为合法中国身份证。
func IsChinaIDCard(text string) bool { return business.IsChinaIDCard(text) }

// ValidateChinaUSCC 验证统一社会信用代码。
func ValidateChinaUSCC(text string) error { return business.ValidateChinaUSCC(text) }

// IsChinaUSCC 判断字符串是否为合法统一社会信用代码。
func IsChinaUSCC(text string) bool { return business.IsChinaUSCC(text) }

// ValidateBankCard 使用 Luhn 算法验证银行卡号。
func ValidateBankCard(text string) error { return business.ValidateBankCard(text) }

// IsBankCard 判断字符串是否为通过 Luhn 校验的银行卡号。
func IsBankCard(text string) bool { return business.IsBankCard(text) }

// ParseChinaIDCard 解析中国身份证号码中的结构化信息。
func ParseChinaIDCard(text string, now time.Time) (*ChinaIDCardInfo, error) {
	info, err := business.ParseChinaIDCard(text, now)
	if err != nil {
		return nil, err
	}
	return &ChinaIDCardInfo{
		Number:       info.Number,
		ProvinceCode: info.ProvinceCode,
		BirthDate:    info.BirthDate,
		Age:          info.Age,
		Gender:       Gender(info.Gender),
	}, nil
}
