package main

import (
	"errors"
	"fmt"
	"github.com/moverest/mnist"
	"github.com/sarthakpranesh/nnGo/cuslib"
)

const (
	Width  = 28
	Height = 28
	TrainingImageFileName = "train-images-idx3-ubyte.gz"
	TrainingLabelFileName = "train-labels-idx1-ubyte.gz"
	TestImageFileName     = "t10k-images-idx3-ubyte.gz"
	TestLabelFileName     = "t10k-labels-idx1-ubyte.gz"
)

var (
	// ErrFormat indicates that the file has not been recognised.
	ErrFormat = errors.New("mnist: invalid format")

	// ErrSize indicates that the labels and images count mismatch.
	ErrSize = errors.New("mnist: size mismatch")

	trainDataSet [][]float64
	targetDataSet [][]float64
)

func main() {
	/*
			-	-	-	mnist	-	-	-
		WORK IN PROGRESS......
	*/

	// load in data
	training, _, err := mnist.Load("")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// convert actual image data to matrix data supported by neural network
	for i := 0; i < 500; i++ {
		imageData, target := training.Get(i)
		var tmp []float64
		var tmp2 []float64
		for _, pixel := range imageData {
			p := float64(pixel) / 255
			tmp = append(tmp, p)
		}
		trainDataSet = append(trainDataSet, tmp)
		for j := 0; j < 10; j++ {
			if j == int(target) {
				tmp2 = append(tmp2, 1)
			} else {
				tmp2 = append(tmp2, 0)
			}
		}
		targetDataSet = append(targetDataSet, tmp2)
	}

	// create and train neural network - HELP NEEDED
	nn := cuslib.NewNN(Width*Height, 10000, 10, 0.1, "tanh", 10)
	nn.Train(trainDataSet, targetDataSet)

	// testing
	for i := 0; i < 10; i++ {
		nn.Predict(trainDataSet[i])
		fmt.Println(targetDataSet[i])
	}
}
