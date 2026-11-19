package gplowercamel

type Kebab struct {
}

func NewKebab() *Kebab {

	return &Kebab{}
}

func (crypt *Kebab) FromWords(value string) string {

	return FromWords(value)
}

func (crypt *Kebab) ToWords(value string) string {

	return ToWords(value)
}

func (crypt *Kebab) FromKebab(value string) string {

	return FromKebab(value)
}
