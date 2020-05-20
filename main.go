package main

import (
	"github.com/sarthakpranesh/percepGo/cuslib"
)

func main() {
	m1 := cuslib.NewMatrix([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
	m1.Show()
	// show working
	m1.Transpose()
	m1.Show()
	// transpose working
	m2 := cuslib.NewMatrix([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}})
	m2.Show()
	m2.AddMat(m1)
	m2.Show()
	// add matrices working
	m3 := cuslib.NewMatrix([][]int{{2, 2, 2, 2}, {2, 2, 2, 2}})
	m3.Show()
	m3.MulMat(m1)
	m3.Show()
	// element multiplication of matrices working
	m4 := cuslib.NewMatrix([][]int{{2, 2, 2, 2}, {2, 2, 2, 2}})
	m4.Show()
	m4.RandomFill()
	m4.Show()
}
