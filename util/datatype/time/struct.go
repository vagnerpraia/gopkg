package gptime

import (
	"time"
)

type Time struct {
	location string
	layout   string
}

func NewTime(location string, layout string) *Time {

	return &Time{
		location: location,
		layout:   layout,
	}
}

func (util *Time) ToStringDefault(in time.Time) string {

	return ToStringDefault(in)
}

func (util *Time) ToString(in time.Time) string {

	return ToString(in, util.layout)
}

func (util *Time) FromString(str string) (time.Time, error) {

	return FromString(str, util.layout)
}

func (util *Time) FromYearOnly(in time.Time) time.Time {

	return FromYearOnly(in)
}

func (util *Time) FromMounthOnly(in time.Time) time.Time {

	return FromMounthOnly(in)
}

func (util *Time) FromWeekOnly(in time.Time) time.Time {

	return FromWeekOnly(in)
}

func (util *Time) FromDateOnly(in time.Time) time.Time {

	return FromDateOnly(in)
}

func (util *Time) FromTimeOnly(in time.Time) time.Time {

	return FromTimeOnly(in)
}

func (util *Time) FromHourOnly(in time.Time) time.Time {

	return FromHourOnly(in)
}

func (util *Time) FromMinuteOnly(in time.Time) time.Time {

	return FromMinuteOnly(in)
}

func (util *Time) FromSecondOnly(in time.Time) time.Time {

	return FromSecondOnly(in)
}

func (util *Time) FromHourUntil(in time.Time) time.Time {

	return FromHourUntil(in)
}

func (util *Time) FromMinuteUntil(in time.Time) time.Time {

	return FromMinuteUntil(in)
}

func (util *Time) FromSecondUntil(in time.Time) time.Time {

	return FromSecondUntil(in)
}

func (util *Time) MergeTimeInDate(date time.Time, schedule time.Time) time.Time {

	return MergeTimeInDate(date, schedule)
}

func (util *Time) DateTimeBeforeOrEqual(base time.Time, check time.Time) bool {

	return DateTimeBeforeOrEqual(base, check)
}

func (util *Time) DateTimeAfterOrEqual(base time.Time, check time.Time) bool {

	return DateTimeAfterOrEqual(base, check)
}

func (util *Time) DateBeforeOrEqual(base time.Time, check time.Time) bool {

	return DateBeforeOrEqual(base, check)
}

func (util *Time) DateAfterOrEqual(base time.Time, check time.Time) bool {

	return DateAfterOrEqual(base, check)
}

func (util *Time) DateBefore(base time.Time, check time.Time) bool {

	return DateBefore(base, check)
}

func (util *Time) DateAfter(base time.Time, check time.Time) bool {

	return DateAfter(base, check)
}

func (util *Time) DateEqual(base time.Time, check time.Time) bool {

	return DateEqual(base, check)
}

func (util *Time) TimeBeforeOrEqual(base time.Time, check time.Time) bool {

	return TimeBeforeOrEqual(base, check)
}

func (util *Time) TimeAfterOrEqual(base time.Time, check time.Time) bool {

	return TimeAfterOrEqual(base, check)
}

func (util *Time) TimeBefore(base time.Time, check time.Time) bool {

	return TimeBefore(base, check)
}

func (util *Time) TimeAfter(base time.Time, check time.Time) bool {

	return TimeAfter(base, check)
}

func (util *Time) TimeEqual(base time.Time, check time.Time) bool {

	return TimeEqual(base, check)
}

func (util *Time) IsSunday(in time.Time) bool {

	return IsSunday(in)
}

func (util *Time) IsMonday(in time.Time) bool {

	return IsMonday(in)
}

func (util *Time) IsTuesday(in time.Time) bool {

	return IsTuesday(in)
}

func (util *Time) IsWednesday(in time.Time) bool {

	return IsWednesday(in)
}

func (util *Time) IsThursday(in time.Time) bool {

	return IsThursday(in)
}

func (util *Time) IsFriday(in time.Time) bool {

	return IsFriday(in)
}

func (util *Time) IsSaturday(in time.Time) bool {

	return IsSaturday(in)
}
