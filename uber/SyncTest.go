package uber

import (
	"fmt"
	"study/uber/model/synclock"
	"sync"
)

func SyncTest() {
	syncMethod()
	syncWait()
}

func syncMethod() {
	fmt.Println("---------------同步锁方法，输出发现flag是等于1000-----------------")
	var data synclock.MyData
	for i := 1; i <= 1000000; i++ {
		go data.SyncWrite(i)
	}
	fmt.Println(data)
}
func syncWait() {
	fmt.Println("---------------同步等待方法，输出后发现也是flag=100---------------")
	var data synclock.MyData
	wg := sync.WaitGroup{}
	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go func(i int) {
			data.Write(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(data)
	fmt.Println("---------------同步等待可以控制协程执行结束，但是go 的执行是并发的，可能出现安全问题-----------------------")
	var data2 synclock.MyData
	for i := 1; i <= 1000000; i++ {
		go data2.Write(i)
	}
	fmt.Println(data2)
}
