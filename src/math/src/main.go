package main

import (
	"fmt"
	"math/src/la4go/Matrix"
)

func main() {
	m1 := Matrix.MatrixI64(10, 10)
	m2 := Matrix.MatrixI64(10, 10)
	//m2.SetElement(5,5,int64(10))
	e := Matrix.Equal(m1, m2)

	fmt.Println(e)
}