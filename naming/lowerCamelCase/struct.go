package gplowercamelcase

type KebabCase struct {
}

func NewKebabCase() *KebabCase {

	return &KebabCase{}
}

func (crypt *KebabCase) FromWords(value string) string {

	return FromWords(value)
}

func (crypt *KebabCase) ToWords(value string) string {

	return ToWords(value)
}

func (crypt *KebabCase) FromKebabCase(value string) string {

	return FromKebabCase(value)
}
