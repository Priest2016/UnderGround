package Matrix

import "errors"

type matrixINT32 struct {
	m      int
	n      int
	matrix [][]int32
}

func MatrixI32(m int, n int) matrixINT32 {
	matrix := make([][]int32, m, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int32, n)
	}

	return matrixINT32{m, n, matrix}
}

func (this matrixINT32) GetM() int {
	return this.m
}

func (this matrixINT32) GetN() int {
	return this.n
}

func (this matrixINT32) GetElement(m int, n int) (interface{}, error) {
	if m <= this.m && n <= this.n {
		return this.matrix[m][n], nil
	}

	return 0, errors.New("Index out of range")
}

func (this matrixINT32) SetElement(m int, n int, value interface{}) error {
	if m <= this.m && n <= this.n {
		val, ok := value.(int32)
		if ok {
			this.matrix[m][n] = val
			return nil
		} else {
			return errors.New("Value isnt integer")
		}
	}

	return errors.New("Index out of range")
}
