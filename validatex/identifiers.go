package validatex

import "errors"

// IsUUID 判断字符串是否为合法 UUID。
func IsUUID(text string) bool {
	return ValidateUUID(text) == nil
}

// ValidateUUID 验证字符串是否为合法 UUID。
func ValidateUUID(text string) error {
	if len(text) != 36 {
		return errors.New("UUID length must be 36")
	}
	for i := range len(text) {
		b := text[i]
		switch i {
		case 8, 13, 18, 23:
			if b != '-' {
				return errors.New("UUID hyphen position is invalid")
			}
		default:
			if !isHexByte(b) {
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

func isHexByte(b byte) bool {
	return '0' <= b && b <= '9' || 'a' <= b && b <= 'f' || 'A' <= b && b <= 'F'
}
