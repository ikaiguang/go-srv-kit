package timepkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// go test -v -count 1 ./kit/time -run TestFormat
func TestFormat(t *testing.T) {
	tNow := now()
	assert.NotEmpty(t, FormatRFC3339(tNow))
	assert.Contains(t, FormatRFC3339(tNow), "T")
}

func TestToday(t *testing.T) {
	today := Today()
	n := time.Now()
	assert.Equal(t, n.Year(), today.Year())
	assert.Equal(t, n.Month(), today.Month())
	assert.Equal(t, n.Day(), today.Day())
	assert.Equal(t, 0, today.Hour())
	assert.Equal(t, 0, today.Minute())
	assert.Equal(t, 0, today.Second())
}

func TestToDay(t *testing.T) {
	input := time.Date(2024, 8, 15, 14, 30, 45, 0, time.Local)
	result := ToDay(input)
	assert.Equal(t, 2024, result.Year())
	assert.Equal(t, time.August, result.Month())
	assert.Equal(t, 15, result.Day())
	assert.Equal(t, 0, result.Hour())
	assert.Equal(t, 0, result.Minute())
	assert.Equal(t, 0, result.Second())
}

func TestToHour(t *testing.T) {
	input := time.Date(2024, 8, 15, 14, 30, 45, 0, time.Local)
	result := ToHour(input)
	assert.Equal(t, 14, result.Hour())
	assert.Equal(t, 0, result.Minute())
	assert.Equal(t, 0, result.Second())
}

func TestToMinute(t *testing.T) {
	input := time.Date(2024, 8, 15, 14, 30, 45, 0, time.Local)
	result := ToMinute(input)
	assert.Equal(t, 14, result.Hour())
	assert.Equal(t, 30, result.Minute())
	assert.Equal(t, 0, result.Second())
}

func TestThisMonth(t *testing.T) {
	input := time.Date(2024, 8, 15, 14, 30, 45, 0, time.Local)
	result := ThisMonth(input)
	assert.Equal(t, 2024, result.Year())
	assert.Equal(t, time.August, result.Month())
	assert.Equal(t, 1, result.Day())
	assert.Equal(t, 0, result.Hour())
}

func TestThisYear(t *testing.T) {
	input := time.Date(2024, 8, 15, 14, 30, 45, 0, time.Local)
	result := ThisYear(input)
	assert.Equal(t, 2024, result.Year())
	assert.Equal(t, time.January, result.Month())
	assert.Equal(t, 1, result.Day())
}

func TestTimestampToTime(t *testing.T) {
	ts := int64(1692100000)
	result := TimestampToTime(ts)
	assert.Equal(t, ts, result.Unix())
}

func TestTimestampToDate(t *testing.T) {
	ts := int64(1692100000)
	result := TimestampToDate(ts, Ymd)
	assert.NotEmpty(t, result)
	assert.Regexp(t, `^\d{4}-\d{2}-\d{2}$`, result)
}

func TestDateToTime(t *testing.T) {
	t.Run("正常解析", func(t *testing.T) {
		result, err := DateToTime(YmdHms, "2024-08-15 14:30:45")
		assert.Nil(t, err)
		assert.Equal(t, 2024, result.Year())
		assert.Equal(t, time.August, result.Month())
		assert.Equal(t, 15, result.Day())
	})

	t.Run("格式不匹配", func(t *testing.T) {
		_, err := DateToTime(YmdHms, "invalid")
		assert.NotNil(t, err)
	})
}

func TestTime9999(t *testing.T) {
	result := Time9999()
	assert.Equal(t, 9999, result.Year())
	assert.Equal(t, time.December, result.Month())
	assert.Equal(t, 31, result.Day())
}

func TestIsToday(t *testing.T) {
	assert.True(t, IsToday(time.Now()))
	assert.False(t, IsToday(time.Now().AddDate(0, 0, -1)))
	assert.False(t, IsToday(time.Now().AddDate(0, 0, 1)))
}

func TestIsYesterday(t *testing.T) {
	assert.True(t, IsYesterday(time.Now().AddDate(0, 0, -1)))
	assert.False(t, IsYesterday(time.Now()))
}

func TestIsThisMonth(t *testing.T) {
	assert.True(t, IsThisMonth(time.Now()))
	assert.False(t, IsThisMonth(time.Now().AddDate(0, -2, 0)))
}

func TestIsThisYear(t *testing.T) {
	assert.True(t, IsThisYear(time.Now()))
	assert.False(t, IsThisYear(time.Now().AddDate(-1, 0, 0)))
}

func TestWeekStart(t *testing.T) {
	// 2024-08-14 是周三
	wed := time.Date(2024, 8, 14, 10, 30, 0, 0, time.Local)
	start := WeekStart(wed)
	assert.Equal(t, time.Monday, start.Weekday())
	assert.Equal(t, 12, start.Day()) // 2024-08-12 是周一
	assert.Equal(t, 0, start.Hour())
}

func TestWeekEnd(t *testing.T) {
	wed := time.Date(2024, 8, 14, 10, 30, 0, 0, time.Local)
	end := WeekEnd(wed)
	assert.Equal(t, time.Sunday, end.Weekday())
	assert.Equal(t, 23, end.Hour())
	assert.Equal(t, 59, end.Minute())
	assert.Equal(t, 59, end.Second())
}

