package gptimeframe

import "time"

type TimeFrame struct {
	layout string
}

func NewTimeFrame(layout string) *TimeFrame {

	return &TimeFrame{
		layout: layout,
	}
}

func DefaultTimeFrame() *TimeFrame {

	return &TimeFrame{
		layout: `^(\d+)([smhd])$`,
	}
}

func (t *TimeFrame) FromDuration(timeFrame string) (time.Duration, error) {

	return FromDuration(timeFrame, t.layout)
}
