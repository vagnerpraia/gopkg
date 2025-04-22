package gpslice

import (
	"time"

	gpmodel "github.com/vagnerpraia/gopkg/model"
)

func IntsEquals(a []int, b []int) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func IntsContains(a []int, b int) bool {

	for _, c := range a {
		if b == c {
			return true
		}
	}

	return false
}

func IntsRemove(a []int, pos int) []int {

	if pos > len(a)-1 {
		return a
	}

	if pos == len(a) {
		return a[:len(a)-1]
	}

	return append(a[:pos], a[pos+1:]...)
}

func IntsMeanSimple(a ...int) float64 {

	sum := 0.0

	for _, b := range a {
		sum += float64(b)
	}

	return sum / float64(len(a))
}

func Float64sEquals(a []float64, b []float64) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Float64sContains(a []float64, b float64) bool {

	for _, c := range a {
		if b == c {
			return true
		}
	}

	return false
}

func Float64sRemove(a []float64, pos int) []float64 {

	if pos > len(a)-1 {
		return a
	}

	if pos == len(a) {
		return a[:len(a)-1]
	}

	return append(a[:pos], a[pos+1:]...)
}

func Float64sMeanSimple(a ...float64) float64 {

	sum := 0.0

	for _, b := range a {
		sum += b
	}

	return sum / float64(len(a))
}

func WeightValuesMeanWeighted(a ...gpmodel.Float64Weighted) float64 {

	sum := 0.0
	divisor := 0.0

	for _, b := range a {
		sum += (b.Value * b.Weight)
		divisor += b.Weight
	}

	return sum / divisor
}

func StringsEquals(a []string, b []string) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func StringsContains(a []string, b string) bool {

	for _, c := range a {
		if b == c {
			return true
		}
	}

	return false
}

func StringsRemove(a []string, pos int) []string {

	if pos > len(a)-1 {
		return a
	}

	if pos == len(a) {
		return a[:len(a)-1]
	}

	return append(a[:pos], a[pos+1:]...)
}

func TimesEquals(a []time.Time, b []time.Time) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TimesContains(a []time.Time, b time.Time) bool {

	for _, c := range a {
		if b == c {
			return true
		}
	}

	return false
}

func TimesRemove(a []time.Time, pos int) []time.Time {

	if pos > len(a)-1 {
		return a
	}

	if pos == len(a) {
		return a[:len(a)-1]
	}

	return append(a[:pos], a[pos+1:]...)
}

func InterfacesEquals(a []interface{}, b []interface{}) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func InterfacesContains(a []interface{}, b interface{}) bool {

	for _, c := range a {
		if b == c {
			return true
		}
	}

	return false
}

func InterfacesRemove(a []interface{}, pos int) []interface{} {

	if pos > len(a)-1 {
		return a
	}

	if pos == len(a) {
		return a[:len(a)-1]
	}

	return append(a[:pos], a[pos+1:]...)
}
