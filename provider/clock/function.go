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
