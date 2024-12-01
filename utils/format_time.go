package utils

import "time"

func FormatTime() time.Time {
	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	reconstructedTime := time.Date(year, month, day, hour, minute, second, 0, time.Local)

	return reconstructedTime
}
