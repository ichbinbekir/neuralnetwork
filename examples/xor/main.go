package main

import (
	"log"

	nn "github.com/ichbinbekir/neuralnetwork"
	"github.com/ichbinbekir/neuralnetwork/function/activation"
	"github.com/ichbinbekir/neuralnetwork/function/loss"
)

const (
	epoch = 1000000
)

func main() {
	nn.Activations["sigmoid"] = activation.Sigmoid{}

	nn.Losses["mse"] = loss.MSE{}

	// You can create it by decoding it with json
	network := nn.Network{
		Layers: []nn.Layer{
			{Neurons: nn.RandomNeurons(2, 2, "sigmoid")},
			{Neurons: nn.RandomNeurons(2, 1, "sigmoid")},
		},
		LearningRate: 0.01,
		Loss:         "mse",
	}

	log.Println("Before training:")
	log.Printf("(0, 0): %.2f", network.Run(0, 0))
	log.Printf("(1, 0): %.2f", network.Run(1, 0))
	log.Printf("(0, 1): %.2f", network.Run(0, 1))
	log.Printf("(1, 1): %.2f", network.Run(1, 1))

	for range epoch {
		network.Train([]float64{0}, 0, 0)
		network.Train([]float64{1}, 1, 0)
		network.Train([]float64{1}, 0, 1)
		network.Train([]float64{0}, 1, 1)
	}
	// You can save it by encoding it with json

	log.Println("After training:")
	log.Printf("(0, 0): %.2f", network.Run(0, 0))
	log.Printf("(1, 0): %.2f", network.Run(1, 0))
	log.Printf("(0, 1): %.2f", network.Run(0, 1))
	log.Printf("(1, 1): %.2f", network.Run(1, 1))
}
