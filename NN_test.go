package nnGo

import (
	"fmt"
	"testing"
)

var nn *NeuralNetwork

func ExampleNewNN () {
	nn = NewNN(2, 10, 1, 1, "sgd", 1000)
	fmt.Println(nn.LearningRate)
	// Output:
	// 1
}

// Training model to learn XOR problem
func ExampleNeuralNetwork_Train () {
	var input [][] float64 = [][] float64 {{1, 1}, {0, 1}, {0, 0}, {1, 0}}
	var target [][] float64 = [][] float64 {{0}, {1}, {0}, {1}}
	nn.Train(input, target)
}

// Predicting Output for XOR problem
func ExampleNeuralNetwork_Predict () {
	nn.Predict([]float64 {1, 0})
}

func BenchmarkNN (b *testing.B) {
	for i := 0; i < b.N; i++ {
		var nn *NeuralNetwork = NewNN(2, 10, 1, 1, "sgd", 1000)
		var input [][] float64 = [][] float64 {{1, 1}, {0, 1}, {0, 0}, {1, 0}}
		var target [][] float64 = [][] float64 {{0}, {1}, {0}, {1}}
		nn.Train(input, target)
		var v [][]float64 = nn.Predict([]float64 {1, 0})
		if v[0][0] < 0.95 {
			b.Error("Expected value greater then 0.95, Got:",v[0][0])
		}
	}
}
