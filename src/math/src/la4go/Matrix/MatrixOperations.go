package Matrix

import "fmt"

func Equal(m1 Matrix, m2 Matrix) bool {
	returnedValue := true

	if m1.GetM() == m2.GetM() && m1.GetN() == m2.GetN() {
		for i:=0;i<m1.GetM();i++{
			for j:=0;j<m1.GetN();j++{
				e1,err:=m1.GetElement(i,j)
				if err!=nil{
					fmt.Println(err.Error())
				}

				e2,err:=m2.GetElement(i,j)
				if err!=nil{
					fmt.Println(err.Error())
				}

				if e1!=e2{
					returnedValue = false
				}
			}
		}
	}

	return returnedValue
}
