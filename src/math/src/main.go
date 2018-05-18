package main

import (
	"math/src/la4go/Matrix"
	"fmt"
)

func main() {
	m:=Matrix.MatrixI(10,10)

	m.SetElement(8,8,19)
	e,err:=m.GetElement(80,8)

	if err!=nil{
		fmt.Println(err.Error())
	} else {
		fmt.Println(e)
	}
}