func TestMonthEnd(t *testing.T) {
	t.Run("31天的月份", func(t *testing.T) {
		input := time.Date(2024, 8, 15, 0, 0, 0, 0, time.Local)
		end := MonthEnd(input)
		assert.Equal(t, 31, end.Day())
		assert.Equal(t, 23, end.Hour())
	})

	t.Run("闰年2月", func(t *testing.T) {
		input := time.Date(2024, 2, 10, 0, 0, 0, 0, time.Local)
		end := MonthEnd(input)
		assert.Equal(t, 29, end.Day())
	})

	t.Run("非闰年2月", func(t *testing.T) {
		input := time.Date(2023, 2, 10, 0, 0, 0, 0, time.Local)
		end := MonthEnd(input)
		assert.Equal(t, 28, end.Day())
	})
}

func TestDaysBetween(t *testing.T) {
	a := time.Date(2024, 8, 15, 0, 0, 0, 0, time.Local)
	b := time.Date(2024, 8, 10, 0, 0, 0, 0, time.Local)
	assert.Equal(t, 5, DaysBetween(a, b))
	assert.Equal(t, 5, DaysBetween(b, a)) // 绝对值
	assert.Equal(t, 0, DaysBetween(a, a))
}

func TestStartOfDay_EndOfDay(t *testing.T) {
	input := time.Date(2024, 8, 15, 14, 30, 45, 0, time.Local)

	start := StartOfDay(input)
	assert.Equal(t, 0, start.Hour())
	assert.Equal(t, 0, start.Minute())

	end := EndOfDay(input)
	assert.Equal(t, 23, end.Hour())
	assert.Equal(t, 59, end.Minute())
	assert.Equal(t, 59, end.Second())
}

func TestAddDays(t *testing.T) {
	base := time.Date(2024, 8, 15, 0, 0, 0, 0, time.Local)
	assert.Equal(t, 20, AddDays(base, 5).Day())
	assert.Equal(t, 10, AddDays(base, -5).Day())
}

func TestAddMonths(t *testing.T) {
	base := time.Date(2024, 8, 15, 0, 0, 0, 0, time.Local)
	assert.Equal(t, time.November, AddMonths(base, 3).Month())
	assert.Equal(t, time.May, AddMonths(base, -3).Month())
}

func TestIsBefore_IsAfter(t *testing.T) {
	a := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	b := time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local)
	assert.True(t, IsBefore(a, b))
	assert.False(t, IsBefore(b, a))
	assert.True(t, IsAfter(b, a))
	assert.False(t, IsAfter(a, b))
}

func TestIsBetween(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local)
	mid := time.Date(2024, 6, 15, 0, 0, 0, 0, time.Local)
	before := time.Date(2023, 6, 15, 0, 0, 0, 0, time.Local)

	assert.True(t, IsBetween(mid, start, end))
	assert.True(t, IsBetween(start, start, end)) // 包含边界
	assert.True(t, IsBetween(end, start, end))   // 包含边界
	assert.False(t, IsBetween(before, start, end))
}

func TestIsExpired(t *testing.T) {
	assert.True(t, IsExpired(time.Now().Add(-time.Hour)))
	assert.False(t, IsExpired(time.Now().Add(time.Hour)))
}

func TestIsZero(t *testing.T) {
	assert.True(t, IsZero(time.Time{}))
	assert.False(t, IsZero(time.Now()))
}

func TestFriendlyDuration(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{"秒", 30 * time.Second, "30秒"},
		{"分钟", 5 * time.Minute, "5分钟"},
		{"分钟+秒", 5*time.Minute + 30*time.Second, "5分钟30秒"},
		{"小时", 2 * time.Hour, "2小时"},
		{"小时+分钟", 2*time.Hour + 30*time.Minute, "2小时30分钟"},
		{"天", 72 * time.Hour, "3天"},
		{"天+小时", 73 * time.Hour, "3天1小时"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, FriendlyDuration(tt.d))
		})
	}
}

func TestTimestampMillis(t *testing.T) {
	ms := TimestampMillis()
	assert.Greater(t, ms, int64(0))
}

func TestMillisToTime(t *testing.T) {
	ms := time.Now().UnixMilli()
	result := MillisToTime(ms)
	assert.Equal(t, ms, result.UnixMilli())
}

func TestDaysInMonth(t *testing.T) {
	assert.Equal(t, 31, DaysInMonth(2024, time.January))
	assert.Equal(t, 29, DaysInMonth(2024, time.February)) // 闰年
	assert.Equal(t, 28, DaysInMonth(2023, time.February)) // 非闰年
	assert.Equal(t, 30, DaysInMonth(2024, time.April))
}

func TestIsLeapYear(t *testing.T) {
	assert.True(t, IsLeapYear(2024))
	assert.True(t, IsLeapYear(2000))
	assert.False(t, IsLeapYear(2023))
	assert.False(t, IsLeapYear(1900))
}

func TestIsWeekend_IsWorkday(t *testing.T) {
	// 2024-08-17 是周六
	sat := time.Date(2024, 8, 17, 0, 0, 0, 0, time.Local)
	// 2024-08-18 是周日
	sun := time.Date(2024, 8, 18, 0, 0, 0, 0, time.Local)
	// 2024-08-19 是周一
	mon := time.Date(2024, 8, 19, 0, 0, 0, 0, time.Local)

	assert.True(t, IsWeekend(sat))
	assert.True(t, IsWeekend(sun))
	assert.False(t, IsWeekend(mon))

	assert.False(t, IsWorkday(sat))
	assert.False(t, IsWorkday(sun))
	assert.True(t, IsWorkday(mon))
}
