package utils

// Interfaces - makes it possible to connect different types of layers
type Layer interface {
	Compute() *[]float64      // Computer the output of the layer
	GetWeights() *[][]float64 // Get the weights of the layer
	GetName() string          // Get the name of the layer
	UseBias() bool            // Check if the layer uses bias
	GetDropOut() float64      // Get the dropout rate of the layer
	GetBias() *[]float64      // Get the bias of the layer - this is normally a 1x1 matrix or one value
	Save(string)              // Save the layer Model
	Load(string)              // Load the layer Model
	// TBC
}
