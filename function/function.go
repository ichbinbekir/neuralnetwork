package function

type Function interface {
	Normal(float64) float64
	Derivative(float64) float64
}
