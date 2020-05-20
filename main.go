package main

import (
	"github.com/sarthakpranesh/percepGo/cuslib"
)

func main() {
	m1 := cuslib.NewMatrix([][]int{{2, 2}, {2, 2}, {2, 2}, {2, 2}})
	m1.Show()
	m2 := cuslib.NewMatrix([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}})
	m2.Show()
	m1.SubMat(m2)
	m1.Show()
}
