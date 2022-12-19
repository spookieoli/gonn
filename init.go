package gonn

import (
	"github.com/spookieoli/gonn/layers"
	"github.com/spookieoli/gonn/utils"
)

// Create a new Gonn instance
func New(lc *[]utils.LayerConfig, threads int, name string) *Gonn {
	// create the Channel for the Compution
	com := make(chan utils.Payload)

	// Start the Threads for the Compution
	for i := 0; i < threads; i++ {
		go utils.Work(com)
	}

	// Create the Neural Network
	nn := &utils.NeuralNetwork{Threads: threads, Com: com}

	// Create the Gonn Instance
	gonn := &Gonn{NN: nn, Name: name}

	// Create the Layers
	for idx, layer := range *lc {
		switch layer.LayerType.(type) {
		case *layers.Input:
			// Create a new Input layer
			*gonn.NN.Layers = append(*gonn.NN.Layers, layers.NewInput(layer.Neurons, com))
		case *layers.Dense:
			// Create a new Dense layer
			*gonn.NN.Layers = append(*gonn.NN.Layers, layers.NewDense(layer.Neurons, layer.Activation, layer.Bias, layer.Name, layer.DropOut, &(*gonn.NN.Layers)[idx-1], com))
		}
	}
	// Set the right Inputs for every Layer
	gonn.SetInputs()
	// Init all the Weights
	gonn.CreateWeights()
	// Return to Caller
	return gonn
}
