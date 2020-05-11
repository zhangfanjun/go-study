package uber

import (
	"fmt"
	"study/uber/service"
	"study/uber/service/impl"
)

func InterfaceTest() {
	//测试地址调用和值调用的区别
	privateTest()

	//测试接口接收的问题
	interfaceTest()
}

func privateTest() {
	fmt.Println("测试实例调用方法")
	w1 := impl.WeChat{Dollar: 123456}
	w2 := &impl.WeChat{Dollar: 123456}
	a1 := impl.Alipay{Dollar: 123456}
	a2 := &impl.Alipay{Dollar: 123456}
	w1.Pay(456)
	w2.Pay(456)
	a1.Pay(456)
	a2.Pay(456)
	fmt.Printf("w1:%.2f,w2:%.2f,a1:%.2f,a2:%.2f\n", w1.Dollar, w2.Dollar, a1.Dollar, a2.Dollar)
	fmt.Println("从上面的输出可以看出，对于普通方法的调用，只能用地址进行传递，与入参类型无关，否则传递给虚拟入参没有意义")

	myMap := map[int]impl.WeChat{1: {Dollar: 123}}
	myMap[1].Pay(3)
	myMap[1].Show()

	myMap2 := map[int]*impl.WeChat{1: {Dollar: 123}}
	myMap2[1].Pay(3)
	myMap2[1].Show()
	fmt.Println("无论是实例还是地址，都能调用所有私有方法，与私有方法的调用类型无关")
}

func interfaceTest() {
	fmt.Println("测试接口接收实例")
	var pay service.PayInterface
	w1 := impl.WeChat{Dollar: 123456}
	w2 := &impl.WeChat{Dollar: 123456}
	//a1 := impl.Alipay{Dollar:123456}
	a2 := &impl.Alipay{Dollar: 123456}
	pay = w1
	fmt.Println(pay)
	pay = w2
	fmt.Println(pay)
	//pay = a1
	pay = a2
	fmt.Println(pay)
	fmt.Println("从上面的例子中，Alipay是地址调用实现接口，所以只能通过地址传递给接口")
}
