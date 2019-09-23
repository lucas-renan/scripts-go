package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

//CodProd contem codigos de produtos
var CodProd = []int64{94010123, 45020094, 94010145, 45010021, 45050016, 94010166, 45020095, 45030155, 94010124, 45050172, 45050015, 94010125, 54050139, 54050081, 45030073, 45030103, 45050155}
var i = 2

//var i = 0

func main() {
	robotgo.Sleep(1)
	//fmt.Println("primeiro elemento", CodProd[i])
	//for i := 0; i < 2; i++ {
	//	fmt.Println("primeiro elemento", CodProd[i])
	//}
	robotgo.KeyTap("%v", CodProd[i])
	//fmt.Println(CodProd
	fmt.Println("%d", CodProd[i])
}
