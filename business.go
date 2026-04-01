package stringx

import "github.com/gtkit/stringx/validatex"

// ValidateChinaIDCard 验证 18 位中国居民身份证号码。
func ValidateChinaIDCard(text string) error { return validatex.ValidateChinaIDCard(text) }

// IsChinaIDCard 判断字符串是否为合法的 18 位中国居民身份证号码。
func IsChinaIDCard(text string) bool { return validatex.IsChinaIDCard(text) }

// ValidateChinaUSCC 验证统一社会信用代码。
func ValidateChinaUSCC(text string) error { return validatex.ValidateChinaUSCC(text) }

// IsChinaUSCC 判断字符串是否为合法统一社会信用代码。
func IsChinaUSCC(text string) bool { return validatex.IsChinaUSCC(text) }

// ValidateBankCard 验证银行卡号是否满足基本长度和 Luhn 校验。
func ValidateBankCard(text string) error { return validatex.ValidateBankCard(text) }

// IsBankCard 判断字符串是否为通过 Luhn 校验的银行卡号。
func IsBankCard(text string) bool { return validatex.IsBankCard(text) }
