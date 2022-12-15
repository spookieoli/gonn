package gonn

import (
	"github.com/spookieoli/gonn/layers"
	"github.com/spookieoli/gonn/utils"
)

// Gonn is the main struct of the library
type Gonn struct {
	// Layers is a slice of layers
	Layers *[]utils.Layer
}

// Set the Inputs field of the layers
func (g *Gonn) SetInputs(lc *[]utils.LayerConfig) {
	for i := range *lc {
		switch (*lc)[i].LayerType.(type) {
		case *layers.Input:
			continue
		default:
			if (*lc)[i].Inputs == 0 {
				(*lc)[i].Inputs = ((*lc)[i-1]).Neurons
			}
		}
	}
}

// Save a model to a file
func (g *Gonn) Save(path string) error {
	return nil
}

// Load a model from a file
func (g *Gonn) Load(path string) error {
	return nil
}
