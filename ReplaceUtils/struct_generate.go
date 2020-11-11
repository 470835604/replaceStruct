package ReplaceUtils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

//excel转换变量类型
var types = map[string]string{
	"int":				"int",
	"float":			"float64",
	"array(int)":		"[]int",
	"array(float)":		"[]float64",
	"array2(int)":		"[][]int",
	"string":			"string",
}

//struct拼接所需字符
const (
	packAge = "package model\n\n"
	titleBeFor = "type  "
	titleCenter = "  struct {\n"
	tailBeFor = "}"
)

//读取excel文件
func ReadExcel(filepath,fileName string) () {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	name:= make([]string,0)
	typeData:= make([]string,0)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows(fileName)
	if rows == nil{
		rows,err = f.GetRows("Sheet1")
	}
	for index, row := range rows {
		if index >= 2 || row[index] == ""{
			break
		}
		for _, colCell := range row {
			if colCell == ""{
				break
			}
			if index == 0{
				name = append(name,colCell)
			}else if index == 1{
				typeData = append(typeData,colCell)
			}
		}
	}
	s := splicingStruct(fileName, name, typeData)

	writeErr := ioutil.WriteFile("model/"+fileName+".go", []byte(s), os.ModePerm)
	if writeErr != nil {
		fmt.Println("写入文件失败，读取地址:","")
	}
	cmd := exec.Command("gofmt", "-w", "model/"+fileName+".go")
	_ = cmd.Run()
	fmt.Println(filepath,"文件转换完成")

}

//拼接struct
func splicingStruct(fileName string,name,typeData []string)string  {
	str := packAge+titleBeFor+fileName+titleCenter
	for i := range name {
		str += "   "+capitalize(name[i])+"   "+types[typeData[i]]+"\n"
	}
	str+=tailBeFor
	fmt.Println(str)
	return str
}


//字符首字母大写转换
func capitalize(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

func GetAllExcel(testFilePath string)  {
	dir, err := ioutil.ReadDir(testFilePath)
	if err != nil{
		fmt.Println("该目录地下没有文件",testFilePath)
	}
	for _,fileOs := range dir {
		if fileOs.IsDir() {
			GetAllExcel(testFilePath+"/"+fileOs.Name())
		}else {
			if !strings.HasPrefix(fileOs.Name(), "~$") && strings.HasSuffix(fileOs.Name(),".xlsx"){
				fmt.Println("开始读取:",fileOs.Name())
				ReadExcel(testFilePath+"/"+fileOs.Name(), strings.TrimSuffix(fileOs.Name(), ".xlsx"))
			}
		}
	}
}

