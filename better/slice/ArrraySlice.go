package slice

import "fmt"

func ArraySliceTest() {
	sliceTest()
}

func sliceTest() {
	var mans []man
	var mans2 []man
	m0 := man{name: "name0", power: 0}
	m1 := man{name: "name1", power: 1}
	m2 := man{name: "name2", power: 2}
	m3 := man{name: "name3", power: 3}
	m4 := man{name: "name4", power: 4}

	fmt.Println("-------------------添加操作--------------------------------")
	fmt.Println("mans：", mans)
	fmt.Println("mans2：", mans2)
	//添加一个元素
	mans = append(mans, m1)
	fmt.Println("mans：", mans)
	//添加两个元素，注意后面只能是单个元素，所以如果是数组，就加三个点，表示逐个添加
	mans2 = append(mans2, m1, m2, m3)
	mans = append(mans, []man{m2, m3}...)
	fmt.Println("mans：", mans)
	fmt.Println("mans2：", mans2)
	//拼接的第一参数必须是数组，因为拼接扩容是在第一的基础上的
	mans = append([]man{m0}, mans...)
	fmt.Println("mans：", mans)

	fmt.Println("---------------------插入操作------------------------")
	index := 1
	//因为数组作为后面的变量的时候，不能添加单个变量，所以这里采用了中间变量
	//temp := append(mans[:index],m4)
	//fmt.Println("mans:",mans)
	//mans = append(temp,mans[index:]...)
	////采用组合的形式完成插入
	//mans2 = append(append(mans2[:index],m4),mans2[index:]...)
	//这个时候根据输出结果，会发现原来数组有一个直被替换了，所以append(mans[:index],m4)的真正意义是向右移动一位然后添加，原来数组会受影响

	temp := append([]man{m4}, mans[index:]...)
	mans = append(mans[:index], temp...)
	//采用组合的形式完成插入
	mans2 = append(mans2[:index], append([]man{m4}, mans2[index:]...)...)
	fmt.Println("mans：", mans)
	fmt.Println("mans2：", mans2)

	fmt.Println("----------------替换操作---------------------")
	//拼接的结果是[0:index)和单个元素拼接是正确的，但是对于数组本身是下标为index的替换操作了，所以这里不能用mans去接收，采用中间变量mans3接收
	mans2 = append(append(mans2[:index], m0), mans2[index+1:]...)
	mans3 := append(mans[:index], m0)
	fmt.Println("mans：", mans)
	fmt.Println("mans2：", mans2)
	fmt.Println("mans3：", mans3)
	fmt.Println("-----------------------------")
	fmt.Println("删除操作")
	mans2 = append(mans2[:index], mans2[index+1:]...)
	fmt.Println("mans2：", mans2)

	fmt.Println("---------------测试扩容问题------------------")
	aa := make([]string, 5)
	aa[0] = "1"
	aa[2] = "2"
	aa[3] = "3"
	aa[4] = "4"
	aa[1] = "5"
	fmt.Println()
	strings := append(aa, "6")
	fmt.Println(strings)
	//测试结果是aa并没有发生变化，扩容的时候，是会copy生成新的数组，而对原来的数组没有影响
	fmt.Println(aa)

}

type man struct {
	name  string
	power int
}
