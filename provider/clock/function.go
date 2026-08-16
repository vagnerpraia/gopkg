package gpclock

import (
	"time"
	_ "time/tzdata"
)

func Now(loc string, layout string) (*time.Time, error) {

	location, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}

	now, err := time.Parse(layout, time.Now().In(location).Format(layout))
	if err != nil {
		return nil, err
	}

	return &now, nil
}

func String(loc string, layout string) (string, error) {

	timer, err := Now(loc, layout)
	if err != nil {
		return "", err
	}

	return timer.Format(layout), nil
}
