package validatex

import (
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/netip"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/idna"
)

// Gender 表示身份证解析后的性别信息。
type Gender string

const (
	GenderUnknown Gender = "unknown"
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
)

// ChinaIDCardInfo 是中国身份证号码的解析结果。
type ChinaIDCardInfo struct {
	Number       string
	ProvinceCode string
	BirthDate    time.Time
	Age          int
	Gender       Gender
}

// ParseEmailAddress 使用 RFC 5322 语法解析邮箱地址。
func ParseEmailAddress(address string) (*mail.Address, error) {
	return mail.ParseAddress(address)
}

// IsEmailStrict 判断字符串是否为严格邮箱格式。
func IsEmailStrict(address string) bool {
	return ValidateEmailStrict(address) == nil
}

// ValidateEmailStrict 验证字符串是否为严格邮箱格式。
func ValidateEmailStrict(address string) error {
	if address == "" {
		return errors.New("email is empty")
	}
	if len(address) > 254 {
		return errors.New("email exceeds maximum length")
	}
	if strings.TrimSpace(address) != address {
		return errors.New("email contains surrounding whitespace")
	}

	at := strings.LastIndexByte(address, '@')
	if at <= 0 || at == len(address)-1 {
		return errors.New("email must contain local part and domain")
	}

	local, domain := address[:at], address[at+1:]
	if strings.Contains(domain, "@") {
		return errors.New("email domain contains extra at sign")
	}
	if len(local) > 64 {
		return errors.New("email local part exceeds maximum length")
	}
	if !isEmailLocalPart(local) {
		return errors.New("email local part is invalid")
	}
	if !isEmailDomain(domain) {
		return errors.New("email domain is invalid")
	}
	return nil
}

// IsURL 判断字符串是否为带 scheme 和 host 的绝对 URL。
func IsURL(raw string) bool {
	return ValidateURL(raw) == nil
}

// ValidateURL 验证字符串是否为带 scheme 和 host 的绝对 URL。
func ValidateURL(raw string) error {
	u, err := parseAbsoluteURL(raw)
	if err != nil {
		return err
	}
	return validateURLHost(u)
}

// IsHTTPURL 判断字符串是否为 http 或 https 绝对 URL。
func IsHTTPURL(raw string) bool {
	return ValidateHTTPURL(raw) == nil
}

// ValidateHTTPURL 验证字符串是否为 http 或 https 绝对 URL。
func ValidateHTTPURL(raw string) error {
	u, err := parseAbsoluteURL(raw)
	if err != nil {
		return err
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("unsupported URL scheme %q", u.Scheme)
	}
	return validateURLHost(u)
}

// IsHostname 判断字符串是否为合法主机名。
func IsHostname(host string) bool {
	return ValidateHostname(host) == nil
}

// ValidateHostname 验证字符串是否为合法主机名。
func ValidateHostname(host string) error {
	_, err := normalizeHostnameASCII(host)
	return err
}

// IsIP 判断字符串是否为合法 IP 地址。
func IsIP(text string) bool {
	return ValidateIP(text) == nil
}

// ValidateIP 验证字符串是否为合法 IP 地址。
func ValidateIP(text string) error {
	if _, err := netip.ParseAddr(text); err != nil {
		return fmt.Errorf("invalid IP address: %w", err)
	}
	return nil
}

// IsIPv4 判断字符串是否为合法 IPv4 地址。
func IsIPv4(text string) bool {
	return ValidateIPv4(text) == nil
}

// ValidateIPv4 验证字符串是否为合法 IPv4 地址。
func ValidateIPv4(text string) error {
	addr, err := netip.ParseAddr(text)
	if err != nil {
		return fmt.Errorf("invalid IPv4 address: %w", err)
	}
	if !addr.Is4() {
		return errors.New("address is not IPv4")
	}
	return nil
}

// IsIPv6 判断字符串是否为合法 IPv6 地址。
func IsIPv6(text string) bool {
	return ValidateIPv6(text) == nil
}

// ValidateIPv6 验证字符串是否为合法 IPv6 地址。
func ValidateIPv6(text string) error {
	addr, err := netip.ParseAddr(text)
	if err != nil {
		return fmt.Errorf("invalid IPv6 address: %w", err)
	}
	if !addr.Is6() {
		return errors.New("address is not IPv6")
	}
	return nil
}

