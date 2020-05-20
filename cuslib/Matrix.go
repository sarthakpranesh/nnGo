package cuslib

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Matrix struct {
	rows	int
	cols	int
	val		[][]float64
}

func NewMatrix(arr [][]float64) Matrix {
	return Matrix{len(arr), len(arr[0]), arr}
}

func NewZeros(rows, cols int) Matrix {
	tmp := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		tmp[i] = make([]float64, cols)
	}
	return Matrix{
		rows: rows,
		cols: cols,
		val:  tmp,
	}
}

func (m *Matrix) RandomFill() {
	for i := 0; i < len(m.val); i++ {
		for j := 0; j < len(m.val[0]); j++ {
			m.val[i][j] = rand.Float64()
		}
	}
}

func (m *Matrix) Add(a float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] += a
		}
	}
}

func (m *Matrix) Sub(a float64) {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.val[i][j] -= a
		}
	}
}

func (m *Matrix) Mul(a float64) {
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
			formatted += "\t" + strconv.FormatFloat(m.val[i][j], 'f', 6, 64)
		}
		formatted += "\n"
	}
	fmt.Printf(formatted)
}

func (m *Matrix) Transpose() {
	r := m.rows
	c := m.cols
	tmp := make([][]float64, c)
	for i := 0; i < c; i++ {
		tmp[i] = make([]float64, r)
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

func (m *Matrix) Map(mapFunc func(x float64) float64) {
	for i, arr1 := range m.val {
		for j, ele := range arr1 {
			m.val[i][j] = mapFunc(ele)
		}
	}
}

func MatrixProduct(m Matrix, m2 Matrix) (*Matrix, error) {
	if m.cols != m2.rows {
		return &Matrix{}, fmt.Errorf("matrix product not possible")
	}
	tmp := make([][]float64, m.rows)
	for i := 0; i < m.rows; i++ {
		tmp[i] = make([]float64, m2.cols)
	}
	m2.Transpose()
	for i, arr1 := range m.val {
		for j, arr2 := range m2.val {
			var sum float64
			for k := 0; k < len(arr2); k++ {
				sum += arr1[k]*arr2[k]
			}
			tmp[i][j] = sum
		}
	}
	return &Matrix{
		rows: m.rows,
		cols: m2.rows, // this will not be m2.cols because we transposed m2 earlier
		val:  tmp,
	}, nil
}