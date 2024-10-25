package timeutil

import "time"

var (
	layout = time.RFC3339
)

func ParseTime(s string) (time.Time, error) {
	return time.Parse(layout, s)
}
