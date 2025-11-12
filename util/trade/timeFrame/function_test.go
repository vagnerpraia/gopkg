package gptimeframe

import (
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

func (s *FunctionSuite) TestFunction_FromDuration() {

	type input struct {
		timeframe string
	}

	type want struct {
		duration    time.Duration
		expectError bool
	}

	tests := []struct {
		name  string
		input input
		want  want
	}{
		{
			name: "FromDuration 1",
			input: input{
				timeframe: "1ms",
			},
			want: want{
				duration:    1 * time.Millisecond,
				expectError: false,
			},
		},
		{
			name: "FromDuration 2",
			input: input{
				timeframe: "5s",
			},
			want: want{
				duration:    5 * time.Second,
				expectError: false,
			},
		},
		{
			name: "FromDuration 3",
			input: input{
				timeframe: "7m",
			},
			want: want{
				duration:    7 * time.Minute,
				expectError: false,
			},
		},
		{
			name: "FromDuration 4",
			input: input{
				timeframe: "3d",
			},
			want: want{
				duration:    3 * 24 * time.Hour,
				expectError: false,
			},
		},
		{
			name: "FromDuration 5",
			input: input{
				timeframe: "10md",
			},
			want: want{
				duration:    0,
				expectError: true,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			got, err := FromDuration(tt.input.timeframe)

			s.Assert().Equal(tt.want.expectError, err != nil)
			s.Assert().Equal(tt.want.duration, got)
		})
	}
}
