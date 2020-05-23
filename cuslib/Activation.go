package cuslib

import "math"

type Activation struct {
	f		func(x float64) float64
	df		func(y float64) float64
}

/*
	Different activation functions
*/
func newSigmoid() *Activation {
	return &Activation{
		f: func(x float64) float64 {
			return 1 / ( 1 + math.Exp(-x) )
		},
		df: func(y float64) float64 {
			return y * (1 - y)
		},
	}
}

func newTanh() *Activation {
	return &Activation{
		f: func(x float64) float64 {
			return math.Tanh(x)
		},
		df: func(y float64) float64 {
			return 1 - (y*y)
		},
	}
}
