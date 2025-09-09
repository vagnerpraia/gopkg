package gpfloat64

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"

	gpslice "github.com/vagnerpraia/gopkg/util/datatype/slice"
)

func ToStringDefault(flt float64) string {

	return strings.Replace(strconv.FormatFloat(flt, 'f', PREC, 64), ".", SEP, -1)
}

func ToString(flt float64, prec int, sep string) string {

	return strings.Replace(strconv.FormatFloat(flt, 'f', prec, 64), ".", sep, -1)
}

func ToMonetary(i float64, currency string, prec int, sep string) string {

	return currency + " " + strings.Replace(strconv.FormatFloat(i, 'f', prec, 64), ".", sep, -1)
}

func ToPercentage(i float64, prec int, sep string) string {

	return strings.Replace(strconv.FormatFloat(i, 'f', prec, 64), ".", sep, -1) + "%"
}

func FromString(str string, sep string, scape []string, re *regexp.Regexp) (float64, error) {

	str = strings.Trim(str, " ")

	if gpslice.StringsContains(scape, str) {
		return 0, nil
	}

	if !re.MatchString(str) {
		return 0, errors.New("invalid format")
	}

	s := strings.Replace(str, sep, ".", -1)

	out, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return out, nil
}

func Normalize(flt float64, normalizer float64, prec int) (float64, error) {

	if normalizer > 0.0 {
		flt = flt / normalizer
	}

	str := strconv.FormatFloat(flt, 'f', prec, 64)

	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0, err
	}

	return flt, nil
}

func Positive(flt float64) float64 {

	if flt < 0 {
		return ReverseSign(flt)
	}

	return flt
}

func Negative(flt float64) float64 {

	if flt > 0 {
		return ReverseSign(flt)
	}

	return flt
}

func Module(flt float64) float64 {

	if flt < 0 {
		return ReverseSign(flt)
	}

	return flt
}

func IsIntegral(flt float64) bool {
	return flt == float64(int(flt))
}

func ReverseSign(flt float64) float64 {

	return flt * -1
}

func Round(flt float64, prec int) float64 {

	factor := math.Pow(10, float64(prec))

	return math.Round(flt*factor) / factor
}
