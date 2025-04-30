package gpsymbol

type Symbol struct {
}

func NewSymbol(normalizer float64) *Symbol {

	return &Symbol{}
}

func (this *Symbol) Normalize(symbol string) string {

	return Normalize(symbol)
}
