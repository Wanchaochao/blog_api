package convert

import "testing"

func TestStrTo_Time(t *testing.T) {
	t1 := "2012-10-24 07:49:00"
	ct1 := StrTo(t1).MustTime().Format("2006-01-02 15:04:05")
	if t1 != ct1 {
		t.Error("StrTo_Time fail")
	}

	t2 := "2012-10-24T07:49:00+08:00"

	ct2, err := StrTo(t2).Time()

	if err != nil {
		t.Error(err)
	}

	ct3 := ct2.Format("2006-01-02 15:04:05")
	if ct3 != t1 {
		t.Error("StrTo_Time RFC3339 fail")
	}
}
