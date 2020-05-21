package cuslib

import (
	"fmt"
	"math"
	"math/rand"
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

	biasIH				float64
	biasHO				float64

	learningRate		float64
	activationFunc		*Activation // lol super shady
}

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

func NewNN(inputNodes, outputNodes, hiddenNodes int, learningRate float64, activationFunc string) *NeuralNetwork {
	var actFunc *Activation
	if activationFunc == "sigmoid" || activationFunc == "sgd" {
		actFunc = newSigmoid()
	}

	return &NeuralNetwork{
		numInputNodes:  inputNodes,
		numOutputNodes: outputNodes,
		numHiddenNodes: hiddenNodes,
		weightsIH:      NewZeros(inputNodes, hiddenNodes),
		weightsHO:      NewZeros(hiddenNodes, outputNodes),
		biasIH:         rand.Float64(),
		biasHO:         rand.Float64(),
		learningRate:   learningRate,
		activationFunc: actFunc,
	}
}

func (nn *NeuralNetwork) Train(input Matrix, target Matrix) {
	nn.weightsIH.RandomFill()
	nn.weightsHO.RandomFill()
	lengthInout := len(input.val)
	h1, err := MatrixProduct(input, nn.weightsIH)
	if err != nil {
		fmt.Println(err)
		return
	}
	h1.Add(nn.biasIH)
	h1.Map(nn.activationFunc.f)

	o1, err2 := MatrixProduct(h1, nn.weightsHO)
	if err2 != nil {
		fmt.Println(err)
		return
	}
	errors, err3 := MatrixSub(target, o1)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	errors.Show()
	o1.Map(nn.activationFunc.df)
	o1.MulMat(errors)
	o1.Mul(nn.learningRate)

	// calculating delta
	h1T := MatrixTranspose(h1)
	weightDeltaHO, err4 := MatrixProduct(h1T, o1)
	if err4 != nil {
		fmt.Println(err)
		return
	}
	nn.weightsHO.AddMat(weightDeltaHO)

	// hidden layer error
	weightHOT := MatrixTranspose(nn.weightsHO)
	hiddenErrors, err5 := MatrixProduct(weightHOT, errors)
	if err5 != nil {
		fmt.Println(err5)
		return
	}

	h1.Map(nn.activationFunc.df)
	h1.MulMat(hiddenErrors)
	h1.Mul(nn.learningRate)

	inputT := MatrixTranspose(input)
	weightDeltaIH, err6 := MatrixProduct(h1, inputT)
	if err6 != nil {
		fmt.Println(err6)
		return
	}
	nn.weightsIH.AddMat(weightDeltaIH)
}

