package function

import "fmt"

/**
* 闭包，匿名函数修改包内的参数
 */
var data string

func ClosureTest() {
	fmt.Println("---------------ClosureTest-----------------")
	data = "未经过修改的变量"
	my := func() {
		data = "使用匿名函数对变量进行修改"
	}
	fmt.Println(data)
	my()
	fmt.Println(data)
}
