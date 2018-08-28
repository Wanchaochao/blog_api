package ding

import (
	"testing"
	"time"
)

func TestLimitedAlarm(t *testing.T) {
	DING_TALK_TOKEN = "xx"
	for i := 0 ; i < 10 ; i ++ {
		LimitedAlarm("test" , time.Second * 10 , "123123")
	}
}