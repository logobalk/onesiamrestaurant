package timeutil

import (
	"testing"
	"time"
)

var (
	timeFunc func() time.Time
)

func init() {
	timeFunc = stdTime
}

func stdTime() time.Time {
	return time.Now()
}

func Now() time.Time {
	return timeFunc()
}

func FreezeAt(t *testing.T, at time.Time) {
	timeFunc = func() time.Time {
		return at
	}
	t.Cleanup(func() {
		timeFunc = stdTime
	})
}
