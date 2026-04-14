package timepkg

import (
	"strconv"
	"time"
)

// Date Format
const (
	RFC3339           = time.RFC3339               // 例子：2022-08-03T16:40:58+08:00
	YmdHmsTZ          = "2006-01-02T15:04:05-0700" // 例子：2022-08-03T16:40:58+0800
	YmdHms            = "2006-01-02 15:04:05"
	YmdHm             = "2006-01-02 15:04"
	YmdH              = "2006-01-02 15"
	Ymd               = "2006-01-02"
	Ym                = "2006-01"
	Y                 = "2006"
	YSecond           = "20060102150405"
	YMinute           = "200601021504"
	YHour             = "2006010215"
	YDay              = "20060102"
	YMonth            = "200601"
	MDay              = "0102"
	YmdHmsMillisecond = "2006-01-02 15:04:05.999"
	YmdHmsMicrosecond = "2006-01-02 15:04:05.999999"
	YmdHmsNanosecond  = "2006-01-02 15:04:05.999999999"
	YMillisecond      = "20060102150405.999"
	YMicrosecond      = "20060102150405.999999"
	YNanosecond       = "20060102150405.999999999"
	YmdHmsMLogger     = "2006-01-02T15:04:05.999"

	YmdHmsChinese = "2006年01月02日 15:04:05"
	YmdChinese    = "2006年01月02日"
	HmsChinese    = "15时04分05秒"
)

// now time.Now()
func now() time.Time {
	return time.Now()
}

