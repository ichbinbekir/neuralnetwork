package neuralnetwork

type Neuron struct {
	Wegihts    []float64
	Bias       float64
	Activation string
}

func (n Neuron) run(inputs ...float64) float64 {
	output := n.Bias
	for w, input := range inputs {
		output += input * n.Wegihts[w]
	}

	return Activations[n.Activation].Normal(output)
}

// func (n *Neuron) train(learningRate, target float64, inputs ...float64) {
// 	output := n.run(inputs...)
// 	gradient := learningRate * mseDerivative(target, output) * Activations[n.Activation].Derivative(output)
// 	for w, input := range inputs {
// 		n.Wegihts[w] -= gradient * input
// 	}
// 	n.Bias -= gradient
// }
