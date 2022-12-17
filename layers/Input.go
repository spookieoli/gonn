package layers

import "github.com/spookieoli/gonn/utils"

// The Inputlayer is a Placeholder for the Input of the Network
type Input struct {
	Neurons int64 // In Input these are the inputs (data points)
	Name    string
}

// Create new Input Layer and return it
func NewInput(neurons int64) *Input {
	return &Input{
		Neurons: neurons,
	}
}

// ForwardPropagate - Forward propagate the input through the layer
func (d *Input) Compute() {
	// Is a Placeholder
	return
}

// Get all the weights of the layer
func (d *Input) GetWeights() *[][]float64 {
	return &[][]float64{}
}

// Get the Name of the layer
func (d *Input) GetName() string {
	return "InputLayer"
}

// Get the dropout rate of the layer
func (d *Input) GetDropOut() float64 {
	return 0
}

// Get the bias of the layer
func (d *Input) GetBias() *[]float64 {
	return &[]float64{0}
}

// Check if the layer uses bias
func (d *Input) UseBias() bool {
	return false
}

// Get the Outputs of the layer
func (d *Input) GetOutputs() any {
	return nil
}

// Get the number of neurons in the layer
func (d *Input) GetNeuronCount() int64 {
	return d.Neurons
}

// Set the Name of the Layer
func (d *Input) SetName(name string) {
	d.Name = name
}

// This is only a Placeholder - Inputs have no Inputs
func (d *Input) SetInputLayer(l *utils.Layer) {
}

// Placeholder
func (d *Input) GetInputLayername() string {
	return ""
}

// Placeholder
func (d *Input) InitWeights() {

}
