package structure

import "fmt"

/**
* 匿名结构体，但是会发现实例的时候太复杂了
* 如果字段的结构体必须配合外部结构体使用或者说通用性很低的时候可以使用
 */
type Car struct {
	Wheel  struct{ Size int }
	Engine struct{ Power int }
}

func NiMingStructTest() {
	car := Car{Wheel: struct{ Size int }{Size: 256}, Engine: struct{ Power int }{Power: 800}}
	fmt.Println(car)
}
