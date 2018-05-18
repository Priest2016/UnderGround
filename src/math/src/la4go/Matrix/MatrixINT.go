package Matrix

import "errors"

type matrixINT struct {
	m int
	n int
	matrix [][]int
}

func MatrixI(m int, n int) matrixINT {
	matrix:=make([][]int, m, n)

	for i:=0;i<m;i++{
		matrix[i]=make([]int, n)
	}

	return matrixINT{m, n, matrix}
}

func (this *matrixINT) GetElement(m int, n int) (int, error) {
	if m<=this.m && n<=this.n {
		return this.matrix[m][n], nil
	}

	return 0, errors.New("Index out of range")
}

func (this *matrixINT) SetElement(m int, n int, value int) error{
	if m<=this.m && n<=this.n {
		this.matrix[m][n]=value
		return nil
	}

	return errors.New("Index out of range")
}