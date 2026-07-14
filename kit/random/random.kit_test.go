package randompkg

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrings(t *testing.T) {
	t.Run("长度正确", func(t *testing.T) {
		for _, size := range []int{0, 1, 3, 10, 100} {
			result := Strings(size)
			assert.Equal(t, size, len(result), "Strings(%d) 长度不正确", size)
		}
	})

	t.Run("字符范围正确", func(t *testing.T) {
		result := Strings(1000)
		for _, c := range result {
			isValid := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
			assert.True(t, isValid, "Strings 包含非法字符: %c", c)
		}
	})

	t.Run("两次生成不同", func(t *testing.T) {
		a := Strings(20)
		b := Strings(20)
		assert.NotEqual(t, a, b, "两次生成的随机字符串不应相同")
	})
}

func TestRandomStringBoundary(t *testing.T) {
	assert.Empty(t, Strings(-1))
	assert.Empty(t, Letter(-1))
	assert.Empty(t, Numeric(-1))
	assert.Empty(t, String(8, ""))
	assert.Empty(t, String(-1, CharsetAlphabet))
}

func TestLetter(t *testing.T) {
	t.Run("长度正确", func(t *testing.T) {
		result := Letter(10)
		assert.Equal(t, 10, len(result))
	})

	t.Run("只包含字母", func(t *testing.T) {
		result := Letter(1000)
		for _, c := range result {
			assert.True(t, unicode.IsLetter(c), "Letter 包含非字母字符: %c", c)
		}
	})
}

func TestNumeric(t *testing.T) {
	t.Run("长度正确", func(t *testing.T) {
		result := Numeric(6)
		assert.Equal(t, 6, len(result))
	})

	t.Run("只包含数字", func(t *testing.T) {
		result := Numeric(1000)
		for _, c := range result {
			assert.True(t, unicode.IsDigit(c), "Numeric 包含非数字字符: %c", c)
		}
	})
}

func TestNumericBetween(t *testing.T) {
	t.Run("正常范围", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			result := NumericBetween(1, 10)
			assert.GreaterOrEqual(t, result, int64(1))
			assert.Less(t, result, int64(10))
		}
	})

	t.Run("min大于max自动交换", func(t *testing.T) {
		result := NumericBetween(10, 1)
		assert.GreaterOrEqual(t, result, int64(1))
		assert.Less(t, result, int64(10))
	})

	t.Run("min等于max", func(t *testing.T) {
		result := NumericBetween(5, 5)
		assert.Equal(t, int64(5), result)
	})
}

func TestInt32Between(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := Int32Between(1, 100)
		assert.GreaterOrEqual(t, result, int32(1))
		assert.Less(t, result, int32(100))
	}
}

func TestIntBetween(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := IntBetween(10, 20)
		assert.GreaterOrEqual(t, result, 10)
		assert.Less(t, result, 20)
	}
}

func TestName(t *testing.T) {
	name := Name()
	assert.NotEmpty(t, name)
	assert.Contains(t, name, "_", "Name 应包含下划线分隔符")
}

func TestAlphabetLower(t *testing.T) {
	result := AlphabetLower(20)
	assert.Equal(t, 20, len(result))
	for _, c := range result {
		assert.True(t, c >= 'a' && c <= 'z', "AlphabetLower 包含非小写字母: %c", c)
	}
}

func TestHex(t *testing.T) {
	result := Hex(32)
	assert.Equal(t, 32, len(result))
	for _, c := range result {
		isHex := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')
		assert.True(t, isHex, "Hex 包含非十六进制字符: %c", c)
	}
}

func TestString(t *testing.T) {
	charset := "ABC"
	result := String(100, charset)
	assert.Equal(t, 100, len(result))
	for _, c := range result {
		assert.Contains(t, charset, string(c))
	}
}

func TestVerifyCode(t *testing.T) {
	code := VerifyCode(6)
	assert.Equal(t, 6, len(code))
	for _, c := range code {
		assert.True(t, unicode.IsDigit(c))
	}
}

func TestPassword(t *testing.T) {
	t.Run("最小长度为8", func(t *testing.T) {
		pwd := Password(4)
		assert.GreaterOrEqual(t, len(pwd), 8)
	})

	t.Run("包含大写小写数字特殊字符", func(t *testing.T) {
		pwd := Password(20)
		assert.Equal(t, 20, len(pwd))
		var hasUpper, hasLower, hasDigit, hasSpecial bool
		for _, c := range pwd {
			switch {
			case c >= 'A' && c <= 'Z':
				hasUpper = true
			case c >= 'a' && c <= 'z':
				hasLower = true
			case c >= '0' && c <= '9':
				hasDigit = true
			default:
				hasSpecial = true
			}
		}
		assert.True(t, hasUpper, "密码应包含大写字母")
		assert.True(t, hasLower, "密码应包含小写字母")
		assert.True(t, hasDigit, "密码应包含数字")
		assert.True(t, hasSpecial, "密码应包含特殊字符")
	})
}

func TestToken(t *testing.T) {
	token := Token(32)
	assert.Equal(t, 32, len(token))
	for _, c := range token {
		isAlphanumeric := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
		assert.True(t, isAlphanumeric, "Token 包含非字母数字字符: %c", c)
	}
}

