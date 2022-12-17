package gonn

import (
	"fmt"
	"os"
	"strconv"

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
func (g *Gonn) SetInputs() {
	// Create map with LayerNames as keys and the Layer as value
	layerMap := make(map[string]*utils.Layer)
	for i, val := range *g.NN.Layers {
		switch val.(type) {
		case *layers.Dense:
			if val.GetName() == "" {
				val.SetName("Dense" + strconv.Itoa(i))
			}
			layerMap[val.GetName()] = &val
		case *layers.Input:
			if val.GetName() == "" {
				val.SetName("Input" + strconv.Itoa(i))
			}
			layerMap[val.GetName()] = &val
		default:
			os.Exit(1)
		}
	}

	// Set all layers there Incomming Layers
	for i := range *g.NN.Layers {
		switch (*g.NN.Layers)[i].(type) {
		case *layers.Input:
			continue // No Inputs for Inputlayers
		default:
			(*g.NN.Layers)[i].SetInputLayer(layerMap[(*g.NN.Layers)[i].GetInputLayername()])
		}
	}
}

// Create the Weights of the Layers
func (g *Gonn) CreateWeights() {
	for i := range *g.NN.Layers {
		switch (*g.NN.Layers)[i].(type) {
		case *layers.Input:
			continue // No weights for Input Layers
		default:
			(*g.NN.Layers)[i].InitWeights()
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
