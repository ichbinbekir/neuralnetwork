package activation

import "math"

type Sigmoid struct{}

func (Sigmoid) Normal(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func (Sigmoid) Derivative(x float64) float64 {
	return x * (1 - x)
}
