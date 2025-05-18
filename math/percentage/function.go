package gppercentage

import (
	"errors"

	gpfloat64 "github.com/vagnerpraia/gopkg/util/datatype/float64"
	gpint "github.com/vagnerpraia/gopkg/util/datatype/int"
)

func PercentualInts(total int, partial int) float64 {

	return (float64(partial) / float64(total)) * 100
}

func PercentualFloat64s(total float64, partial float64) float64 {

	return (partial / total) * 100
}

func ValueInt(value int, percentage float64) float64 {

	return (percentage / 100) * float64(value)
}

func ValueSumInt(value int, percentage float64) float64 {

	return float64(value) + (ValueInt(value, percentage))
}

func ValueFloat64(value float64, percentage float64) float64 {

	return (percentage / 100) * value
}

func ValueSumFloat64(value float64, percentage float64) float64 {

	return value + (ValueFloat64(value, percentage))
}

func VariationInts(num1 int, num2 int) (float64, error) {

	if num1 == 0 {
		return 0, errors.New("invalid value")
	}

	if num1 < 0 {
		divisor := gpint.ReverseSign(num1)
		return ((float64(num2) - float64(num1)) / float64(divisor)) * 100, nil
	}

	return ((float64(num2) - float64(num1)) / float64(num1)) * 100, nil
}

func VariationFloat64s(num1 float64, num2 float64) (float64, error) {

	if num1 == 0 {
		return 0, errors.New("invalid value")
	}

	if num1 < 0 {
		divisor := gpfloat64.ReverseSign(num1)
		return ((float64(num2) - float64(num1)) / float64(divisor)) * 100, nil
	}

	return ((float64(num2) - float64(num1)) / float64(num1)) * 100, nil
}
