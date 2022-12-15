package utils

// Interfaces - makes it possible to connect different types of layers
type Layer interface {
	Compute() *[]float64      // Computer the output of the layer
	GetWeights() *[][]float64 // Get the weights of the layer
	GetName() string          // Get the name of the layer
	UseBias() bool            // Check if the layer uses bias
	GetDropOut() float64      // Get the dropout rate of the layer
	GetBias() *[]float64      // Get the bias of the layer - this is normally a 1x1 matrix or one value
	// TBC
}

// LayerConfig is a struct that holds the configuration for a layer
type LayerConfig struct {
	LayerType  Layer
	Neurons    int64
	Activation func(any) any
	Bias       bool
	Name       string
	DropOut    float64
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
