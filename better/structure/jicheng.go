package structure

import "fmt"

/**
* 内嵌结构体，实际上就是相当于java的继承
 */
type A struct {
	name string
	age  int
}
type B struct {
	A
	Sex string
}

func fatherAndSunTest() {
	var b B
	b.name = "子类的name"
	b.age = 18
	b.Sex = "man"
	fmt.Println("B继承了A", b)
}
