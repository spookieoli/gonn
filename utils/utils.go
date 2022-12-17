package utils

// Interfaces - makes it possible to connect different types of layers
type Layer interface {
	Compute()                 // Computer the output of the layer
	GetWeights() *[][]float64 // Get the weights of the layer
	GetName() string          // Get the name of the layer
	UseBias() bool            // Check if the layer uses bias
	GetDropOut() float64      // Get the dropout rate of the layer
	GetBias() *[]float64      // Get the bias of the layer - this is normally a 1x1 matrix or one value
	GetOutputs() any          // Get the outputs of the layer
	GetNeuronCount() int64    // Get the number of neurons of the layer
	// TBC
}

// LayerConfig is a struct that holds the configuration for a layer
type LayerConfig struct {
	LayerType      Layer
	Neurons        int64
	Inputs         int64 // Inputs is optional - if not set it will be set to the number of neurons of the previous layer
	Activation     func(any) any
	Bias           bool
	Name           string
	DropOut        float64
	InputLayerName string // The Name of the Layer delivering Data to the actual Layer, if empty it will be set to the name of the previous layer
}

// Struct to control the Training of the Model
type TrainConfig struct {
	Epochs   int
	Batch    int
	Threads  int
	Shuffle  bool
	Verbose  bool
	Progress bool
}

// The Model interface is used to save and load models
type Model interface {
	Save(string) error       // Save the Model
	Load(string) error       // Load the Model
	Train(TrainConfig) error // Train the Model (int) is the Number of Epochs, (int) is the Batch Size, (int) is the Number of Threads
	Predict(int) error       // Predict the Output of the Model (int) is number of threads
}

// Payload for the Com channel
type Payload struct {
	Fun func(any) any
	Arg any
}

// Struct for the Neural Network
type NeuralNetwork struct {
	Layers  *[]Layer
	Com     chan Payload
	Threads int
}

// This is the Work Function which will Computer several different Functions if the Payload is sent to the Com channel
func Work(c chan Payload) {
	for u := range c {
		u.Fun(u.Arg)
	}
}
