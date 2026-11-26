package stringx

import "github.com/gtkit/stringx/randomx"

// RNG 提供一组轻量随机数方法。
//
// 当前实现直接复用 math/rand/v2 的并发安全顶层函数，以减少共享状态带来的风险。
type RNG struct{}

// RandInt 返回区间 [min, max) 内的随机整数。
//
// 当 max 小于等于 min 时，返回 min。
func RandInt(low, high int) int { return randomx.RandInt(low, high) }

// RUint 返回区间 [0, n) 内的随机整数。
//
// 当 n 小于等于 0 时，返回 0。
func RUint(n int) int { return randomx.RUint(n) }

// RandIntHandler 连续生成 count 个随机整数，并将结果回调给 handler。
//
// 生成范围为 [0, maxN)。为了兼容旧行为，回调中的索引 i 按倒序传入。
func RandIntHandler(maxN, count int, handler func(num, i int)) {
	randomx.RandIntHandler(maxN, count, handler)
}

// Uint32 返回一个随机 uint32。
func (rng RNG) Uint32() uint32 { return randomx.RNG(rng).Uint32() }

// Uint32n 返回区间 [0, maxN) 内的随机 uint32。
func (rng RNG) Uint32n(maxN uint32) uint32 { return randomx.RNG(rng).Uint32n(maxN) }

// Uint64 返回一个随机 uint64。
func (rng RNG) Uint64() uint64 { return randomx.RNG(rng).Uint64() }

// Uint64n 返回区间 [0, maxN] 内的随机 uint64。
func (rng RNG) Uint64n(maxN uint64) uint64 { return randomx.RNG(rng).Uint64n(maxN) }
