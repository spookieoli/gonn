package layers

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
}

// Create new Dense Layer and return it
func NewDense(neurons int64, activation func(any) any, bias bool, name string, dropout float64, neuronsbefore int64) *Dense {
	// if bias is true we have to increase the number of neurons by 1
	if bias {
		neurons++
	}

	// Create the Weights
	weights := make([][]float64, neurons)
	for i := range weights {
		weights[i] = make([]float64, neuronsbefore)
	}

	// Return the new Dense Layer
	return &Dense{
		Neurons:    neurons,
		Activation: activation,
		UsesBias:   bias,
		Name:       name,
		DropOut:    dropout,
	}
}

// ForwardPropagate - Forward propagate the input through the layer
func (d *Dense) Compute() *[]float64 {
	// TBD
	return nil
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
