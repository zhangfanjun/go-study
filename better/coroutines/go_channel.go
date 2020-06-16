package coroutines

import (
	"fmt"
	"strconv"
	"time"
)

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

/**
* 遍历消费通道
* 注意，遍历的时候，属于无线等待，所以会导致通道阻塞，需要跳出for循环
 */
func GoChannelRange() {
	var ch = make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- strconv.Itoa(i)
		}
	}()
	go func() {
		ch <- "ok"
	}()
	for data := range ch {
		fmt.Println("data:", data)
		if "ok" == data {
			fmt.Println("消费到ok")
			break
		}
	}
}

/**
* 利用无缓存通道两端相互等待的特性，实现生产者和消费者
 */
func Factory() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
}
func producer(c chan int) {
	fmt.Println("生产")
	fmt.Println("产出")
	c <- 1
	for {
		time.Sleep(5 * time.Second)
		data := <-c
		if data == 0 {
			fmt.Println("生产")
			fmt.Println("产出")
			c <- 1
		} else {
			c <- data
		}
	}

}
func consumer(c chan int) {
	for {
		time.Sleep(3 * time.Second)
		data := <-c
		if data == 1 {
			fmt.Println("消费")
			c <- 0
		} else {
			c <- data
		}
	}
}
