package main

import (
	"github.com/sarthakpranesh/percepGo/cuslib"
)

func main() {
	nn1 := cuslib.NewNN(2, 1, 3, 1.6, "sgd")
	input := cuslib.NewMatrix([][]float64{{1, 1}, {1, 0}, {0, 1}, {0, 0}})
	input.Show()
	nn1.Train(input)
}