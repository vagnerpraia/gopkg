package gpclock

import (
	"time"
)

type Clock struct {
	location string
	layout   string
}

func NewClock(location string, layout string) *Clock {

	return &Clock{
		location: location,
		layout:   layout,
	}
}

func NewClockAmericaSaoPaulo() *Clock {

	location := "America/Sao_Paulo"
	layout := "2006-01-02T15:04:05Z"

	return &Clock{
		location: location,
		layout:   layout,
	}
}

func (clock *Clock) Now() (*time.Time, error) {

	return Now(clock.location, clock.layout)
}
