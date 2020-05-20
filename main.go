package main

import (
	"fmt"
	"github.com/sarthakpranesh/percepGo/cuslib"
)

func main() {
	m := cuslib.NewMatrix([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}})
	m.Show()
	err := m.Transpose()
	if err != nil {
		fmt.Println(err)
	}
	m.Show()
}