// IsCIDR 判断字符串是否为合法 CIDR。
func IsCIDR(text string) bool {
	return ValidateCIDR(text) == nil
}

// ValidateCIDR 验证字符串是否为合法 CIDR。
func ValidateCIDR(text string) error {
	if _, err := netip.ParsePrefix(text); err != nil {
		return fmt.Errorf("invalid CIDR: %w", err)
	}
	return nil
}

// IsMAC 判断字符串是否为合法 MAC 地址。
func IsMAC(text string) bool {
	return ValidateMAC(text) == nil
}

// ValidateMAC 验证字符串是否为合法 MAC 地址。
func ValidateMAC(text string) error {
	if _, err := net.ParseMAC(text); err != nil {
		return fmt.Errorf("invalid MAC address: %w", err)
	}
	return nil
}

// IsUUID 判断字符串是否为合法 UUID。
func IsUUID(text string) bool {
	return ValidateUUID(text) == nil
}

// ValidateUUID 验证字符串是否为合法 UUID。
func ValidateUUID(text string) error {
	if len(text) != 36 {
		return errors.New("UUID length must be 36")
	}
	for i, r := range text {
		switch i {
		case 8, 13, 18, 23:
			if r != '-' {
				return errors.New("UUID hyphen position is invalid")
			}
		default:
			if !isHexRune(byte(r)) {
				return errors.New("UUID contains non-hex character")
			}
		}
	}
	return nil
}

// IsUUIDv4 判断字符串是否为合法 UUID v4。
func IsUUIDv4(text string) bool {
	return ValidateUUIDv4(text) == nil
}

// ValidateUUIDv4 验证字符串是否为合法 UUID v4。
func ValidateUUIDv4(text string) error {
	if err := ValidateUUID(text); err != nil {
		return err
	}
	if text[14] != '4' {
		return errors.New("UUID version is not v4")
	}
	switch text[19] {
	case '8', '9', 'a', 'A', 'b', 'B':
		return nil
	default:
		return errors.New("UUID variant is invalid")
	}
}

// IsE164 判断字符串是否符合 E.164。
func IsE164(text string) bool {
	return ValidateE164(text) == nil
}

// ValidateE164 验证字符串是否符合 E.164。
func ValidateE164(text string) error {
	if len(text) < 2 || len(text) > 16 {
		return errors.New("E.164 length must be between 2 and 16")
	}
	if text[0] != '+' {
		return errors.New("E.164 number must start with plus sign")
	}
	if text[1] == '0' {
		return errors.New("E.164 country code cannot start with 0")
	}
	for i := 1; i < len(text); i++ {
		if text[i] < '0' || text[i] > '9' {
			return errors.New("E.164 number must contain digits only")
		}
	}
	return nil
}

// ValidateChinaMobilePhone 使用本地规则严格验证中国手机号码。
//
// 支持可选的 `+86` 或 `86` 前缀，并允许空格、连字符和括号作为格式化字符。
func ValidateChinaMobilePhone(text string) error {
	number, err := normalizeChinaMobilePhone(text)
	if err != nil {
		return err
	}
	if _, ok := chinaMobilePrefixes[number[:3]]; !ok {
		return errors.New("phone number prefix is invalid")
	}
	return nil
}

// ValidateChinaIDCard 验证 18 位中国居民身份证。
func ValidateChinaIDCard(text string) error {
	if len(text) != 18 {
		return errors.New("ID card length must be 18")
	}

	text = strings.ToUpper(text)
	for i := 0; i < 17; i++ {
		if text[i] < '0' || text[i] > '9' {
			return errors.New("ID card must contain digits in first 17 positions")
		}
	}
	if !(text[17] >= '0' && text[17] <= '9' || text[17] == 'X') {
		return errors.New("ID card last character must be digit or X")
	}
	if _, ok := chinaProvinceCodes[text[:2]]; !ok {
		return errors.New("ID card province code is invalid")
	}

	birth := text[6:14]
	date, err := time.Parse("20060102", birth)
	if err != nil || date.Format("20060102") != birth {
		return errors.New("ID card birth date is invalid")
	}

	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(text[i]-'0') * chinaIDCardWeights[i]
	}
	expected := chinaIDCardChecksumCodes[sum%11]
	if text[17] != expected {
		return errors.New("ID card checksum is invalid")
	}
	return nil
}

