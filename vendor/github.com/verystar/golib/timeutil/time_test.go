package timeutil

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	date := Date()
	tm, err := ParseDate(date)

	if err != nil || FormatDate(tm) != date {
		t.Fatalf("parse date error")
	}
}

func TestParseDateTime(t *testing.T) {
	date := DateTime()
	tm, err := ParseDateTime(date)

	if err != nil || FormatDateTime(tm) != date {
		t.Fatalf("parse datetime error")
	}
}

func TestFormatDate(t *testing.T) {
	testDate := "2012-10-24"
	tm, err := ParseDate(testDate)

	if err != nil || FormatDate(tm) != testDate {
		t.Fatalf("format date error")
	}
}

func TestFormatDateTime(t *testing.T) {
	testDate := "2012-10-24 07:21:15"
	tm, err := ParseDateTime(testDate)

	if err != nil || FormatDateTime(tm) != testDate {
		t.Fatalf("format date error")
	}
}

func TestToUnix(t *testing.T) {
	testDate := "2012-10-24 07:21:15"

	if ToUnix(testDate) != 1351034475 {
		t.Fatalf("to unix error")
	}
}
