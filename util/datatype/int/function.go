package gpint

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	gpslice "github.com/vagnerpraia/gopkg/util/datatype/slice"
)

func FromString(str string, scape []string, slice gpslice.Slice, re *regexp.Regexp) (int, error) {

	str = strings.Trim(str, " ")

	if slice.StringsContains(scape, str) {
		return 0, nil
	}

	if !re.MatchString(str) {
		return 0, errors.New("invalid format")
	}

	out, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return out, nil
}

func ReverseSign(in int) int {

	return in * -1
}

func Positive(in int) int {

	if in < 0 {
		return ReverseSign(in)
	}

	return in
}

func Negative(in int) int {

	if in > 0 {
		return ReverseSign(in)
	}

	return in
}

func Module(in int) int {

	if in < 0 {
		return ReverseSign(in)
	}

	return in
}
