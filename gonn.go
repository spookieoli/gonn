package gonn

import (
	"github.com/spookieoli/gonn/utils"
)

// Gonn is the main struct of the library
type Gonn struct {
	// Layers is a slice of layers
	Layers *[]utils.Layer
}

// Save a model to a file
func (g *Gonn) Save(path string) error {
	return nil
}

// Load a model from a file
func (g *Gonn) Load(path string) error {
	return nil
}
