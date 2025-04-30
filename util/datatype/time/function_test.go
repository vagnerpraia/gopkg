package gptime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FunctionSuite struct {
	suite.Suite
}

func TestFunctionSuite(t *testing.T) {

	suite.Run(t, new(FunctionSuite))
}

func (s *FunctionSuite) SetupSuite() {}

func (s *FunctionSuite) TestFunction_FromWeekOnly() {

	type args struct {
		date time.Time
	}

	tests := []struct {
		name     string
		args     args
		expected time.Time
	}{
		{
			name: "FromWeekOnly 1",
			args: args{
				date: time.Date(2022, time.December, 25, 20, 50, 0, 0, time.Local),
			},
			expected: time.Date(2022, time.December, 25, 0, 0, 0, 0, time.Local),
		},
		{
			name: "FromWeekOnly 2",
			args: args{
				date: time.Date(2022, time.December, 28, 12, 0, 0, 0, time.Local),
			},
			expected: time.Date(2022, time.December, 25, 0, 0, 0, 0, time.Local),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			actual := FromWeekOnly(tt.args.date)

			fmt.Println(tt.args.date)
			fmt.Println(actual)

			s.Assert().Equal(tt.expected, actual)
		})
	}
}
