package cuslib

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Matrix struct {
	rows	int
	cols	int
	val		[][]int
}

func NewMatrix(arr [][]int) *Matrix {
	return &Matrix{len(arr), len(arr[0]), arr}
}

func (m *Matrix) RandomFill() {
	for i := 0; i < len(m.val); i++ {
		for j := 0; j < len(m.val[0]); j++ {
			m.val[i][j] = rand.Int()
		}
	}
}

func (m *Matrix) AddMat(m2 *Matrix) error {
	if m.rows != m2.rows {
		return fmt.Errorf("number of rows don't match")
	}
	if m.cols != m2.cols {
		return fmt.Errorf("number of cols don't match")
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] += m2.val[i][j]
		}
	}
	return nil
}

func (m *Matrix) MulMat(m2 *Matrix) error {
	if m.rows != m2.rows {
		return fmt.Errorf("number of rows don't match")
	}
	if m.cols != m2.cols {
		return fmt.Errorf("number of cols don't match")
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] *= m2.val[i][j]
		}
	}
	return nil
}

func (m *Matrix) Show() {
	var formatted string
	formatted += "Matrix -> \n"
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			formatted += "\t" + strconv.Itoa(m.val[i][j])
		}
		formatted += "\n"
	}
	fmt.Printf(formatted)
}

func (m *Matrix) Transpose() error {
	if m.cols != m.rows {
		return fmt.Errorf("number of rows (%v) is not equal to number of columns (%v)", m.rows, m.cols)
	}
	tmp := make([][]int, m.cols)
	for i := 0; i < m.cols; i++ {
		tmp[i] = make([]int, m.rows)
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			tmp[i][j] = m.val[j][i]
		}
	}
	m.val = tmp
	return nil
}
