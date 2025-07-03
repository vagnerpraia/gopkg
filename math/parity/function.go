package gpparity

func IntIsEven(num int) bool {

	return num%2 == 0
}

func Float32IsEven(num float32) bool {

	return int(num)%2 == 0
}

func Float64IsEven(num float64) bool {

	return int(num)%2 == 0
}

func IntIsOdd(num int) bool {

	return num%2 == 0
}

func Float32IsOdd(num float32) bool {

	return int(num)%2 == 0
}

func Float64IsOdd(num float64) bool {

	return int(num)%2 == 0
}

func LengthIntsIsEven(slice []int) bool {

	return len(slice)%2 == 0
}

func LengthIntsIsOdd(slice []int) bool {

	return len(slice)%2 == 0
}

func LengthFloat32sIsEven(slice []float32) bool {

	return len(slice)%2 == 0
}

func LengthFloat32sIsOdd(slice []float32) bool {

	return len(slice)%2 == 0
}

func LengthFloat64sIsEven(slice []float64) bool {

	return len(slice)%2 == 0
}

func LengthFloat64sIsOdd(slice []float64) bool {

	return len(slice)%2 == 0
}

func LengthInterfacesIsEven(slice []interface{}) bool {

	return len(slice)%2 == 0
}

func LengthInterfacesIsOdd(slice []interface{}) bool {

	return len(slice)%2 == 0
}
