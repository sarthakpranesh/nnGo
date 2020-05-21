package main

import (
	"github.com/sarthakpranesh/percepGo/cuslib"
)

func main() {
	input := [][]float64 {{1, 1}, {1, 0}, {0, 1}, {0, 0}}
	target := [][]float64 {{0}, {1}, {1}, {0}}
	nn := cuslib.NewNN(2, 2, 1, 1.6, "sgd")
	nn.Train(input, target)
}