package layers

// The Inputlayer is a Placeholder for the Input of the Network
type Input struct {
	Neurons int64 // In Input these are the inputs (data points)
}

// Create new Input Layer and return it
func NewInput(neurons int64) *Input {
	return &Input{
		Neurons: neurons,
	}
}

// ForwardPropagate - Forward propagate the input through the layer
func (d *Input) Compute() *[]float64 {
	// TBD
	return nil
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
