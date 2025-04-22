package gpmodel

type Float64Weighted struct {
	Weight float64
	Value  float64
}

func NewFloat64Weighted(value int, weight float64) *Float64Weighted {

	return &Float64Weighted{
		Value:  float64(value),
		Weight: weight,
	}
}
