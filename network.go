package neuralnetwork

type Network struct {
	Layers       []Layer
	LearningRate float64
	Loss         string
}

func (n Network) Run(inputs ...float64) []float64 {
	outputs := n.Layers[0].run(inputs...)
	for l, layer := range n.Layers {
		if l == 0 {
			continue
		}
		outputs = layer.run(outputs...)
	}
	return outputs
}

func (n *Network) Train(targets []float64, inputs ...float64) {
	outputs := make([][]float64, len(n.Layers))
	outputs[0] = n.Layers[0].run(inputs...)

	for l, layer := range n.Layers {
		if l == 0 {
			continue
		}
		outputs[l] = layer.run(outputs[l-1]...)
	}

	gradients := make([][]float64, len(n.Layers))
	for l := len(n.Layers) - 1; l >= 0; l-- {
		gradients[l] = make([]float64, len(n.Layers[l].Neurons))
		if len(n.Layers)-1 == l {
			for i, neuron := range n.Layers[l].Neurons {
				gradients[l][i] = Activations[neuron.Activation].Derivative(outputs[l][i]) * Losses[n.Loss].Derivative(targets[i], outputs[l][i])
			}
			continue
		}

		for j, neuron := range n.Layers[l].Neurons {
			var sum float64
			for i, neu := range n.Layers[l+1].Neurons {
				sum += gradients[l+1][i] * neu.Wegihts[j]
			}
			gradients[l][j] = sum * Activations[neuron.Activation].Derivative(outputs[l][j])
		}
	}

	for l := range n.Layers {
		if l == 0 {
			for i := range n.Layers[l].Neurons {
				for w := range n.Layers[l].Neurons[i].Wegihts {
					n.Layers[l].Neurons[i].Wegihts[w] -= n.LearningRate * gradients[l][i] * inputs[w]
				}
				n.Layers[l].Neurons[i].Bias -= n.LearningRate * gradients[l][i]
			}
			continue
		}

		for i := range n.Layers[l].Neurons {
			for w := range n.Layers[l].Neurons[i].Wegihts {
				n.Layers[l].Neurons[i].Wegihts[w] -= n.LearningRate * gradients[l][i] * outputs[l-1][w]
			}
			n.Layers[l].Neurons[i].Bias -= n.LearningRate * gradients[l][i]
		}
	}
}
