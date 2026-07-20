package gpstring

type String struct {
}

func NewInt() (*String, error) {

	return &String{}, nil
}

func (that *String) NormalizePath(str string) string {

	return NormalizePath(str)
}
