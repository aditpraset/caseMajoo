package helpers

import (
	"time"
)

func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}
