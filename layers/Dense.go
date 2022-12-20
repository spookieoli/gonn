package layers

import "github.com/spookieoli/gonn/utils"

type Dense struct {
	// Weights is a 2D matrix of weights
	Weights *[][]float64
	// Bias is a 1D matrix of bias
	Bias *[]float64
	// Neurons is the number of neurons in the layer
	Neurons int64
	// Activation is the activation function of the layer
	Activation func(any) any
	// Name is the name of the layer
	Name string
	// DropOut is the dropout rate of the layer
	DropOut float64
	// UseBias is a boolean that tells if the layer uses bias
	UsesBias bool
	// The Name of the Inputlayer
	InputLayername string
	// Layer that delivers Data to this Layer
	InputLayer *utils.Layer
	// The Communication Channel
	Com chan utils.Payload
}

// Create new Dense Layer and return it
func NewDense(neurons int64, activation func(any) any, bias bool, name string, dropout float64, lb *utils.Layer, com chan utils.Payload) *Dense {
	// if bias is true we have to increase the number of neurons by 1
	if bias {
		neurons++
	}

	// Return the new Dense Layer
	return &Dense{
		Neurons:    neurons,
		Activation: activation,
		UsesBias:   bias,
		Name:       name,
		DropOut:    dropout,
		Com:        com,
	}
}

// Set the Weights of every Neuron in the Layer
func (d *Dense) InitWeights() {
	// Create the Weights
	weights := make([][]float64, d.Neurons)
	for i := range weights {
		weights[i] = make([]float64, (*d.InputLayer).GetNeuronCount())
	}
	// Fill Weights with random Data
	for i := range weights {
		for j := range weights[i] {
			weights[i][j] = utils.RandFloat64(-1, 1)
		}
	}
}

// Get all the weights of the layer
func (d *Dense) GetWeights() *[][]float64 {
	return d.Weights
}

// Get the Name of the layer
func (d *Dense) GetName() string {
	return d.Name
}

// Get the dropout rate of the layer
func (d *Dense) GetDropOut() float64 {
	return d.DropOut
}

// Get the bias of the layer
func (d *Dense) GetBias() *[]float64 {
	return d.Bias
}

// Check if the layer uses bias
func (d *Dense) UseBias() bool {
	return d.UsesBias
}

// Calculate the output of the layer
func (d *Dense) GetOutputs() any {
	return nil
}

// Get the number of neurons in the layer
func (d *Dense) GetNeuronCount() int64 {
	return d.Neurons
}

// Set the Name of the Layer
func (d *Dense) SetName(name string) {
	d.Name = name
}

// Set the InputLayer
func (d *Dense) SetInputLayer(l *utils.Layer) {
	d.InputLayer = l
}

// Get the InputLayername
func (d *Dense) GetInputLayername() string {
	return d.InputLayername
}
