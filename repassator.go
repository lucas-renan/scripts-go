package main

import (
	"github.com/go-vgo/robotgo"
)

//CodProd codigo dos produtos
var CodProd = []int64{94010123, 45020094, 94010145, 45010021, 45050016, 94010166, 45020095, 45030155, 94010124, 45050172, 45050015, 94010125, 54050139, 54050081, 45030073, 45030103, 45050155}

//ValCred valor de credito e cartao sim
var ValCred = []int64{0}

//ValDebi valor de debito
var ValDebi = []int64{0}
var i = 0
var j = 0

//var z = 0

func main() {

	//fmt.Println(CodProd[i])
	robotgo.Sleep(6)
	robotgo.SetKeyDelay(2000)
	credito()
	//debito()
	//cartaoSiM()
}
func credito() {
	robotgo.SetKeyDelay(10000)
	robotgo.KeyTap("space")
	robotgo.Sleep(4) //
	robotgo.KeyTap("n", "alt")
	robotgo.Sleep(4) //
	for j := 0; j < 9; j++ {
		robotgo.KeyTap("tab")
		robotgo.Sleep(1)
	}
	robotgo.KeyTap("down")
	robotgo.Sleep(1)
	robotgo.KeyTap("P")
	robotgo.KeyTap("A")
	robotgo.KeyTap("R")
	robotgo.KeyTap("T")
	robotgo.Sleep(4) //
	for j := 0; j < 4; j++ {
		robotgo.KeyTap("down")
	}
	robotgo.KeyTap("tab")
	for j := 0; j < 4; j++ {
		robotgo.KeyTap("down")
	}
	for j := 0; j < 15; j++ {
		robotgo.KeyTap("tab")
	}
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
	robotgo.Sleep(1)
	robotgo.KeyTap("tab")
	robotgo.Sleep(1)
	robotgo.KeyTap("down")
	robotgo.Sleep(1)
	//;-----codigo do  procedimento
	//	Send % CodProcd[i]
	robotgo.KeyTap("%d", CodProd[i])
	//;-----------
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
	robotgo.Sleep(1)
	robotgo.KeyTap("enter")
	robotgo.Sleep(1)
	for j := 0; j < 8; j++ {
		robotgo.KeyTap("tab")
		robotgo.Sleep(1)
	}
	robotgo.Sleep(1)
	robotgo.KeyTap("down")
	for j := 0; j < 10; j++ {
		robotgo.KeyTap("BackSpace")
		robotgo.Sleep(1)
	}
	//;-----------valor na opcao cred
	//	Send % Cred[i]
	robotgo.KeyTap("%d", ValCred[i])
	//;----------------
	robotgo.Sleep(1)
	robotgo.KeyTap("tab")
	robotgo.Sleep(1)
	robotgo.KeyTap("s", "alt")
}
func debito() {
	robotgo.KeyTap("tab")
	robotgo.KeyTap("%d", ValDebi[i])
}
func cartaoSiM() {
	robotgo.KeyTap("tab")
}
