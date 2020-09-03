package nnGo

import "math"

// Activation is a simple struct that holds Activation functions for forward progations (F) 
// and back propagation (DF)
type Activation struct {
	F		func(x float64) float64
	DF		func(y float64) float64
}

// NewSigmoid returns a pointer to an Activation, set to use Sigmoid
func NewSigmoid() *Activation {
	return &Activation{
		F: func(x float64) float64 {
			return 1 / ( 1 + math.Exp(-x) )
		},
		DF: func(y float64) float64 {
			return y * (1 - y)
		},
	}
}

// NewTanh returns a pointer to an Activation, set to use Tanh
func NewTanh() *Activation {
	return &Activation{
		F: func(x float64) float64 {
			return math.Tanh(x)
		},
		DF: func(y float64) float64 {
			return 1 - (y*y)
		},
	}
}
