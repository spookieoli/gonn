package gonn

import (
	"fmt"

	"github.com/spookieoli/gonn/layers"
	"github.com/spookieoli/gonn/utils"
)

// Gonn is the main struct of the library
type Gonn struct {
	// Layers is a slice of layers
	NN   *utils.NeuralNetwork
	Name string
}

// Set the Inputs field of the layers
func (g *Gonn) SetInputs(lc *[]utils.LayerConfig) {
	for i := range *lc {
		switch (*lc)[i].LayerType.(type) {
		case *layers.Input:
			continue
		// TODO: What in case of a Concat layer?
		default:
			if (*lc)[i].Inputs == 0 {
				if (*lc)[i-1].Bias {
					(*lc)[i].Inputs = ((*lc)[i-1]).Neurons + 1
				} else {
					(*lc)[i].Inputs = ((*lc)[i-1]).Neurons
				}
			}
		}
	}
}

// Check if the first Layer is an Input Layer
func (g *Gonn) CheckInput() bool {
	typeLayer := fmt.Sprintf("%T", (*g.NN.Layers)[0])
	return typeLayer == "*layers.Input"
}

// Save a model to a file
func (g *Gonn) Save(path string) error {
	return nil
}

// Load a model from a file
func (g *Gonn) Load(path string) error {
	return nil
}

// Predict the output of the model
func (g *Gonn) Predict() error {
	return nil
}
