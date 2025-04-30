package vpbalance

func Calculate(ask float64, bid float64, normalizer float64) float64 {

	if normalizer > 0 {
		return (bid - ask) * normalizer
	}

	return (bid - ask)
}
