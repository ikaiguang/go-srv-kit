package randompkg

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	CharsetLowercase = "abcdefghijklmnopqrstuvwxyz"
	CharsetAlphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	CharsetNumeral   = "1234567890"
	CharsetHex       = "1234567890abcdef"
)

// NewRandHandler ...
func NewRandHandler() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Name random name
func Name() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36) + "_" + Strings(5)
}

// Strings : A-Z a-z 0-9
func Strings(size int) string {
	randHandler := NewRandHandler()
	res := make([]byte, size)
	for i := range res {
		switch randHandler.Intn(3) {
		case 0:
			res[i] = byte(randHandler.Intn(26) + 65) // A-Z
		case 1:
			res[i] = byte(randHandler.Intn(26) + 97) // a-z
		default:
			res[i] = byte(randHandler.Intn(10) + 48) // 0-9
		}
	}
	return string(res)
}

// Letter : A-Z a-z
func Letter(size int) string {
	randHandler := NewRandHandler()
	res := make([]byte, size)
	for i := range res {
		switch randHandler.Intn(2) {
		case 0:
			res[i] = byte(randHandler.Intn(26) + 65) // A-Z
		default:
			res[i] = byte(randHandler.Intn(26) + 97) // a-z
		}
	}
	return string(res)
}

// Numeric 0-9
func Numeric(size int) string {
	randHandler := NewRandHandler()
	res := make([]byte, size)
	for i := range res {
		res[i] = byte(randHandler.Intn(10) + 48) // 0-9
	}
	return string(res)
}

// AlphabetLower 从小写字符集生成指定长度的随机字符串
func AlphabetLower(n int) string {
	return String(n, CharsetLowercase)
}

func Hex(n int) string {
	return String(n, CharsetHex)
}

// String returns a random string n characters long, composed of entities from charset.
func String(n int, charset string) string {
	randomHandler := NewRandHandler()
	randStr := make([]byte, n)
	charLen := len(charset)
	for i := range randStr {
		randStr[i] = charset[randomHandler.Intn(charLen)]
	}
	return string(randStr)
}

// Int32Between random number between min-max
func Int32Between(min, max int32) int32 {
	if min == max {
		return min
	}

	if min >= max {
		min, max = max, min
	}

	return NewRandHandler().Int31n(max-min) + min
}

// Int64Between random number between min-max
func Int64Between(min, max int64) int64 {
	if min == max {
		return min
	}

	if min >= max {
		min, max = max, min
	}

	return NewRandHandler().Int63n(max-min) + min
}

// NumericBetween random number between min-max
func NumericBetween(min, max int64) int64 {
	return Int64Between(min, max)
}

// IntBetween random number between min-max
func IntBetween(min, max int) int {
	if min == max {
		return min
	}

	if min >= max {
		min, max = max, min
	}

	return NewRandHandler().Intn(max-min) + min
}

// ==================== 常用应用场景 ====================

// 扩展字符集常量
const (
	CharsetAlphanumeric = CharsetAlphabet + CharsetNumeral
	CharsetUppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetSpecial      = "!@#$%^&*()-_=+[]{}|;:,.<>?"
	CharsetPassword     = CharsetAlphanumeric + "!@#$%^&*"
)

// VerifyCode 生成纯数字验证码（短信/邮箱验证码）
// 常用长度：4 位或 6 位
func VerifyCode(length int) string {
	return Numeric(length)
}

// Password 生成随机密码（包含大小写字母、数字、特殊字符）
func Password(length int) string {
	if length < 8 {
		length = 8
	}
	randHandler := NewRandHandler()
	res := make([]byte, length)
	// 确保至少包含一个大写、一个小写、一个数字、一个特殊字符
	res[0] = CharsetUppercase[randHandler.Intn(len(CharsetUppercase))]
	res[1] = CharsetLowercase[randHandler.Intn(len(CharsetLowercase))]
	res[2] = CharsetNumeral[randHandler.Intn(len(CharsetNumeral))]
	res[3] = CharsetSpecial[randHandler.Intn(len(CharsetSpecial))]
	// 剩余位置随机填充
	for i := 4; i < length; i++ {
		res[i] = CharsetPassword[randHandler.Intn(len(CharsetPassword))]
	}
	// 打乱顺序
	randHandler.Shuffle(length, func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	return string(res)
}

// Token 生成 URL 安全的随机 token（用于 API token、重置密码链接等）
func Token(length int) string {
	return String(length, CharsetAlphanumeric)
}

// OrderNo 生成订单号：时间戳前缀 + 随机数字后缀
// 格式：20060102150405 + N 位随机数字，默认 6 位
func OrderNo(randomSuffixLen int) string {
	if randomSuffixLen < 4 {
		randomSuffixLen = 4
	}
	return time.Now().Format("20060102150405") + Numeric(randomSuffixLen)
}

// TraceID 生成 32 位十六进制 trace ID（兼容 OpenTelemetry 格式）
func TraceID() string {
	return Hex(32)
}

// Bool 随机返回 true 或 false
func Bool() bool {
	return NewRandHandler().Intn(2) == 0
}

// Element 从切片中随机选取一个元素
func Element[T any](slice []T) T {
	var zero T
	if len(slice) == 0 {
		return zero
	}
	return slice[NewRandHandler().Intn(len(slice))]
}

// Shuffle 随机打乱切片（原地修改）
func Shuffle[T any](slice []T) {
	NewRandHandler().Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// Sample 从切片中随机选取 n 个不重复元素
func Sample[T any](slice []T, n int) []T {
	length := len(slice)
	if n >= length {
		copied := make([]T, length)
		copy(copied, slice)
		Shuffle(copied)
		return copied
	}
	// Fisher-Yates 部分洗牌
	copied := make([]T, length)
	copy(copied, slice)
	randHandler := NewRandHandler()
	for i := 0; i < n; i++ {
		j := i + randHandler.Intn(length-i)
		copied[i], copied[j] = copied[j], copied[i]
	}
	return copied[:n]
}

// WeightedIndex 按权重随机选择索引
// weights 为各选项的权重值，返回被选中的索引
func WeightedIndex(weights []int) int {
	if len(weights) == 0 {
		return -1
	}
	total := 0
	for _, w := range weights {
		total += w
	}
	if total <= 0 {
		return NewRandHandler().Intn(len(weights))
	}
	r := NewRandHandler().Intn(total)
	for i, w := range weights {
		r -= w
		if r < 0 {
			return i
		}
	}
	return len(weights) - 1
}
