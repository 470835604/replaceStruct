
<h1>介绍</h1>
<p>replaceStruct是一个将excel表一键生成为Golang Struct的应用程序。
   内置excel表读取、简单json数据转换函数</p>

<h1>参数列表</h1>
<ul>
    <li>int</li>
    <li>float64</li>
    <li>[]int</li>
    <li>[]float64</li>
    <li>[][]int</li>
    <li>string</li>
</ul>
<h1>使用方法</h1>
     <p>去获取<a href="github.com/Byfengfeng/replaceStruct">github.com/Byfengfeng/replaceStruct</a></p>
<h1>使用</h1>

```
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

```


<h1>生成结果示例</h1>


```
package model

type LotteryStageReward struct {

	ID        int

	LotteryId int

	Num       int

	OmId      []int

	Count     []int

	ItemType  []int

	Level     []int

	Flashtime int

}
```

