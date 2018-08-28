package collection

import (
	"testing"
	"time"
)

type t_users struct {
	UserId       int       `json:"user_id"`
	UserName     string    `json:"user_name" `
	Status       int       `json:"status" `
	Timezone     string    `json:"timezone" `
	Lang         string    `json:"lang" `
	CreateTime   time.Time `json:"create_time" `
	UpdateTime   time.Time `json:"update_time" `
}

var m = map[string]string{
	"user_id":        "4",
	"user_name":      "test",
	"status":         "1",
	"timezone":       "Asia/Shanghai",
	"lang":           "zh-CN",
	"create_time":    "2015-03-18 18:20:28",
	"update_time":    "2017-09-20 10:29:59",
}

func BenchmarkScanStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := &t_users{}
		ScanStruct(m, user)
	}
}

func TestScanStruct(t *testing.T) {
	user := &t_users{}
	err := ScanStruct(m, user)

	if err != nil {
		t.Error(err)
	}

	if user.UserId != 4 {
		t.Error("Parse user_id error:", user.UserId)
	}

	if user.CreateTime.Format("2006-01-02 15:04:05") != "2015-03-18 18:20:28" {
		t.Error("Parse create_time error:", user.CreateTime)
	}

}
