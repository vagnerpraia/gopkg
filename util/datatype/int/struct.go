package gpint

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	gpslice "github.com/vagnerpraia/gopkg/util/datatype/slice"
)

type Int struct {
	scape []string
	re    *regexp.Regexp
}

func NewInt(scape []string) (*Int, error) {

	re, err := regexp.Compile("[0-9]*")
	if err != nil {
		return nil, err
	}

	return &Int{
		scape: scape,
		re:    re,
	}, nil
}

func (this *Int) FromString(str string) (int, error) {

	str = strings.Trim(str, " ")

	if gpslice.StringsContains(this.scape, str) {
		return 0, nil
	}

	if !this.re.MatchString(str) {
		return 0, errors.New("invalid format")
	}

	out, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return out, nil
}

func (this *Int) ReverseSign(in int) int {

	return ReverseSign(in)
}

func (this *Int) Positive(in int) int {

	return Positive(in)
}

func (this *Int) Negative(in int) int {

	return Negative(in)
}

func (this *Int) Module(in int) int {

	return Module(in)
}
