package business

import (
	"errors"
	"strings"
	"time"
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

// ValidateChinaIDCard 验证 18 位中国居民身份证。
func ValidateChinaIDCard(text string) error {
	_, _, err := validateChinaIDCard(text)
	return err
}

// IsChinaIDCard 判断字符串是否为合法中国身份证。
func IsChinaIDCard(text string) bool {
	return ValidateChinaIDCard(text) == nil
}

// ParseChinaIDCard 解析中国身份证号码中的结构化信息。
func ParseChinaIDCard(text string, now time.Time) (ChinaIDCardInfo, error) {
	number, birthDate, err := validateChinaIDCard(text)
	if err != nil {
		return ChinaIDCardInfo{}, err
	}

	age := now.Year() - birthDate.Year()
	birthdayThisYear := time.Date(now.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, now.Location())
	if now.Before(birthdayThisYear) {
		age--
	}

	gender := GenderFemale
	if (number[16]-'0')%2 == 1 {
		gender = GenderMale
	}

	return ChinaIDCardInfo{
		Number:       number,
		ProvinceCode: number[:2],
		BirthDate:    birthDate,
		Age:          age,
		Gender:       gender,
	}, nil
}

func validateChinaIDCard(text string) (string, time.Time, error) {
	if len(text) != 18 {
		return "", time.Time{}, errors.New("ID card length must be 18")
	}

	number := strings.ToUpper(text)
	for i := range 17 {
		if number[i] < '0' || number[i] > '9' {
			return "", time.Time{}, errors.New("ID card must contain digits in first 17 positions")
		}
	}
	if (number[17] < '0' || number[17] > '9') && number[17] != 'X' {
		return "", time.Time{}, errors.New("ID card last character must be digit or X")
	}
	if _, ok := chinaProvinceCodes[number[:2]]; !ok {
		return "", time.Time{}, errors.New("ID card province code is invalid")
	}

	birth := number[6:14]
	date, err := time.Parse("20060102", birth)
	if err != nil || date.Format("20060102") != birth {
		return "", time.Time{}, errors.New("ID card birth date is invalid")
	}
	if number[17] != chinaIDCardChecksum(number) {
		return "", time.Time{}, errors.New("ID card checksum is invalid")
	}
	return number, date, nil
}

func chinaIDCardChecksum(number string) byte {
	sum := 0
	for i, weight := range chinaIDCardWeights {
		sum += int(number[i]-'0') * weight
	}
	return chinaIDCardChecksumCodes[sum%11]
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

var (
	chinaIDCardWeights       = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	chinaIDCardChecksumCodes = [...]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)
