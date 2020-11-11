package main

import (
	"os"
	"testProject/ReplaceUtils"
)

func main() {
	_, err := os.Stat("model")
	if err != nil{
		dirErr := os.Mkdir("model", os.ModePerm)
		if dirErr == nil{
			//fmt.Println("创建目录失败")
		}
	}
	folder, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ReplaceUtils.GetAllExcel(folder)
}
