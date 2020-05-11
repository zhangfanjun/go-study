package structure

type Person struct {
	name string
	age  int
}

/*构造方法其实是一个公共方法，这个方法初始化某个实例然后进行返回*/
func GetZhangSan() Person {
	return Person{name: "张三"}
}
