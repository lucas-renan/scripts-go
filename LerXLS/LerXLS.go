package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	const endereco = `C:\Users\renan.m\desktop\teste.xls`
	excelFileName := endereco
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Println("nao consigo abrir")
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Println(text)
			}
		}
	}
}
