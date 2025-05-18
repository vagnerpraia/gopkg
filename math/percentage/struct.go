package gppercentage

type Percentage struct {
}

func NewPercentage() (*Percentage, error) {

	return &Percentage{}, nil
}

func (this *Percentage) PercentualInts(total int, partial int) float64 {

	return PercentualInts(total, partial)
}

func (this *Percentage) PercentualFloat64s(total float64, partial float64) float64 {

	return PercentualFloat64s(total, partial)
}

func (this *Percentage) ValueInt(value int, percentage float64) float64 {

	return ValueInt(value, percentage)
}

func (this *Percentage) ValueFloat64(value float64, percentage float64) float64 {

	return ValueFloat64(value, percentage)
}

func (this *Percentage) VariationInts(num1 int, num2 int) (float64, error) {

	return VariationInts(num1, num2)
}

func (this *Percentage) VariationFloat64s(num1 float64, num2 float64) (float64, error) {

	return VariationFloat64s(num1, num2)
}
