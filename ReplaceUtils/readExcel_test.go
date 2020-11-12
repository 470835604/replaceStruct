package ReplaceUtils

import (
	"fmt"
	"testing"
)

func TestS(t *testing.T) {
	//types := Types
		//fmt.Println(types)
	//const filePath = "C:/Users/rust/Desktop/ErectsDegrade.xlsx"
	//const fileName = "ErectsDegrade"
	//readExcel := ReadExcel(filePath, fileName)
	//fmt.Println(readExcel)
	s := "[[1.23,23],[2.123],[3]]"
	var array [][]float64
	fmt.Println(JsonAnalysis(array,s))
}

