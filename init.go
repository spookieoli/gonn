package gonn

import "github.com/spookieoli/gonn/utils"

type Gonn struct {
	// Layers is a slice of layers
	Layers *[]utils.Layer
}

// Create a new Gonn instance
func New(lc *[]utils.LayerConfig) *Gonn {
	// Create the Layers
	for _, layer := range *lc {
		switch layer.LayerType {
			// TBD
	}
	return &Gonn{}
}