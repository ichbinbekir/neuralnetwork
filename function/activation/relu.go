package activation

type ReLU struct{}

func (ReLU) Normal(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func (ReLU) Derivative(x float64) float64 {
	if x > 0 {
		return 1
	}
	return 0
}
