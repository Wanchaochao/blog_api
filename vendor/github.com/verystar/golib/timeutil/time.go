package timeutil

import "time"

func Date() string {
	return time.Now().Format("2006-01-02")
}

func DateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ParseDate(tstr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation("2006-01-02", tstr, loc)
}

func ParseDateTime(tstr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation("2006-01-02 15:04:05", tstr, loc)
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ToUnix(tstr string) int64 {
	t, err := ParseDateTime(tstr)

	if err != nil {
		return 0
	}
	return t.Unix()
}
