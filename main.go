package main

import (
	"github.com/sarthakpranesh/percepGo/cuslib"
)

func main() {
	nn1 := cuslib.NewNN(2, 1, 3, 0.6, "sgd")
	input := cuslib.NewMatrix([][]float64{{1, 1}, {1, 0}, {0, 1}, {0, 0}})
	target := cuslib.NewMatrix([][]float64{{0}, {1}, {1}, {0}})
	for i := 0; i < 5; i++ {
		nn1.Train(input, target)
	}
}