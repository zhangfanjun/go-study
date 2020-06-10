package coroutines

import (
	"fmt"
	"sync"
)

/**
* go的协程入参可以是一个公共的变量，但是这个时候涉及到了一个并发的安全问题
* 所以这个时候，每个协程在修改共享资源的时候，需要添加一个同步锁（互斥锁）
 */
func GoSyncLock() {
	var count int = 0
	lock := &sync.Mutex{}
	for i := 0; i < 100; i++ {
		go addData(&count, lock)
	}
}

func addData(data *int, lock *sync.Mutex) {
	lock.Lock()
	if (*data) < 10 {
		*data++
		fmt.Println(*data)
	}
	lock.Unlock()
}
