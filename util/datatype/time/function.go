package gptime

import (
	"time"
)

func ToStringDefault(in time.Time) string {

	return in.Format(LAYOUT_DEFAULT)
}

func ToString(in time.Time, layout string) string {

	return in.Format(layout)
}

func FromString(str string, layout string) (time.Time, error) {

	out, err := time.Parse(layout, str)
	if err != nil {
		return out, err
	}

	return out, err
}

func FromYearOnly(in time.Time) time.Time {

	return time.Date(
		in.Year(),
		1,
		1,
		0,
		0,
		0,
		0,
		in.Location(),
	)
}

func FromMounthOnly(in time.Time) time.Time {

	return time.Date(
		in.Year(),
		in.Month(),
		1,
		0,
		0,
		0,
		0,
		in.Location(),
	)
}

func FromWeekOnly(in time.Time) time.Time {

	in = FromDateOnly(in)

	for {
		if in.Weekday() == time.Sunday {
			break
		}

		in = in.AddDate(0, 0, -1)
	}

	return in
}

func FromDateOnly(in time.Time) time.Time {

	return time.Date(
		in.Year(),
		in.Month(),
		in.Day(),
		0,
		0,
		0,
		0,
		in.Location(),
	)
}

func FromTimeOnly(in time.Time) time.Time {

	return time.Date(
		0,
		1,
		1,
		in.Hour(),
		in.Minute(),
		in.Second(),
		in.Nanosecond(),
		in.Location(),
	)
}

func FromHourOnly(in time.Time) time.Time {

	return time.Date(
		0,
		1,
		1,
		in.Hour(),
		0,
		0,
		0,
		in.Location(),
	)
}

func FromMinuteOnly(in time.Time) time.Time {

	return time.Date(
		0,
		1,
		1,
		0,
		in.Minute(),
		0,
		0,
		in.Location(),
	)
}

func FromSecondOnly(in time.Time) time.Time {

	return time.Date(
		0,
		1,
		1,
		0,
		0,
		in.Second(),
		0,
		in.Location(),
	)
}

func FromHourUntil(in time.Time) time.Time {

	return time.Date(
		0,
		1,
		1,
		in.Hour(),
		0,
		0,
		0,
		in.Location(),
	)
}

func FromMinuteUntil(in time.Time) time.Time {

	return time.Date(
		0,
		1,
		1,
		in.Hour(),
		in.Minute(),
		0,
		0,
		in.Location(),
	)
}

func FromSecondUntil(in time.Time) time.Time {

	return time.Date(
		0,
		1,
		1,
		in.Hour(),
		in.Minute(),
		in.Second(),
		0,
		in.Location(),
	)
}

func MergeTimeInDate(date time.Time, schedule time.Time) time.Time {

	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		schedule.Hour(),
		schedule.Minute(),
		schedule.Second(),
		schedule.Nanosecond(),
		schedule.Location(),
	)
}

func DateTimeBeforeOrEqual(base time.Time, check time.Time) bool {

	return base.Before(check) || base.Equal(check)
}

func DateTimeAfterOrEqual(base time.Time, check time.Time) bool {

	return base.After(check) || base.Equal(check)
}

func DateBeforeOrEqual(base time.Time, check time.Time) bool {

	dateBase := FromDateOnly(base)
	dateCheck := FromDateOnly(check)

	return dateBase.Before(dateCheck) || dateBase.Equal(dateCheck)
}

func DateAfterOrEqual(base time.Time, check time.Time) bool {

	dateBase := FromDateOnly(base)
	dateCheck := FromDateOnly(check)

	return dateBase.After(dateCheck) || dateBase.Equal(dateCheck)
}

func DateBefore(base time.Time, check time.Time) bool {

	dateBase := FromDateOnly(base)
	dateCheck := FromDateOnly(check)

	return dateBase.Before(dateCheck)
}

func DateAfter(base time.Time, check time.Time) bool {

	dateBase := FromDateOnly(base)
	dateCheck := FromDateOnly(check)

	return dateBase.After(dateCheck)
}

func DateEqual(base time.Time, check time.Time) bool {

	dateBase := FromDateOnly(base)
	dateCheck := FromDateOnly(check)

	return dateBase.Equal(dateCheck)
}

func TimeBeforeOrEqual(base time.Time, check time.Time) bool {

	timeBase := FromTimeOnly(base)
	timeCheck := FromTimeOnly(check)

	return timeBase.Before(timeCheck) || timeBase.Equal(timeCheck)
}

func TimeAfterOrEqual(base time.Time, check time.Time) bool {

	timeBase := FromTimeOnly(base)
	timeCheck := FromTimeOnly(check)

	return timeBase.After(timeCheck) || timeBase.Equal(timeCheck)
}

func TimeBefore(base time.Time, check time.Time) bool {

	dateBase := FromTimeOnly(base)
	dateCheck := FromTimeOnly(check)

	return dateBase.Before(dateCheck)
}

func TimeAfter(base time.Time, check time.Time) bool {

	dateBase := FromTimeOnly(base)
	dateCheck := FromTimeOnly(check)

	return dateBase.After(dateCheck)
}

func TimeEqual(base time.Time, check time.Time) bool {

	dateBase := FromTimeOnly(base)
	dateCheck := FromTimeOnly(check)

	return dateBase.Equal(dateCheck)
}

func IsSunday(in time.Time) bool {

	return in.Weekday() == time.Sunday
}

func IsMonday(in time.Time) bool {

	return in.Weekday() == time.Monday
}

func IsTuesday(in time.Time) bool {

	return in.Weekday() == time.Tuesday
}

func IsWednesday(in time.Time) bool {

	return in.Weekday() == time.Wednesday
}

func IsThursday(in time.Time) bool {

	return in.Weekday() == time.Thursday
}

func IsFriday(in time.Time) bool {

	return in.Weekday() == time.Friday
}

func IsSaturday(in time.Time) bool {

	return in.Weekday() == time.Saturday
}
