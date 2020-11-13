package nnGo

import (
	"fmt"
)

// Model is a simple interface having two methods Train and Predict
type Model interface {
	Train(input [][]float64, t [][]float64)
	Predict(input []float64)
}

// NeuralNetwork is a struct that contains all the values necessary to train a Neural Network model
type NeuralNetwork struct {
	NumInputNodes		int				// Num of input nodes on the Neural Network
	NumOutputNodes		int				// Num of output nodes on the Neural Network
	NumHiddenNodes		int				// Num of hidden nodes on the Neural Network

	WeightsIH			Matrix			// weights from input to hidden layer
	WeightsHO			Matrix			// weights from hidden to output layer

	BiasIH				Matrix
	BiasHO				Matrix

	LearningRate		float64			// Rate at which the Network would learn/fit data
	ActivationFunc		*Activation		// Pointer to an Activation function	

	Epochs				int 			// num of Epochs to loop
}

// NewNN creates and returns a pointers to a NeuralNetwork
func NewNN(inputNodes, hiddenNodes, outputNodes int, LearningRate float64, ActivationFunc string, e int) *NeuralNetwork {
	var actFunc *Activation
	switch ActivationFunc {
	case "sigmoid":
		actFunc = NewSigmoid()
	case "sgd":
		actFunc = NewSigmoid()
	case "tanh":
		actFunc = NewTanh()
	default:
		fmt.Println("Unsupported activation function or no activation function passed: Defaulting to `SIGMOID`")
		actFunc = NewSigmoid()
	}

	return &NeuralNetwork{
		NumInputNodes:  inputNodes,
		NumOutputNodes: outputNodes,
		NumHiddenNodes: hiddenNodes,
		WeightsIH:      NewRandom(hiddenNodes, inputNodes),
		WeightsHO:      NewRandom(outputNodes, hiddenNodes),
		BiasIH:         NewRandom(hiddenNodes, 1),
		BiasHO:         NewRandom(outputNodes, 1),
		LearningRate:   LearningRate,
		ActivationFunc: actFunc,
		Epochs:			e,
	}
}

// Train takes in two 2D arrays of type float64 and trains the 
// NeuralNetwork to fit the data. One arguments should
// be the input data to the Neural Network and another one should be
// One Hot encoding of the actual category of the correspoding input data.
func (n *NeuralNetwork) Train(input [][]float64, t [][]float64) {
	for i := 0; i < n.Epochs; i++ {
		var errorSum float64
		fmt.Printf("Epoch %v: ", i+1)
		for i, arr := range input {
			inputMatrix := NewColMatrix(arr)

			// Calculate value of hidden nodes
			hidden, err := MatrixProduct(n.WeightsIH, inputMatrix)
			if err != nil {
				fmt.Println(err)
				return
			}
			hidden.Map(n.ActivationFunc.F)

			// Calculate values of output nodes
			output, err2 := MatrixProduct(n.WeightsHO, hidden)
			if err2 != nil {
				fmt.Println("Error: ", err2)
				return
			}
			output.Map(n.ActivationFunc.F)

			/*
				Calculate the output error
				Error = target - output
			*/
			target := NewColMatrix(t[i])
			errorOutput, err3 := MatrixSub(target, output)
			if err3 != nil {
				fmt.Println("Error: ", err)
				return
			}

			// storing errorOutput
			errorSum += MatrixAve(errorOutput)

			// We have errorOutput need to calculate gradient
			// calculate grads for output
			gradients := MatrixMap(output, n.ActivationFunc.DF)
			gradients.MulMat(errorOutput)
			gradients.Mul(n.LearningRate)

			// calculate delta for WeightsHO
			hiddenT := MatrixTranspose(hidden)
			WeightsHODelta, err4 := MatrixProduct(gradients, hiddenT)
			if err4 != nil {
				fmt.Println("Error: ", err4)
				return
			}
			n.WeightsHO.AddMat(WeightsHODelta) // update the WeightsHO
			n.BiasHO.AddMat(gradients)         // updating the bias

			/*
				Calculate hidden error
				Error = transpose(weightHO)*errorOutput
			*/
			WeightsHOT := MatrixTranspose(n.WeightsHO)
			errorHidden, err5 := MatrixProduct(WeightsHOT, errorOutput)
			if err5 != nil {
				fmt.Println("Error: ", err2)
				return
			}

			// calculate hidden gradients
			hiddenGradient := MatrixMap(hidden, n.ActivationFunc.DF)
			hiddenGradient.MulMat(errorHidden)
			hiddenGradient.Mul(n.LearningRate)

			// calculate delta for WeightsIH
			inputMatrixT := MatrixTranspose(inputMatrix)
			WeightsIHDelta, err6 := MatrixProduct(hiddenGradient, inputMatrixT)
			if err6 != nil {
				fmt.Println("Error: ", err6)
				return
			}
			n.WeightsIH.AddMat(WeightsIHDelta) // update the WeightsIH
			n.BiasIH.AddMat(hiddenGradient)    // updating the bias
		}
		fmt.Printf("Error = %v\n", errorSum / float64(len(input)))
	}
}

// Predict takes 1D array of input data and returns a One-Hot encoded 2D array
func (n *NeuralNetwork) Predict(input []float64) [][] float64 {
	inputMatrix := NewColMatrix(input)

	// Calculate value of hidden nodes
	hidden, err := MatrixProduct(n.WeightsIH, inputMatrix)
	if err != nil {
		fmt.Println(err)
		return [][]float64 {}
	}
	hidden.Map(n.ActivationFunc.F)

	// Calculate values of output nodes
	output, err2 := MatrixProduct(n.WeightsHO, hidden)
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return [][]float64 {}
	}
	output.Map(n.ActivationFunc.F)
	output.Show()
	return output.Val
}

// Mutate takes in a NeuralNetwork and a function of type func (x float64) float64, and returns
// the mutation of the NeuralNetwork by applying the function to all Matrix in the NeuralNetwork
func Mutate(n NeuralNetwork, mapping func(x float64) float64) NeuralNetwork {
	n.WeightsIH.Map(mapping)
	n.WeightsHO.Map(mapping)
	n.BiasIH.Map(mapping)
	n.BiasHO.Map(mapping)
	return n
}

