package gptimeframe

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func FromDuration(timeFrame string, layout string) (time.Duration, error) {

	re := regexp.MustCompile(layout)
	matches := re.FindStringSubmatch(timeFrame)

	if matches == nil {
		return 0, fmt.Errorf("invalid time frame format: %s", timeFrame)
	}

	num, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("invalid number in time frame: %s", matches[1])
	}

	switch matches[2] {
	case "s":
		return time.Duration(num) * time.Second, nil

	case "m":
		return time.Duration(num) * time.Minute, nil

	case "h":
		return time.Duration(num) * time.Hour, nil

	case "d":
		return time.Duration(num) * 24 * time.Hour, nil

	default:
		return 0, fmt.Errorf("invalid time frame unit: %s", matches[2])
	}
}
