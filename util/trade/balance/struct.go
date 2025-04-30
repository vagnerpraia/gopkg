package vpbalance

type Balance struct {
	normalizer float64
}

func NewBalance(normalizer float64) *Balance {

	return &Balance{
		normalizer: normalizer,
	}
}

func (this *Balance) Calculate(ask float64, bid float64) float64 {

	return Calculate(bid, ask, this.normalizer)
}
