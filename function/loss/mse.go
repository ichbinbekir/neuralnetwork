package loss

import "math"

type MSE struct{}

func (MSE) Normal(target, output float64) float64 {
	return math.Pow(output-target, 2) / 2
}

func (MSE) Derivative(target, output float64) float64 {
	return output - target
}
