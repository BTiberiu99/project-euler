package utils

import "time"

func Timeit(f func()) time.Duration {
	now := time.Now()
	f()
	return time.Since(now)
}
