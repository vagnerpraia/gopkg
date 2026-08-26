package gptemplate

import (
	"errors"
)

type Delimiter string

const (
	delimiterCurlyBracketsAlias  = "curly-brackets"
	delimiterSquareBracketsAlias = "square-brackets"
	delimiterRoundBracketsAlias  = "round-brackets"

	DelimiterCurlyBrackets  Delimiter = delimiterCurlyBracketsAlias
	DelimiterSquareBrackets Delimiter = delimiterSquareBracketsAlias
	DelimiterRoundBrackets  Delimiter = delimiterRoundBracketsAlias
)

func NewDelimiter(str string) (Delimiter, error) {

	switch str {
	case delimiterCurlyBracketsAlias:
		return DelimiterCurlyBrackets, nil

	case delimiterSquareBracketsAlias:
		return DelimiterSquareBrackets, nil

	case delimiterRoundBracketsAlias:
		return DelimiterRoundBrackets, nil
	}

	return "", errors.New("os not found")
}

func (that Delimiter) String() string {

	switch that {
	case DelimiterCurlyBrackets:
		return delimiterCurlyBracketsAlias

	case DelimiterSquareBrackets:
		return delimiterSquareBracketsAlias

	case DelimiterRoundBrackets:
		return delimiterRoundBracketsAlias
	}

	return ""
}

func (that Delimiter) Pair() string {

	switch that {
	case DelimiterCurlyBrackets:
		return "{}"

	case DelimiterSquareBrackets:
		return "[]"

	case DelimiterRoundBrackets:
		return "()"
	}

	return ""
}

func (that Delimiter) Open() byte {

	return that.Pair()[0]
}

func (that Delimiter) Close() byte {

	return that.Pair()[1]
}

func (that Delimiter) IsCurlyBrackets() bool {

	return that == DelimiterCurlyBrackets
}

func (that Delimiter) IsSquareBrackets() bool {

	return that == DelimiterSquareBrackets
}

func (that Delimiter) IsRoundBrackets() bool {

	return that == DelimiterRoundBrackets
}
