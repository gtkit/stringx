package unsafex

import "unsafe"

// String2Bytes 在不发生额外内存拷贝的情况下，将字符串转换为字节切片。
//
// 返回的切片与原字符串共享底层数据，只适合只读场景，调用方不得修改返回切片。
func String2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// Bytes2String 在不发生额外内存拷贝的情况下，将字节切片转换为字符串。
//
// 返回的字符串与原字节切片共享底层数据；原切片后续修改会影响字符串视图。
func Bytes2String(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
