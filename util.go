package neuralnetwork

import "math/rand/v2"

func RandomNeurons(nWeight, len int, activation string) []Neuron {
	neurons := make([]Neuron, len)
	for n := range neurons {
		neurons[n].Wegihts = make([]float64, nWeight)
		neurons[n].Activation = activation
		for w := range neurons[n].Wegihts {
			neurons[n].Wegihts[w] = rand.Float64()*2 - 1
		}
	}
	return neurons
}
