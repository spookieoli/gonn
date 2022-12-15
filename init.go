package gonn

import (
	"github.com/spookieoli/gonn/layers"
	"github.com/spookieoli/gonn/utils"
)

// Create a new Gonn instance
func New(lc *[]utils.LayerConfig) *Gonn {
	nn := &Gonn{}
	// Create the Layers
	for _, layer := range *lc {
		switch layer.LayerType.(type) {
		case *layers.Dense:
			// Create a new Dense layer
			*nn.Layers = append(*nn.Layers, layers.NewDense(layer.Neurons, layer.Activation, layer.Bias, layer.Name, layer.DropOut))
		}
	}
	return nn
}
