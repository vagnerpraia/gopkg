package gpparity

type Parity struct {
}

func NewParity() (*Parity, error) {

	return &Parity{}, nil
}

func (that *Parity) IntIsEven(num int) bool {

	return IntIsEven(num)
}

func (that *Parity) Float32IsEven(num float32) bool {

	return Float32IsEven(num)
}

func (that *Parity) Float64IsEven(num float64) bool {

	return Float64IsEven(num)
}

func (that *Parity) IntIsOdd(num int) bool {

	return IntIsOdd(num)
}

func (that *Parity) Float32IsOdd(num float32) bool {

	return Float32IsOdd(num)
}

func (that *Parity) Float64IsOdd(num float64) bool {

	return Float64IsOdd(num)
}

func (that *Parity) LengthIntsIsEven(slice []int) bool {

	return LengthIntsIsEven(slice)
}

func (that *Parity) LengthIntsIsOdd(slice []int) bool {

	return LengthIntsIsOdd(slice)
}

func (that *Parity) LengthFloat32sIsEven(slice []float32) bool {

	return LengthFloat32sIsEven(slice)
}

func (that *Parity) LengthFloat32sIsOdd(slice []float32) bool {

	return LengthFloat32sIsOdd(slice)
}

func (that *Parity) LengthFloat64sIsEven(slice []float64) bool {

	return LengthFloat64sIsEven(slice)
}

func (that *Parity) LengthFloat64sIsOdd(slice []float64) bool {

	return LengthFloat64sIsOdd(slice)
}

func (that *Parity) LengthInterfacesIsEven(slice []interface{}) bool {

	return LengthInterfacesIsEven(slice)
}

func (that *Parity) LengthInterfacesIsOdd(slice []interface{}) bool {

	return LengthInterfacesIsOdd(slice)
}
