package ding

import (
	"sync"
	"time"
)

// 有限的发送

//var send_map = make(map[string]time.Time)
var send_map sync.Map

func LimitedAlarm(key string, duration time.Duration, content string, at ...string) error {
	current := time.Now()
	if data, ok := send_map.Load(key); ok {
		// 如果没有超过时间限制 则不发送
		t := data.(time.Time)
		if t.Add(duration).Unix() > current.Unix() {
			return nil
		}
	}

	err := Alarm(content, at...)
	if err == nil {
		// 发送成功
		send_map.Store(key, current)
	}
	return err
}
