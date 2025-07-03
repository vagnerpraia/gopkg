package gpstats

type Percentage struct {
}

func NewPercentage() (*Percentage, error) {

	return &Percentage{}, nil
}

func (that *Percentage) IntsMean(ints []int) float64 {

	return IntsMean(ints)
}

func (that *Percentage) Float64sMean(floats []float64) float64 {

	return Float64sMean(floats)
}

func (that *Percentage) IntsMedian(ints []int) int {

	return IntsMedian(ints)
}

func (that *Percentage) Float64sMedian(floats []float64) float64 {

	return Float64sMedian(floats)
}

func (that *Percentage) IntsMode(ints []int) int {

	return IntsMode(ints)
}

func (that *Percentage) Float64sMode(floats []float64) float64 {

	return Float64sMode(floats)
}

func (that *Percentage) IntsModes(ints []int) []int {

	return IntsModes(ints)
}

func (that *Percentage) Float64sModes(floats []float64) []float64 {

	return Float64sModes(floats)
}
