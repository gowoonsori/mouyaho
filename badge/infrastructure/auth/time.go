package auth

import "time"

func unix100nano(timeNow time.Time) int64 {
	return timeNow.Unix()*1e7 + int64(timeNow.Nanosecond())/100
}

func tillNext100nano(lastTimestamp int64) int64 {
	timestamp := unix100nano(time.Now())
	for timestamp <= lastTimestamp {
		timestamp = unix100nano(time.Now())
	}
	return timestamp
}
