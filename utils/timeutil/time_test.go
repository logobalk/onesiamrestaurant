package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnNowAtSpecificTime(t *testing.T) {
	var (
		SpecificTime = time.Unix(123456789, 0)
	)

	FreezeAt(t, SpecificTime)

	now := Now()
	assert.Equal(t, SpecificTime, now)
}
