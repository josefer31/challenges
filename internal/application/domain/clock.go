package domain

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type ClockImpl struct{}

func (receiver ClockImpl) Now() time.Time {

	return time.Now().UTC()
}

func NewClock() Clock { return ClockImpl{} }
