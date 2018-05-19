package Matrix

import "errors"

type matrixINT8 struct {
	m      int
	n      int
	matrix [][]int8
}

func MatrixI8(m int, n int) matrixINT8 {
	matrix := make([][]int8, m, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int8, n)
	}

	return matrixINT8{m, n, matrix}
}

func (this matrixINT8) GetM() int {
	return this.m
}

func (this matrixINT8) GetN() int {
	return this.n
}

func (this matrixINT8) GetElement(m int, n int) (interface{}, error) {
	if m <= this.m && n <= this.n {
		return this.matrix[m][n], nil
	}

	return 0, errors.New("Index out of range")
}

func (this matrixINT8) SetElement(m int, n int, value interface{}) error {
	if m <= this.m && n <= this.n {
		val, ok := value.(int8)
		if ok {
			this.matrix[m][n] = val
			return nil
		} else {
			return errors.New("Value isnt integer")
		}
	}

	return errors.New("Index out of range")
}

