package Matrix

import "errors"

type matrixINT64 struct {
	m      int
	n      int
	matrix [][]int64
}

func MatrixI64(m int, n int) matrixINT64 {
	matrix := make([][]int64, m, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int64, n)
	}

	return matrixINT64{m, n, matrix}
}

func (this matrixINT64) GetM() int {
	return this.m
}

func (this matrixINT64) GetN() int {
	return this.n
}

func (this matrixINT64) GetElement(m int, n int) (interface{}, error) {
	if m <= this.m && n <= this.n {
		return this.matrix[m][n], nil
	}

	return 0, errors.New("Index out of range")
}

func (this matrixINT64) SetElement(m int, n int, value interface{}) error {
	if m <= this.m && n <= this.n {
		val, ok := value.(int64)
		if ok {
			this.matrix[m][n] = val
			return nil
		} else {
			return errors.New("Value isnt integer")
		}
	}

	return errors.New("Index out of range")
}
