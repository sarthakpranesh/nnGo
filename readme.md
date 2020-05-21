# Toy Neural Network
<p>
    A small library that has it's own Matrix implementation and a 3 layered 
    (input, hidden, output) Neural Network which utilizes this Matrix implementation. This is just a toy library that 
    I created for revisiting my knowledge on the topic of Neural Networks and help me better myself at
    GO programming language.
</p>

## Pre - Requisites
   Make sure you have `GO` Installed.

## Setting It Up
1. git clone https://github.com/sarthakpranesh/nnGo.git
2. cd nnGo
3. That's it hurray! You can use the main.go as a playground to test the library out!!!
 There is also a XOR example that pre-exists in that file to help you out.
 
 ## About library
 The `cuslib` folder in this repo holds two files, Matrix.go and NN.go. Matrix.go contains all required 
 components that should be there to generate and perform operations on Matrices, where as the NN.go file 
 contains code that utilises the Matrix implementation to provide a 3 layered fully connected Neural Network.
 
 <br/>
 
- Neural Network - generators and operations
    - Actual `Struct` - 
        ```go
        type NeuralNetwork struct {
            numInputNodes		int
            numOutputNodes		int
            numHiddenNodes		int
        
            weightsIH			Matrix	// weights from input to hidden layer
            weightsHO			Matrix	// weights from hidden to output layer
        
            biasIH			Matrix
            biasHO			Matrix
        
            learningRate		float64
            activationFunc		*Activation  // Activation function
        
            epochs			int  // num of epochs to loop
        }
        ``` 
    - `NewNN` - takes in six arguments as the following: (number_of_input_nodes int, number_of_hidden_nodes int, 
    number_of_output_nodes int, learning_rate float64, activation_func_name string, epochs int).
    Returns a NeuralNetwork struct. The only activation function supported right now is the `sigmoid` function, 
    which can be selected using the string "sigmoid" or "sgd".
         ```go
        eg:-
        nn := cuslib.NewNN(2, 10, 1, 1, "sgd", 500)
        ``` 
    - `Train` - method on NeuralNetwork struct that takes in two arguments one 2D float64 input and another
     2D float64 target, and attempts to adjusts the 
    internal weights to learn the inputs using 
    target values.
        ```go
        eg:-
        input = [][]float64 {{1, 1}, {0, 1}, {0, 0}, {1, 0}}
        target = [][]float64 {{0}, {1}, {0}, {1}}
        nn.Train(input, target)
        ``` 
    - `Predict` - method on NeuralNetwork struct that takes in 1D float64 input array and predicts the output
        ```go
        eg:-
        nn.Predict([]float64 {1, 1})
        ``` 
 
- Matrix - generators and operations
    - Actual `Struct` - 
        ```go
        type Matrix struct {
            rows	int
            cols	int
            val		[][]float64
        }
        ``` 
    - `NewMatrix` - takes in a 2D array of float64 numbers and returns the generated Matrix
        ```go
        eg:- 
        data := [][]float64 {{1, 1}, {1, 0}, {0, 1}, {0, 0}}
        matrix := cuslib.NewMatrix(data)
        ```
    - `NewColMatrix` - takes in 1D array of float64 numbers and returns the generated column matrix
        ```go
        eg:- 
        data := []float64 {1, 0, 2, 3}
        colMatrix := cuslib.NewColMatrix(data)
        ```
    - `NewZeros` - takes in number of rows and cols of type int, and returns a Matrix of that order 
    having all elements as Zeros.
         ```go
        eg:- 
        zeroMatrix := cuslib.NewZeros(3, 3) // would generate a 3 x 3 zero matrix
        ```
    - `NewRandom` - takes in number of rows and cols of type int, and returns a Matrix of that order 
    having random float numbers as elements
        ```go
        eg:- 
        randMatrix := cuslib.NewRandom(3, 3) // would generate a 3 x 3 random matrix
        ```
    - `Add` - method on Matrix struct, adds a scalar value to all elements of the Matrix
    - `Sub` - method on Matrix struct, subtracts a scalar value from all elements of the Matrix
    - `Mul` - method on Matrix struct, multiplies scalar value to all elements of the Matrix
    - `AddMat` - method on Matrix struct, takes in another Matrix and does element wise addition
    - `SubMat` - method on Matrix struct, takes in another Matrix and does element wise subtraction
    - `MulMat` - method on Matrix struct, takes in another Matrix and does element wise multiplication
    - `Transpose` - method on Matrix struct which transposes the matrix
    - `Map` - method on Matrix struct, takes in a function which operates on float64 and applies it 
    to all elements in the matrix
    - `Show` - method on Matrix struct that displays the matrix in the terminal in a convenient manner
    - `MatrixSum` - function takes in a Matrix and returns the sum of its elements
    - `MatrixAdd` - function takes in two Matrices and returns the element wise added Matrix
    - `MatrixSub` - function takes in two Matrices and returns the difference Matrix of the two Matrices
    - `MatrixMul` - function takes in two Matrices and returns the element wise multiplied Matrix
    - `MatrixProduct` - function takes in two Matrices and returns there Matrix product
    - `MatrixTranspose` - function takes in a Matrix and returns a transposed Matrix
    - `MatrixMap` - function takes in a Matrix and a function that operates on float64, returns a Matrix whose
    elements belong to the passed matrix but are processed with the function passed

## Looking for help
<p>
    Currently I only have one example ( XOR example ) implemented using this library and only one activation function supported.
    I would love if anyone of you can help me by implementing more examples or define different activation functions in 
    this Toy library. Also looking for improving the existing code quality. 
</p>
    