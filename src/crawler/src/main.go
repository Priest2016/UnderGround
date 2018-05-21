package main

import (
	"crawler/src/web"
	"fmt"
)

func main() {
	p := web.Page("http://zernon.ru")
	fmt.Println(p.GetSource())
}