// Today 今天0时
func Today() time.Time {
	t := time.Now()

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// ToDay 2019-08-21 22:07:07 -> 2019-08-21 00:00:00
func ToDay(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// ToHour 2019-08-21 22:07:07 -> 2019-08-21 22:00:00
func ToHour(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// ToMinute 2019-08-21 22:07:07 -> 2019-08-21 22:07:00
func ToMinute(t time.Time) time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// ThisMonth 2019-08-21 22:07:07 -> 2019-08-01 00:00:00
func ThisMonth(t time.Time) time.Time {
	y, m, _ := t.Date()

	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// ThisYear 2019-08-21 22:07:07 -> 2019-01-01 00:00:00
func ThisYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// TimestampToTime timestamp to Time
func TimestampToTime(u int64) time.Time {
	return time.Unix(u, 0)
}

// TimestampToDate timestamp to date
func TimestampToDate(u int64, format string) string {
	return time.Unix(u, 0).Format(format)
}

// DateToTime date to time
func DateToTime(format, date string) (time.Time, error) {
	return time.ParseInLocation(format, date, time.Local)
}

func Time9999() time.Time {
	return time.Date(9999, 12, 31, 23, 59, 59, 0, time.Local)
}

// FormatRFC3339 to RFC3339
func FormatRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ==================== 常用应用场景 ====================

// IsToday 判断是否是今天
func IsToday(t time.Time) bool {
	n := time.Now()
	return t.Year() == n.Year() && t.Month() == n.Month() && t.Day() == n.Day()
}

// IsYesterday 判断是否是昨天
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	return t.Year() == yesterday.Year() && t.Month() == yesterday.Month() && t.Day() == yesterday.Day()
}

// IsThisWeek 判断是否在本周内（周一为一周起始）
func IsThisWeek(t time.Time) bool {
	start := WeekStart(time.Now())
	end := start.AddDate(0, 0, 7)
	return !t.Before(start) && t.Before(end)
}

// IsThisMonth 判断是否在本月内
func IsThisMonth(t time.Time) bool {
	n := time.Now()
	return t.Year() == n.Year() && t.Month() == n.Month()
}

// IsThisYear 判断是否在本年内
func IsThisYear(t time.Time) bool {
	return t.Year() == time.Now().Year()
}

// WeekStart 获取指定时间所在周的周一 00:00:00
func WeekStart(t time.Time) time.Time {
	weekday := t.Weekday()
	if weekday == time.Sunday {
		weekday = 7
	}
	d := time.Duration(1-weekday) * 24 * time.Hour
	return ToDay(t.Add(d))
}

// WeekEnd 获取指定时间所在周的周日 23:59:59
func WeekEnd(t time.Time) time.Time {
	start := WeekStart(t)
	return start.AddDate(0, 0, 7).Add(-time.Second)
}

// MonthEnd 获取指定时间所在月的最后一天 23:59:59
func MonthEnd(t time.Time) time.Time {
	return ThisMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

// DaysBetween 计算两个时间之间的天数差（绝对值）
func DaysBetween(a, b time.Time) int {
	a = ToDay(a)
	b = ToDay(b)
	diff := a.Sub(b)
	days := int(diff.Hours() / 24)
	if days < 0 {
		return -days
	}
	return days
}

// StartOfDay 获取指定时间当天的 00:00:00（等价于 ToDay）
func StartOfDay(t time.Time) time.Time {
	return ToDay(t)
}

// EndOfDay 获取指定时间当天的 23:59:59
func EndOfDay(t time.Time) time.Time {
	return ToDay(t).Add(24*time.Hour - time.Second)
}

// AddDays 增加天数（支持负数）
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddMonths 增加月数（支持负数）
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// IsBefore 判断 a 是否在 b 之前
func IsBefore(a, b time.Time) bool {
	return a.Before(b)
}

// IsAfter 判断 a 是否在 b 之后
func IsAfter(a, b time.Time) bool {
	return a.After(b)
}

// IsBetween 判断 t 是否在 start 和 end 之间（包含边界）
func IsBetween(t, start, end time.Time) bool {
	return !t.Before(start) && !t.After(end)
}

// IsExpired 判断指定时间是否已过期（在当前时间之前）
func IsExpired(t time.Time) bool {
	return t.Before(time.Now())
}

// IsZero 判断时间是否为零值
func IsZero(t time.Time) bool {
	return t.IsZero()
}

// FriendlyDuration 将 Duration 转为友好的中文描述
// 例如：1h30m → "1小时30分钟"，72h → "3天"
func FriendlyDuration(d time.Duration) string {
	if d < time.Minute {
		return formatInt(int(d.Seconds())) + "秒"
	}
	if d < time.Hour {
		m := int(d.Minutes())
		s := int(d.Seconds()) % 60
		if s == 0 {
			return formatInt(m) + "分钟"
		}
		return formatInt(m) + "分钟" + formatInt(s) + "秒"
	}
	if d < 24*time.Hour {
		h := int(d.Hours())
		m := int(d.Minutes()) % 60
		if m == 0 {
			return formatInt(h) + "小时"
		}
		return formatInt(h) + "小时" + formatInt(m) + "分钟"
	}
	days := int(d.Hours()) / 24
	h := int(d.Hours()) % 24
	if h == 0 {
		return formatInt(days) + "天"
	}
	return formatInt(days) + "天" + formatInt(h) + "小时"
}

// TimeAgo 将时间转为"多久前"的友好描述
// 例如：刚刚、5分钟前、3小时前、昨天、3天前、2个月前、1年前
func TimeAgo(t time.Time) string {
	d := time.Since(t)
	switch {
	case d < time.Minute:
		return "刚刚"
	case d < time.Hour:
		return formatInt(int(d.Minutes())) + "分钟前"
	case d < 24*time.Hour:
		return formatInt(int(d.Hours())) + "小时前"
	case d < 48*time.Hour:
		return "昨天"
	case d < 30*24*time.Hour:
		return formatInt(int(d.Hours()/24)) + "天前"
	case d < 365*24*time.Hour:
		return formatInt(int(d.Hours()/24/30)) + "个月前"
	default:
		return formatInt(int(d.Hours()/24/365)) + "年前"
	}
}

// TimestampMillis 获取毫秒级时间戳
func TimestampMillis() int64 {
	return time.Now().UnixMilli()
}

// TimestampMicros 获取微秒级时间戳
func TimestampMicros() int64 {
	return time.Now().UnixMicro()
}

// MillisToTime 毫秒时间戳转 time.Time
func MillisToTime(ms int64) time.Time {
	return time.UnixMilli(ms)
}

// DaysInMonth 获取指定年月的天数
func DaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// IsLeapYear 判断是否为闰年
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// IsWeekend 判断是否为周末（周六或周日）
func IsWeekend(t time.Time) bool {
	day := t.Weekday()
	return day == time.Saturday || day == time.Sunday
}

// IsWorkday 判断是否为工作日（周一至周五）
func IsWorkday(t time.Time) bool {
	return !IsWeekend(t)
}

// formatInt int 转字符串（内部辅助函数）
func formatInt(n int) string {
	return strconv.Itoa(n)
}
