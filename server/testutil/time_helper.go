package testutil

import (
	"time"
)

// FakeTime is the mock of the time.
var FakeTime time.Time

// SetFakeTime sets FakeTime.
func SetFakeTime(t time.Time) {
	FakeTime = t
}

// ResetFakeTime resets FakeTime.
func ResetFakeTime() {
	FakeTime = time.Time{}
}

// TimeNow returns the current time.
func TimeNow() time.Time {
	if !FakeTime.IsZero() {
		return FakeTime
	}
	return time.Now()
}
