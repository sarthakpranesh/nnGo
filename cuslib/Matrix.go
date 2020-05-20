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
			m.val[i][j] = rand.Intn(10)
		}
	}
}

func (m *Matrix) Add(a int) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] += a
		}
	}
}

func (m *Matrix) Sub(a int) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] -= a
		}
	}
}

func (m *Matrix) Mul(a int) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] *= a
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

func (m *Matrix) SubMat(m2 *Matrix) error {
	if m.rows != m2.rows {
		return fmt.Errorf("number of rows don't match")
	}
	if m.cols != m2.cols {
		return fmt.Errorf("number of cols don't match")
	}
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] -= m2.val[i][j]
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
	formatted += "Matrix -> Dim: {" + strconv.Itoa(m.rows) + ", " + strconv.Itoa(m.cols) + "}  \n"
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			formatted += "\t" + strconv.Itoa(m.val[i][j])
		}
		formatted += "\n"
	}
	fmt.Printf(formatted)
}

func (m *Matrix) Transpose() {
	r := m.rows
	c := m.cols
	tmp := make([][]int, c)
	for i := 0; i < c; i++ {
		tmp[i] = make([]int, r)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			tmp[j][i] = m.val[i][j]
		}
	}
	m.rows = c
	m.cols = r
	m.val = tmp
}
