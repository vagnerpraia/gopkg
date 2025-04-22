package gpslice

import (
	"time"

	gpmodel "github.com/vagnerpraia/gopkg/model"
)

type Slice struct {
}

func NewSlice() Slice {

	return Slice{}
}

func (util *Slice) IntsEquals(a []int, b []int) bool {

	return IntsEquals(a, b)
}

func (util *Slice) IntsContains(a []int, b int) bool {

	return IntsContains(a, b)
}

func (util *Slice) IntsRemove(a []int, b int) []int {

	return IntsRemove(a, b)
}

func (util *Slice) IntsMeanSimple(a ...int) float64 {

	return IntsMeanSimple(a...)
}

func (util *Slice) Float64sEquals(a []float64, b []float64) bool {

	return Float64sEquals(a, b)
}

func (util *Slice) Float64sContains(a []float64, b float64) bool {

	return Float64sContains(a, b)
}

func (util *Slice) Float64sRemove(a []float64, b int) []float64 {

	return Float64sRemove(a, b)
}

func (util *Slice) Float64sMeanSimple(a ...float64) float64 {

	return Float64sMeanSimple(a...)
}

func (util *Slice) WeightValuesMeanWeighted(a ...gpmodel.Float64Weighted) float64 {

	return WeightValuesMeanWeighted(a...)
}

func (util *Slice) StringsEquals(a []string, b []string) bool {

	return StringsEquals(a, b)
}

func (util *Slice) StringsContains(a []string, b string) bool {

	return StringsContains(a, b)
}

func (util *Slice) StringsRemove(a []string, b int) []string {

	return StringsRemove(a, b)
}

func (util *Slice) TimesEquals(a []time.Time, b []time.Time) bool {

	return TimesEquals(a, b)
}

func (util *Slice) TimesContains(a []time.Time, b time.Time) bool {

	return TimesContains(a, b)
}

func (util *Slice) TimesRemove(a []time.Time, pos int) []time.Time {

	return TimesRemove(a, pos)
}

func (util *Slice) InterfacesEquals(a []interface{}, b []interface{}) bool {

	return InterfacesEquals(a, b)
}

func (util *Slice) InterfacesContains(a []interface{}, b interface{}) bool {

	return InterfacesContains(a, b)
}

func (util *Slice) InterfacesRemove(a []interface{}, b int) []interface{} {

	return InterfacesRemove(a, b)
}
