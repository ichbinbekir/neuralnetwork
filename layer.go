package neuralnetwork

type Layer struct {
	Neurons []Neuron
}

func (l Layer) run(inputs ...float64) []float64 {
	outputs := make([]float64, len(l.Neurons))
	for o, neuron := range l.Neurons {
		outputs[o] = neuron.run(inputs...)
	}
	return outputs
}

// func (l *Layer) train(learningRate float64, targets []float64, inputs ...float64) {
// 	for i, n := range l.Neurons {
// 		n.train(learningRate, targets[i], inputs...)
// 	}
// }
