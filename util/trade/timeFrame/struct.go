package gptimeframe

import "time"

type TimeFrame struct{}

func NewTimeFrame(layout string) *TimeFrame {

	return &TimeFrame{}
}

func DefaultTimeFrame() *TimeFrame {

	return &TimeFrame{}
}

func (t *TimeFrame) FromDuration(timeFrame string) (time.Duration, error) {

	return FromDuration(timeFrame)
}