func TestSecureRandom(t *testing.T) {
	t.Run("SecureBytes", func(t *testing.T) {
		got, err := SecureBytes(16)
		require.NoError(t, err)
		assert.Len(t, got, 16)

		empty, err := SecureBytes(0)
		require.NoError(t, err)
		assert.Empty(t, empty)
	})

	t.Run("SecureString", func(t *testing.T) {
		got, err := SecureString(32, "ABC")
		require.NoError(t, err)
		assert.Len(t, got, 32)
		for _, c := range got {
			assert.Contains(t, "ABC", string(c))
		}

		empty, err := SecureString(0, "")
		require.NoError(t, err)
		assert.Empty(t, empty)

		_, err = SecureString(8, "")
		require.Error(t, err)
	})

	t.Run("SecureToken", func(t *testing.T) {
		got, err := SecureToken(32)
		require.NoError(t, err)
		assert.Len(t, got, 32)
		for _, c := range got {
			isAlphanumeric := (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
			assert.True(t, isAlphanumeric)
		}
	})

	t.Run("SecureHex", func(t *testing.T) {
		got, err := SecureHex(32)
		require.NoError(t, err)
		assert.Len(t, got, 32)
		for _, c := range got {
			isHex := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')
			assert.True(t, isHex)
		}
	})

	t.Run("SecureBase64URL", func(t *testing.T) {
		got, err := SecureBase64URL(16)
		require.NoError(t, err)
		assert.NotEmpty(t, got)
		assert.NotContains(t, got, "+")
		assert.NotContains(t, got, "/")
		assert.NotContains(t, got, "=")
	})
}

func TestOrderNo(t *testing.T) {
	t.Run("最小后缀长度为4", func(t *testing.T) {
		no := OrderNo(2)
		assert.GreaterOrEqual(t, len(no), 14+4) // 时间戳14位 + 最少4位随机数
	})

	t.Run("正常长度", func(t *testing.T) {
		no := OrderNo(6)
		assert.Equal(t, 14+6, len(no))
	})
}

func TestTraceID(t *testing.T) {
	id := TraceID()
	assert.Equal(t, 32, len(id))
}

func TestBool(t *testing.T) {
	trueCount := 0
	total := 1000
	for i := 0; i < total; i++ {
		if Bool() {
			trueCount++
		}
	}
	// 概率应大致在 30%-70% 之间
	assert.Greater(t, trueCount, total/5)
	assert.Less(t, trueCount, total*4/5)
}

func TestElement(t *testing.T) {
	t.Run("正常切片", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		elem := Element(slice)
		assert.Contains(t, slice, elem)
	})

	t.Run("空切片返回零值", func(t *testing.T) {
		var slice []int
		elem := Element(slice)
		assert.Equal(t, 0, elem)
	})

	t.Run("单元素切片", func(t *testing.T) {
		slice := []int{42}
		assert.Equal(t, 42, Element(slice))
	})
}

func TestShuffle(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffled := make([]int, len(original))
	copy(shuffled, original)
	Shuffle(shuffled)

	// 元素应相同（排序后）
	assert.ElementsMatch(t, original, shuffled)
}

func TestSample(t *testing.T) {
	t.Run("取子集", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		result := Sample(slice, 3)
		assert.Equal(t, 3, len(result))
		// 每个元素都应在原切片中
		for _, v := range result {
			assert.Contains(t, slice, v)
		}
	})

	t.Run("n大于等于切片长度", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Sample(slice, 10)
		assert.Equal(t, 3, len(result))
		assert.ElementsMatch(t, slice, result)
	})

	t.Run("不修改原切片", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		original := make([]int, len(slice))
		copy(original, slice)
		Sample(slice, 3)
		assert.Equal(t, original, slice)
	})

	t.Run("n小于等于0返回空切片", func(t *testing.T) {
		assert.Empty(t, Sample([]int{1, 2, 3}, 0))
		assert.Empty(t, Sample([]int{1, 2, 3}, -1))
	})
}

func TestWeightedIndex(t *testing.T) {
	t.Run("空权重返回-1", func(t *testing.T) {
		assert.Equal(t, -1, WeightedIndex(nil))
		assert.Equal(t, -1, WeightedIndex([]int{}))
	})

	t.Run("单一权重", func(t *testing.T) {
		assert.Equal(t, 0, WeightedIndex([]int{1}))
	})

	t.Run("权重分布合理", func(t *testing.T) {
		weights := []int{90, 10}
		counts := make([]int, 2)
		total := 10000
		for i := 0; i < total; i++ {
			idx := WeightedIndex(weights)
			require.GreaterOrEqual(t, idx, 0)
			require.Less(t, idx, 2)
			counts[idx]++
		}
		// 第一个权重 90%，应占大多数
		assert.Greater(t, counts[0], total*70/100)
	})

	t.Run("全零权重不panic", func(t *testing.T) {
		idx := WeightedIndex([]int{0, 0, 0})
		assert.GreaterOrEqual(t, idx, 0)
		assert.Less(t, idx, 3)
	})
}
