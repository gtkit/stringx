package stringx

import (
	"strconv"
	"time"
)

// MillTime 返回当前 Unix 毫秒时间戳对应的十进制字符串。
func MillTime() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

// UnixTime 返回当前 Unix 秒级时间戳对应的十进制字符串。
func UnixTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
