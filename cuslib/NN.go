package cuslib

import (
	"fmt"
)

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

	epochs				int // num of epochs to loop
}

func NewNN(inputNodes, hiddenNodes, outputNodes int, learningRate float64, activationFunc string, e int) *NeuralNetwork {
	var actFunc *Activation
	if activationFunc == "sigmoid" || activationFunc == "sgd" {
		actFunc = newSigmoid()
	} else if activationFunc == "tanh" {
		actFunc = newTanh()
	} else {
		fmt.Println("Unsupported activation function or no activation function passed: Defaulting to `SIGMOID`")
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
		epochs:			e,
	}
}

func (n *NeuralNetwork) Train(input [][]float64, t [][]float64) {
	for i := 0; i < n.epochs; i++ {
		var errorSum float64
		fmt.Printf("Epoch %v: ", i+1)
		for i, arr := range input {
			inputMatrix := NewColMatrix(arr)

			// Calculate value of hidden nodes
			hidden, err := MatrixProduct(n.weightsIH, inputMatrix)
			if err != nil {
				fmt.Println(err)
				return
			}
			hidden.Map(n.activationFunc.f)

			// Calculate values of output nodes
			output, err2 := MatrixProduct(n.weightsHO, hidden)
			if err2 != nil {
				fmt.Println("Error: ", err2)
				return
			}
			output.Map(n.activationFunc.f)

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
			gradients := MatrixMap(output, n.activationFunc.df)
			gradients.MulMat(errorOutput)
			gradients.Mul(n.learningRate)

			// calculate delta for weightsHO
			hiddenT := MatrixTranspose(hidden)
			weightsHODelta, err4 := MatrixProduct(gradients, hiddenT)
			if err4 != nil {
				fmt.Println("Error: ", err4)
				return
			}
			n.weightsHO.AddMat(weightsHODelta) // update the weightsHO
			n.biasHO.AddMat(gradients)         // updating the bias

			/*
				Calculate hidden error
				Error = transpose(weightHO)*errorOutput
			*/
			weightsHOT := MatrixTranspose(n.weightsHO)
			errorHidden, err5 := MatrixProduct(weightsHOT, errorOutput)
			if err5 != nil {
				fmt.Println("Error: ", err2)
				return
			}

			// calculate hidden gradients
			hiddenGradient := MatrixMap(hidden, n.activationFunc.df)
			hiddenGradient.MulMat(errorHidden)
			hiddenGradient.Mul(n.learningRate)

			// calculate delta for weightsIH
			inputMatrixT := MatrixTranspose(inputMatrix)
			weightsIHDelta, err6 := MatrixProduct(hiddenGradient, inputMatrixT)
			if err6 != nil {
				fmt.Println("Error: ", err6)
				return
			}
			n.weightsIH.AddMat(weightsIHDelta) // update the weightsIH
			n.biasIH.AddMat(hiddenGradient)    // updating the bias
		}
		fmt.Printf("Error = %v\n", errorSum / float64(len(input)))
	}
}

func (n *NeuralNetwork) Predict(input []float64) [][] float64 {
	inputMatrix := NewColMatrix(input)

	// Calculate value of hidden nodes
	hidden, err := MatrixProduct(n.weightsIH, inputMatrix)
	if err != nil {
		fmt.Println(err)
		return [][]float64 {}
	}
	hidden.Map(n.activationFunc.f)

	// Calculate values of output nodes
	output, err2 := MatrixProduct(n.weightsHO, hidden)
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return [][]float64 {}
	}
	output.Map(n.activationFunc.f)
	//output.Show()
	return output.val
}

func (n *NeuralNetwork) GeneMutate(mapping func(x float64) float64) *NeuralNetwork {
	temp := &NeuralNetwork{
		numInputNodes:  n.numInputNodes,
		numOutputNodes: n.numOutputNodes,
		numHiddenNodes: n.numHiddenNodes,
		weightsIH:      MatrixMap(n.weightsIH, mapping),
		weightsHO:      MatrixMap(n.weightsHO, mapping),
		biasIH:         MatrixMap(n.biasIH, mapping),
		biasHO:         MatrixMap(n.biasHO, mapping),
		learningRate:   n.learningRate,
		activationFunc: n.activationFunc,
		epochs:         n.epochs,
	}
	return temp
}

