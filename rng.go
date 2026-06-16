package stringx

import (
	"math"
	rand2 "math/rand/v2"
)

// RNG 提供一组轻量随机数方法。
//
// 当前实现直接复用 math/rand/v2 的并发安全顶层函数，以减少共享状态带来的风险。
type RNG struct{}

// RandInt 返回区间 [min, max) 内的随机整数。
//
// 当 max 小于等于 min 时，返回 min。
func RandInt(low, high int) int {
	if high <= low {
		return low
	}

	span := uint(high) - uint(low)                            //nolint:gosec // G115: intentional modular arithmetic for full-width int ranges.
	return int(uint(low) + uint(rand2.Uint64N(uint64(span)))) //nolint:gosec // G115: result is constrained to the int range by construction.
}

// RUint 返回区间 [0, n) 内的随机整数。
//
// 当 n 小于等于 0 时，返回 0。
func RUint(n int) int {
	if n <= 0 {
		return 0
	}

	return int(rand2.Uint64N(uint64(n))) //nolint:gosec // G115: n is positive int, so the generated value is always < n.
}

// RandIntHandler 连续生成 count 个随机整数，并将结果回调给 handler。
//
// 生成范围为 [0, maxN)。为了兼容旧行为，回调中的索引 i 按倒序传入。
func RandIntHandler(maxN, count int, handler func(num, i int)) {
	if maxN <= 0 || count <= 0 || handler == nil {
		return
	}

	for i := count - 1; i >= 0; i-- {
		handler(RandInt(0, maxN), i)
	}
}

// Uint32 返回一个随机 uint32。
func (RNG) Uint32() uint32 {
	return rand2.Uint32()
}

// Uint32n 返回区间 [0, maxN) 内的随机 uint32。
func (RNG) Uint32n(maxN uint32) uint32 {
	if maxN == 0 {
		return 0
	}

	return rand2.Uint32N(maxN)
}

// Uint64 返回一个随机 uint64。
func (RNG) Uint64() uint64 {
	return rand2.Uint64()
}

// Uint64n 返回区间 [0, maxN] 内的随机 uint64。
func (RNG) Uint64n(maxN uint64) uint64 {
	if maxN == math.MaxUint64 {
		return rand2.Uint64()
	}

	return rand2.Uint64N(maxN + 1)
}
