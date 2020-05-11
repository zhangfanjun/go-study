package function

import "fmt"

/**
* 接口的实现包括两种，
* 一种是结构体的方法实现了接口的方法
* 另一种是匿名一个方法实现接口的方法
 */
func ServiceImplTest() {
	impl := ServiceImpl{data: "5566"}
	interfaceTest(&impl)
	my := myFunc(func() {
		fmt.Println("这是一个匿名函数，myFunc函数实现了接口")
	})
	interfaceTest(my)
}

/*入参为接口的函数*/
func interfaceTest(s Service) {
	s.Handle()
}

/*接口*/
type Service interface {
	Handle()
}

/*结构体*/
type ServiceImpl struct {
	data string
}

/*实现接口的函数*/
func (s *ServiceImpl) Handle() {
	fmt.Println(s.data)
}

/*定义一个func类型实现接口，func的结构和接口的函数结构一致*/
type myFunc func()

func (m myFunc) Handle() {
	m()
}
