package Matrix

import "errors"

type matrixINT16 struct {
	m      int
	n      int
	matrix [][]int16
}

func MatrixI16(m int, n int) matrixINT16 {
	matrix := make([][]int16, m, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int16, n)
	}

	return matrixINT16{m, n, matrix}
}

func (this matrixINT16) GetM() int {
	return this.m
}

func (this matrixINT16) GetN() int {
	return this.n
}

func (this matrixINT16) GetElement(m int, n int) (interface{}, error) {
	if m <= this.m && n <= this.n {
		return this.matrix[m][n], nil
	}

	return 0, errors.New("Index out of range")
}

func (this matrixINT16) SetElement(m int, n int, value interface{}) error {
	if m <= this.m && n <= this.n {
		val, ok := value.(int16)
		if ok {
			this.matrix[m][n] = val
			return nil
		} else {
			return errors.New("Value isnt integer")
		}
	}

	return errors.New("Index out of range")
}