// IsChinaIDCard 判断字符串是否为合法中国身份证。
func IsChinaIDCard(text string) bool {
	return ValidateChinaIDCard(text) == nil
}

// ValidateChinaUSCC 验证统一社会信用代码。
func ValidateChinaUSCC(text string) error {
	if len(text) != 18 {
		return errors.New("USCC length must be 18")
	}

	text = strings.ToUpper(text)
	sum := 0
	for i := 0; i < 17; i++ {
		value, ok := chinaUSCCValueMap[text[i]]
		if !ok {
			return errors.New("USCC contains invalid character")
		}
		sum += value * chinaUSCCWeights[i]
	}
	lastValue, ok := chinaUSCCValueMap[text[17]]
	if !ok {
		return errors.New("USCC contains invalid character")
	}
	check := (31 - sum%31) % 31
	if lastValue != check {
		return errors.New("USCC checksum is invalid")
	}
	return nil
}

// IsChinaUSCC 判断字符串是否为合法统一社会信用代码。
func IsChinaUSCC(text string) bool {
	return ValidateChinaUSCC(text) == nil
}

// ValidateBankCard 使用 Luhn 算法验证银行卡号。
func ValidateBankCard(text string) error {
	if len(text) < 12 || len(text) > 19 {
		return errors.New("bank card length must be between 12 and 19")
	}
	sum := 0
	double := false
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] < '0' || text[i] > '9' {
			return errors.New("bank card must contain digits only")
		}
		digit := int(text[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	if sum%10 != 0 {
		return errors.New("bank card luhn checksum is invalid")
	}
	return nil
}

// IsBankCard 判断字符串是否为通过 Luhn 校验的银行卡号。
func IsBankCard(text string) bool {
	return ValidateBankCard(text) == nil
}

func normalizeChinaMobilePhone(text string) (string, error) {
	if len(text) > 64 {
		return "", errors.New("phone input is too long")
	}

	var compact strings.Builder
	compact.Grow(len(text))

	for _, r := range text {
		switch {
		case '0' <= r && r <= '9':
			compact.WriteRune(r)
		case r == '+' || r == '(' || r == ')':
			compact.WriteRune(r)
		case r == ' ' || r == '-':
		default:
			return "", errors.New("phone input contains unsupported characters")
		}
	}

	normalized := compact.String()
	switch {
	case strings.HasPrefix(normalized, "(+86)"):
		normalized = normalized[len("(+86)"):]
	case strings.HasPrefix(normalized, "(86)"):
		normalized = normalized[len("(86)"):]
	case strings.HasPrefix(normalized, "+86"):
		normalized = normalized[len("+86"):]
	case strings.HasPrefix(normalized, "86"):
		normalized = normalized[len("86"):]
	}

	if strings.ContainsAny(normalized, "+()") {
		return "", errors.New("phone input format is invalid")
	}
	if len(normalized) != 11 {
		return "", errors.New("phone number must contain 11 digits")
	}
	if normalized[0] != '1' {
		return "", errors.New("phone number must start with 1")
	}
	return normalized, nil
}

// ParseChinaIDCard 解析中国身份证号码中的结构化信息。
func ParseChinaIDCard(text string, now time.Time) (*ChinaIDCardInfo, error) {
	if err := ValidateChinaIDCard(text); err != nil {
		return nil, err
	}

	text = strings.ToUpper(text)
	birthDate, _ := time.Parse("20060102", text[6:14])
	age := now.Year() - birthDate.Year()
	birthdayThisYear := time.Date(now.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, now.Location())
	if now.Before(birthdayThisYear) {
		age--
	}

	gender := GenderFemale
	if (text[16]-'0')%2 == 1 {
		gender = GenderMale
	}

	return &ChinaIDCardInfo{
		Number:       text,
		ProvinceCode: text[:2],
		BirthDate:    birthDate,
		Age:          age,
		Gender:       gender,
	}, nil
}

func isEmailLocalPart(local string) bool {
	if local == "" {
		return false
	}
	synthetic := local + "@example.com"
	addr, err := mail.ParseAddress(synthetic)
	return err == nil && addr.Name == "" && addr.Address == synthetic
}

func isEmailDomain(domain string) bool {
	if strings.HasPrefix(domain, "[") && strings.HasSuffix(domain, "]") {
		literal := domain[1 : len(domain)-1]
		if strings.HasPrefix(literal, "IPv6:") {
			return isIPv6(strings.TrimPrefix(literal, "IPv6:"))
		}
		return isIP(literal)
	}
	return isHostname(domain)
}

func parseAbsoluteURL(raw string) (*url.URL, error) {
	if raw == "" || strings.TrimSpace(raw) != raw {
		return nil, errors.New("URL is empty or contains surrounding whitespace")
	}
	u, err := url.ParseRequestURI(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		if err != nil {
			return nil, fmt.Errorf("invalid URL: %w", err)
		}
		return nil, errors.New("URL must be absolute and contain scheme and host")
	}
	return u, nil
}

func validateURLHost(u *url.URL) error {
	host := u.Hostname()
	if host == "" {
		return errors.New("URL host is empty")
	}
	if ValidateHostname(host) == nil || ValidateIP(host) == nil {
		return nil
	}
	return errors.New("URL host is invalid")
}

func isHostname(host string) bool {
	_, err := normalizeHostnameASCII(host)
	return err == nil
}

func normalizeHostnameASCII(host string) (string, error) {
	if host == "" {
		return "", errors.New("hostname is empty")
	}
	host = strings.TrimSuffix(host, ".")
	if host == "" {
		return "", errors.New("hostname is empty after trimming trailing dot")
	}
	if _, err := netip.ParseAddr(host); err == nil {
		return "", errors.New("hostname cannot be a raw IP address")
	}
	asciiHost, err := idna.Lookup.ToASCII(host)
	if err != nil {
		return "", fmt.Errorf("hostname IDNA conversion failed: %w", err)
	}
	if asciiHost == "" || len(asciiHost) > 253 {
		return "", errors.New("hostname length is invalid")
	}
	labels := strings.Split(asciiHost, ".")
	for _, label := range labels {
		if len(label) == 0 || len(label) > 63 {
			return "", errors.New("hostname label length is invalid")
		}
		if label[0] == '-' || label[len(label)-1] == '-' {
			return "", errors.New("hostname label cannot start or end with hyphen")
		}
		for _, b := range []byte(label) {
			if !isASCIIAlphaNum(b) && b != '-' {
				return "", errors.New("hostname contains invalid character")
			}
		}
	}
	return asciiHost, nil
}

func isASCIIAlphaNum(b byte) bool {
	return 'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9'
}

func isHexRune(b byte) bool {
	return '0' <= b && b <= '9' || 'a' <= b && b <= 'f' || 'A' <= b && b <= 'F'
}

func isIP(text string) bool {
	_, err := netip.ParseAddr(text)
	return err == nil
}

func isIPv6(text string) bool {
	addr, err := netip.ParseAddr(text)
	return err == nil && addr.Is6()
}

var chinaProvinceCodes = map[string]struct{}{
	"11": {}, "12": {}, "13": {}, "14": {}, "15": {},
	"21": {}, "22": {}, "23": {},
	"31": {}, "32": {}, "33": {}, "34": {}, "35": {}, "36": {}, "37": {},
	"41": {}, "42": {}, "43": {}, "44": {}, "45": {}, "46": {},
	"50": {}, "51": {}, "52": {}, "53": {}, "54": {},
	"61": {}, "62": {}, "63": {}, "64": {}, "65": {},
	"71": {}, "81": {}, "82": {}, "91": {},
}

var chinaIDCardWeights = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var chinaIDCardChecksumCodes = [...]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

const chinaUSCCCharset = "0123456789ABCDEFGHJKLMNPQRTUWXY"

var chinaUSCCValueMap = func() map[byte]int {
	m := make(map[byte]int, len(chinaUSCCCharset))
	for i := 0; i < len(chinaUSCCCharset); i++ {
		m[chinaUSCCCharset[i]] = i
	}
	return m
}()

var chinaUSCCWeights = [...]int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}

var chinaMobilePrefixes = func() map[string]struct{} {
	prefixes := make(map[string]struct{}, 55)
	addRange := func(start, end int) {
		for prefix := start; prefix <= end; prefix++ {
			prefixes[strconv.Itoa(prefix)] = struct{}{}
		}
	}

	addRange(130, 139)
	addRange(145, 149)
	addRange(150, 159)
	for _, prefix := range []int{162, 165, 166, 167} {
		prefixes[strconv.Itoa(prefix)] = struct{}{}
	}
	addRange(170, 179)
	addRange(180, 189)
	addRange(190, 199)
	return prefixes
}()
