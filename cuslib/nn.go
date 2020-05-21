package cuslib

import (
	"fmt"
	"math"
)

type Activation struct {
	f		func(x float64) float64
	df		func(y float64) float64
}

type NeuralNetwork struct {
	numInputNodes		int
	numOutputNodes		int
	numHiddenNodes		int

	weightsIH			Matrix	// weights from input to hidden layer
	weightsHO			Matrix	// weights from hidden to output layer

	biasIH				Matrix
	biasHO				Matrix

	learningRate		float64
	activationFunc		*Activation // lol super shady
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

func NewNN(inputNodes, hiddenNodes, outputNodes int, learningRate float64, activationFunc string) *NeuralNetwork {
	var actFunc *Activation
	if activationFunc == "sigmoid" || activationFunc == "sgd" {
		actFunc = newSigmoid()
	}

	return &NeuralNetwork{
		numInputNodes:  inputNodes,
		numOutputNodes: outputNodes,
		numHiddenNodes: hiddenNodes,
		weightsIH:      NewRandom(hiddenNodes, inputNodes),
		weightsHO:      NewRandom(outputNodes, hiddenNodes),
		biasIH:         NewRandom(hiddenNodes, 1),
		biasHO:         NewRandom(outputNodes, 1),
		learningRate:   learningRate,
		activationFunc: actFunc,
	}
}

func (n *NeuralNetwork) FeedForward(input []float64) Matrix{
	inputMatrix := NewColMatrix(input)

	// Calculate value of hidden nodes
	hidden, err := MatrixProduct(n.weightsIH, inputMatrix)
	if err != nil {
		fmt.Println(err)
		return Matrix{}
	}
	hidden.Map(n.activationFunc.f)

	// Calculate values of output nodes
	output, err2 := MatrixProduct(n.weightsHO, hidden)
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return Matrix{}
	}
	output.Map(n.activationFunc.f)
	return output
}

func (n *NeuralNetwork) BackPropagation(input []float64) {

}

func (n *NeuralNetwork) Train(input [][]float64, t [][]float64) {
	for i, arr := range input {
		output := n.FeedForward(arr)

		/*
			Calculate the output error
			Error = target - output
		*/
		target := NewColMatrix(t[i])
		errorOutput, err := MatrixSub(target, output)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		/*
			Calculate hidden error
			Error = transpose(weightHO)*errorOutput
		*/
		weightsHOT := MatrixTranspose(n.weightsHO)
		errorHidden, err2 := MatrixProduct(weightsHOT, errorOutput)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			return
		}
		errorHidden.Show()
	}
}

