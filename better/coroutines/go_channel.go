package coroutines

import "fmt"

/**
* 使用通道做并发
* 通道具有每次传递一个数据元，没有发送或者没有接收都会导致阻塞，根据这个特性可以设计协程的等待
 */
func GoChannelWait() {
	//声明一个通道
	var ch = make(chan int)
	go func() {
		fmt.Println("协程执行TODO")
		fmt.Println("协程执行结束，向通道发送数据666")
		ch <- 666
	}()
	fmt.Println("开启了go之后，开始消费通道，等待通道生产数据")
	<-ch
	fmt.Println("消费成功")
}
