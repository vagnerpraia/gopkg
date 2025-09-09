package gpfloat64

import (
	"regexp"
)

type Float64 struct {
	prec     int
	sep      string
	currency string
	scape    []string
	re       *regexp.Regexp
}

func NewFloat64(prec int, sep string, currency string) (*Float64, error) {

	re, err := regexp.Compile("(^[0-9]*$)|(^[0-9]*[\\" + sep + "][0-9]{0,2})")
	if err != nil {
		return nil, err
	}

	return &Float64{
		prec:     prec,
		sep:      sep,
		currency: currency,
		re:       re,
	}, nil
}

func (f *Float64) ToStringDefault(flt float64) string {

	return ToStringDefault(flt)
}

func (f *Float64) ToString(flt float64) string {

	return ToString(flt, f.prec, f.sep)
}

func (f *Float64) ToMonetary(flt float64) string {

	return ToMonetary(flt, f.currency, f.prec, f.sep)
}

func (f *Float64) ToPercentage(flt float64) string {

	return ToPercentage(flt, f.prec, f.sep)
}

func (f *Float64) FromString(str string) (float64, error) {

	return FromString(str, f.sep, f.scape, f.re)
}

func (f *Float64) Normalize(flt float64, normalizer float64) (float64, error) {

	return Normalize(flt, normalizer, f.prec)
}

func (f *Float64) Positive(flt float64) float64 {

	return Positive(flt)
}

func (f *Float64) Negative(flt float64) float64 {

	return Negative(flt)
}

func (f *Float64) Module(flt float64) float64 {

	return Module(flt)
}

func (f *Float64) IsIntegral(flt float64) bool {

	return IsIntegral(flt)
}

func (f *Float64) ReverseSign(flt float64) float64 {

	return ReverseSign(flt)
}

func (f *Float64) Round(flt float64, prec int) float64 {

	return Round(flt, prec)
}
