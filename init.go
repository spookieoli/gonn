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
	// Check if the first layer is an Input layer
	if !gonn.CheckInput() {
		panic("Your first layer must be an Input layer")
	}

	// Create the Layers
	for idx, layer := range *lc {
		switch layer.LayerType.(type) {
		case *layers.Input:
			// Create a new Input layer
			*gonn.NN.Layers = append(*gonn.NN.Layers, layers.NewInput(layer.Neurons))
		case *layers.Dense:
			// Create a new Dense layer
			*gonn.NN.Layers = append(*gonn.NN.Layers, layers.NewDense(layer.Neurons, layer.Activation, layer.Bias, layer.Name, layer.DropOut, &(*gonn.NN.Layers)[idx-1]))
		}
	}
	return gonn
}
