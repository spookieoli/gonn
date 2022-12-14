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
	UseBias bool
}
