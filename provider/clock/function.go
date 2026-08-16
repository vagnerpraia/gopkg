package gpclock

import (
	"time"
)

func Now(loc string, layout string) (*time.Time, error) {

	location, err := time.LoadLocation(loc)
	if err != nil {
		return nil, err
	}

	t := time.Now().In(location).Format(layout)
	now, err := time.Parse(layout, t)
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
