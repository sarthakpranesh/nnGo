package nnGo

import (
	"testing"
	"fmt"
)

type TestMethods struct {
	data		[][] float64
	answer		[][] float64
}

var testMethodAddMat [] TestMethods = [] TestMethods {
	TestMethods{
		[][] float64 {{1,1,1},{1,1,1},{1,1,1}},
		[][] float64 {{2,3,4},{2,3,4},{2,3,4}},
	},
	TestMethods{
		[][] float64 {{0,0,0},{0,0,0},{0,0,0}},
		[][] float64 {{1,2,3}, {1,2,3}, {1,2,3}},
	},
	TestMethods{
		[][] float64 {{-1,-1,-1},{-1,-1,-1},{-1,-1,-1}},
		[][] float64 {{0,1,2},{0,1,2},{0,1,2}},
	},
	TestMethods{
		[][] float64 {{1,-1,1},{1,-2,1},{1,-3,1}},
		[][] float64 {{2,1,4},{2,0,4},{2,-1,4}},
	},
}

var testMethodSubMat [] TestMethods = [] TestMethods {
	TestMethods{
		[][] float64 {{1,1,1},{1,1,1},{1,1,1}},
		[][] float64 {{0,1,2},{0,1,2},{0,1,2}},
	},
	TestMethods{
		[][] float64 {{0,0,0},{0,0,0},{0,0,0}},
		[][] float64 {{1,2,3}, {1,2,3}, {1,2,3}},
	},
	TestMethods{
		[][] float64 {{-1,-1,-1},{-1,-1,-1},{-1,-1,-1}},
		[][] float64 {{2,3,4},{2,3,4},{2,3,4}},
	},
	TestMethods{
		[][] float64 {{1,-1,1},{1,-2,1},{1,-3,1}},
		[][] float64 {{0,3,2},{0,4,2},{0,5,2}},
	},
}

var testMethodMulMat [] TestMethods = [] TestMethods {
	TestMethods{
		[][] float64 {{1,1,1},{1,1,1},{1,1,1}},
		[][] float64 {{1,2,3},{1,2,3},{1,2,3}},
	},
	TestMethods{
		[][] float64 {{1,2,3},{1,2,3},{1,2,3}},
		[][] float64 {{1,4,9}, {1,4,9}, {1,4,9}},
	},
}

func compareMatrix (m1, m2 Matrix) bool {
	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			if m1.Val[i][j] != m2.Val[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMatrix (t *testing.T) {
	var fm1 [][] float64 = [][] float64 {{1,2,3}, {1,2,3}, {1,2,3}}

	fmt.Println("Running Table Test: AddMat");
	for _, v := range testMethodAddMat {
		var m1 Matrix = NewMatrix(fm1)
		var m2 Matrix = NewMatrix(v.data)
		m1.AddMat(m2);
		if compareMatrix(m1, m2) {
			t.Error("Expected:", v.answer, "Got:", m1.Val)
		}
	}

	fmt.Println("Running Test: Transpose")
	var m1 Matrix = NewMatrix(fm1)
	var m2 Matrix = NewMatrix([][] float64 {{1,1,1}, {2,2,2},{3,3,3}})
	m1.Transpose()
	if compareMatrix(m1,m2) {
		t.Error("Expected:", m2.Val, "Got:", m1.Val)
	}

	fmt.Println("Running Table Test: SubMat")
	for _, v := range testMethodSubMat {
		var m1 Matrix = NewMatrix(fm1)
		var m2 Matrix = NewMatrix(v.data)
		m1.SubMat(m2);
		if compareMatrix(m1, m2) {
			t.Error("Expected:", v.answer, "Got:", m1.Val)
		}
	}

	fmt.Println("Running Table Test: MulMat")
	for _, v := range testMethodMulMat {
		var m1 Matrix = NewMatrix(fm1)
		var m2 Matrix = NewMatrix(v.data)
		m1.MulMat(m2);
		if compareMatrix(m1, m2) {
			t.Error("Expected:", v.answer, "Got:", m1.Val)
		}
	}
}
