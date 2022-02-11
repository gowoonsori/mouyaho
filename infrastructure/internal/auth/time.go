package auth

import "time"

func Unix100nano(timeNow time.Time) int64 {
	return timeNow.Unix()*1e7 + int64(timeNow.Nanosecond())/100
}

func TillNext100nano(lastTimestamp int64) int64 {
	timestamp := Unix100nano(time.Now())
	for timestamp <= lastTimestamp {
		timestamp = Unix100nano(time.Now())
	}
	return timestamp
}
