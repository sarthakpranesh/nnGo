package main

import "github.com/sarthakpranesh/percepGo/cuslib"

var (
	input = [][]float64 {{1, 1}, {0, 1}, {0, 0}, {1, 0}}
	target = [][]float64 {{0}, {1}, {0}, {1}}
)

func main() {
	/*
			-	-	-	XOR Problem	-	-	-
		Below is an example of a simple XOR problem solved using
		cuslib Neural Network Library and Matrix Library.

		Spoiler: Both the libraries are only for toying around. I
		built them just for fun and learning GO programming. Feel
		free to play around. Some important useful notes below

		The Neural Network is essentially a three fully connected
		layer network ( did not have the will to sit and code for
		more dynamically given layer numbers ). You can read more
		about it in "./cuslib/NN.go"
	*/
	nn := cuslib.NewNN(2, 10, 1, 1, "sgd", 500)
	nn.Train(input, target)
	nn.Predict([]float64 {1, 1})
	nn.Predict([]float64 {0, 1})
	nn.Predict([]float64 {1, 0})
	nn.Predict([]float64 {0, 0})
}
