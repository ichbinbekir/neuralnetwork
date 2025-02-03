package neuralnetwork

import "github.com/ichbinbekir/neuralnetwork/function"

// prrrt 😂
type localFunc interface {
	Normal(float64, float64) float64
	Derivative(float64, float64) float64
}

var Activations map[string]function.Function = make(map[string]function.Function)
var Losses map[string]localFunc = make(map[string]localFunc)
