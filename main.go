package main

import (
	"github.com/sarthakpranesh/percepGo/cuslib"
)

func main() {
	m := cuslib.NewMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	m.Show()
	m.Transpose()
	m.Show()
}
