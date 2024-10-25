package timeutil

import (
	"time"
)

var (
	layout = time.RFC3339
)

func ParseTime(s string) (time.Time, error) {
	return time.Parse(layout, s)
}

func ParseStartTimeAndEndTime(startTimeStr, endTimeStr string) (time.Time, time.Time, error) {
	startTime := time.Time{}
	endTime := time.Time{}

	if startTimeStr != "" {
		t, err := ParseTime(startTimeStr)
		startTime = t
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
	}

	if endTimeStr != "" {
		t, err := ParseTime(endTimeStr)
		endTime = t
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
	}

	return startTime, endTime, nil
}
