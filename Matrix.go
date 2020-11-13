package nnGo

import (
	"fmt"
	"math/rand"
	"strconv"
)

// Matrix is a struct that represents a mathematical Matrix
type Matrix struct {
	Rows	int
	Cols	int
	Val		[][]float64
}

// NewMatrix takes in an 2D array of float64 numbers and returns a Matrix
func NewMatrix(arr [][]float64) Matrix {
	return Matrix{len(arr), len(arr[0]), arr}
}

// NewColMatrix takes in a 1D array of float64 numbers and returns a Matrix (mathematically a vector)
func NewColMatrix(arr []float64) Matrix {
	tmp := make([][]float64, len(arr))
	for i := 0; i < len(arr); i++ {
		tmp[i] = make([]float64, 1)
		tmp[i][0] = arr[i]
	}
	return Matrix{
		Rows: len(arr),
		Cols: 1,
		Val:  tmp,
	}
}

// NewZeros takes in number of rows and columns of type int and returns a Matrix of Val zero
func NewZeros(Rows, Cols int) Matrix {
	tmp := make([][]float64, Rows)
	for i := 0; i < Rows; i++ {
		tmp[i] = make([]float64, Cols)
	}
	return Matrix{
		Rows: Rows,
		Cols: Cols,
		Val:  tmp,
	}
}

// NewRandom takes in number of rows and column of type int and returns a Matrix of randomly filled values
func NewRandom(Rows, Cols int) Matrix {
	m := NewZeros(Rows, Cols)
	m.RandomFill()
	return m
}

// RandomFill fills the Matrix with random values
func (m *Matrix) RandomFill() {
	for i := 0; i < len(m.Val); i++ {
		for j := 0; j < len(m.Val[0]); j++ {
			m.Val[i][j] = rand.NormFloat64()
		}
	}
}

// Add takes in a float64 number and adds it to all elements of the Matrix
func (m *Matrix) Add(a float64) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Val[i][j] += a
		}
	}
}

// Sub takes in a float64 number and substracts it to all elements of the Matrix
func (m *Matrix) Sub(a float64) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Val[i][j] -= a
		}
	}
}

// Mul takes in a float64 number and Multiplies it to all elements of the Matrix
func (m *Matrix) Mul(a float64) {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Val[i][j] *= a
		}
	}
}

// Show is used to visualize Matrix
func (m *Matrix) Show() {
	var formatted string
	formatted += "Matrix -> Dim: {" + strconv.Itoa(m.Rows) + ", " + strconv.Itoa(m.Cols) + "}  \n"
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			formatted += "\t" + strconv.FormatFloat(m.Val[i][j], 'f', 6, 64)
		}
		formatted += "\n"
	}
	fmt.Printf(formatted)
}

// Transpose performs the mathamatical transpose operation on the Matrix
func (m *Matrix) Transpose() {
	r := m.Rows
	c := m.Cols
	tmp := make([][]float64, c)
	for i := 0; i < c; i++ {
		tmp[i] = make([]float64, r)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			tmp[j][i] = m.Val[i][j]
		}
	}
	m.Rows = c
	m.Cols = r
	m.Val = tmp
}

// AddMat takes in another Matrix as argument and performs element wise addition operation
func (m *Matrix) AddMat(m2 Matrix) error {
	if m.Rows != m2.Rows {
		return fmt.Errorf("number of Rows don't match")
	}
	if m.Cols != m2.Cols {
		return fmt.Errorf("number of Cols don't match")
	}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Val[i][j] += m2.Val[i][j]
		}
	}
	return nil
}

// SubMat takes in another Matrix as argument and performs element wise substraction operation
func (m *Matrix) SubMat(m2 Matrix) error {
	if m.Rows != m2.Rows {
		return fmt.Errorf("number of Rows don't match")
	}
	if m.Cols != m2.Cols {
		return fmt.Errorf("number of Cols don't match")
	}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Val[i][j] -= m2.Val[i][j]
		}
	}
	return nil
}

// MulMat takes in another Matrix as argument and performs element wise multiplication operation
func (m *Matrix) MulMat(m2 Matrix) error {
	if m.Rows != m2.Rows {
		return fmt.Errorf("number of Rows don't match")
	}
	if m.Cols != m2.Cols {
		return fmt.Errorf("number of Cols don't match")
	}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Val[i][j] *= m2.Val[i][j]
		}
	}
	return nil
}

// Map takes in a function of type func(x float64) float64 and applies this function to all elements
// of the matrix
func (m *Matrix) Map(mapFunc func(x float64) float64) {
	for i, arr1 := range m.Val {
		for j, ele := range arr1 {
			m.Val[i][j] = mapFunc(ele)
		}
	}
}

// MatrixSum takes in a Matrix and return the sum of all elements of that Matrix
func MatrixSum(m Matrix) float64 {
	var sum float64
	for _, arr1 := range m.Val {
		 for _, ele := range arr1 {
		 	sum += ele
		 }
	}
	return  sum
}

// MatrixAdd takes in two Matrix and returns their element wise addition
func MatrixAdd(m Matrix, m2 Matrix) (Matrix, error) {
	err := m.AddMat(m2)
	if err != nil {
		return Matrix{}, err
	}
	return m, nil
}

// MatrixSub takes in two Matrix and returns their element wise subtraction
func MatrixSub(m Matrix, m2 Matrix) (Matrix, error) {
	err := m.SubMat(m2)
	if err != nil {
		return Matrix{}, err
	}
	return m, nil
}

// MatrixMul takes in two Matrix and returns their element wise Multiplication
func MatrixMul(m Matrix, m2 Matrix) (Matrix, error) {
	err := m.MulMat(m2)
	if err != nil {
		return Matrix{}, err
	}
	return m, nil
}

// MatrixProduct takes in two Matrix and returns their Matrix product (Cross product)
func MatrixProduct(m Matrix, m2 Matrix) (Matrix, error) {
	if m.Cols != m2.Rows {
		return Matrix{}, fmt.Errorf("matrix product not possible")
	}
	tmp := make([][]float64, m.Rows)
	for i := 0; i < m.Rows; i++ {
		tmp[i] = make([]float64, m2.Cols)
	}
	m2.Transpose()
	for i, arr1 := range m.Val {
		for j, arr2 := range m2.Val {
			var sum float64
			for k := 0; k < len(arr2); k++ {
				sum += arr1[k]*arr2[k]
			}
			tmp[i][j] = sum
		}
	}
	return Matrix{
		Rows: m.Rows,
		Cols: m2.Rows, // this will not be m2.Cols because we transposed m2 earlier
		Val:  tmp,
	}, nil
}

// MatrixTranspose takes in a Matrix and returns it's transpose Matrix
func MatrixTranspose(m Matrix) Matrix {
	m.Transpose()
	return m
}

// MatrixMap takes in a Matrix and a function, returns a Matrix created by applying 
// the function to all elements of the Matrix passed in as argument
func MatrixMap(m Matrix, f func(x float64) float64) Matrix {
	m.Map(f)
	return m
}

// MatrixAve takes in a Matrix and returns the average of type float64 of all the elements in that matrix
func MatrixAve(m Matrix) float64 {
	return MatrixSum(m)/float64(m.Rows * m.Cols)
}