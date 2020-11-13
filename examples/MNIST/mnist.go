package main

import (
	"errors"
	"fmt"

	"github.com/moverest/mnist"
	"github.com/sarthakpranesh/nnGo"
)

const (
	Width                 = 28
	Height                = 28
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

	trainDataSet  [][]float64
	targetDataSet [][]float64
)

func main() {
	/*
			-	-	-	mnist	-	-	-
		Make sure you download all four files from the below
		URL and paste them as it is in the examples folder

		mnist dataset - http://yann.lecun.com/exdb/mnist/

		After doing the above you can successfully run: go run mnist.go
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
	nn := nnGo.NewNN(Width*Height, 10000, 10, 0.001, "sgd", 5)
	nn.Train(trainDataSet, targetDataSet)

	// testing
	for i := 0; i < 4; i++ {
		fmt.Println("Actual: ", targetDataSet[i])
		nn.Predict(trainDataSet[i])
	}
}
