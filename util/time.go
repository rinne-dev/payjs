package util

import "time"

// ExpiredAfter 生成过期时间
func ExpiredAfter(d time.Duration) string {
	return time.Now().Add(d).Format("20060102150405")
}
