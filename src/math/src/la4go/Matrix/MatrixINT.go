package Matrix

import "errors"

type matrixINT struct {
	m      int
	n      int
	matrix [][]int
}

func MatrixI(m int, n int) matrixINT {
	matrix := make([][]int, m, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	return matrixINT{m, n, matrix}
}

func (this matrixINT) GetM() int {
	return this.m
}

func (this matrixINT) GetN() int {
	return this.n
}

func (this matrixINT) GetElement(m int, n int) (interface{}, error) {
	if m <= this.m && n <= this.n {
		return this.matrix[m][n], nil
	}

	return 0, errors.New("Index out of range")
}

func (this matrixINT) SetElement(m int, n int, value interface{}) error {
	if m <= this.m && n <= this.n {
		val, ok := value.(int)
		if ok {
			this.matrix[m][n] = val
			return nil
		} else {
			return errors.New("Value isnt integer")
		}
	}

	return errors.New("Index out of range")
}
