package util

import (
	"testing"
	"time"
)

func TestExpiredAfter(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	after, _ := time.ParseInLocation("20060102150405", ExpiredAfter(24*time.Hour), loc)
	if after.Sub(now)-(24*time.Hour) > time.Microsecond {
		t.Errorf("util.ExpiredAfter test failed\n")
	}
}
