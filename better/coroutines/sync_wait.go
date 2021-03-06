package coroutines

import (
	"fmt"
	"sync"
)

/**
* 协程的同步能够解决共享资源的安全修改，但是不能用for循环不断检索count来判断协程是否执行结束，因为for循环会导致程序崩了
* 所以这个时候采用协程等待，
* 经过多次测试，发现线程等待具有线程同步安全的效果，但是实际上没有，runtime.Gosched()让当前goroutine暂停可以让效果更明显，go build -race检查共享资源的问题
* 但是wait后面的代码仅仅是指在所有的done执行结束后执行，测试发现打印10可能比wait之后的输出还要慢。
 */
func GoSyncWait() {
	var count int = 0
	wg := &sync.WaitGroup{}
	for i := 0; i <= 20; i++ {
		fmt.Println("第", i)
		wg.Add(1)
		go deleteData(&count, wg)
	}
	wg.Wait()
	fmt.Println("所有的协程执行完毕")
}

func deleteData(data *int, wg *sync.WaitGroup) {
	//注意这里应该在最后执行，否则会出现了上面说的打印10可能比wait之后的输出还要慢这个问题，wg是安全的，但是go里面对共享资源的操作不是安全的
	defer wg.Done()
	if *data < 10 {
		*data++
		fmt.Println(*data)
	}
}
