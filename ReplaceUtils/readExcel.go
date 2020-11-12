package ReplaceUtils

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

/**
	读取excel文件
 */
func ReadExcel(filePath,fileName string)map[int]map[string]string {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	tableData := make(map[int]map[string]string)

	variable := make([]string,0)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows(fileName)
	for index, row := range rows {
		if row[0] == ""{
			break
		}

		values := make(map[string]string)
		for i, colCell := range row {
			if colCell == ""{
				break
			}
			if index < 2{
				variable = append(variable,colCell)
			}else if index > 2{
				values[variable[i]] = colCell
			}
		}
		if index > 2{
			tableData[index-3] = values
		}

	}
	return tableData
}

/**
	json将string转换基本数据类型格式
*/
func JsonAnalysis(dataType interface{},data string) interface{} {
	err := json.Unmarshal([]byte(data), &dataType)
	if err != nil{
		panic(err)
	}
	return dataType
}