package coroutines

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/**
* 对于协程的共享资源读写，除了互斥锁，还有atomic
 */
func GoSyncAtomic() {
	var count int64 = 0
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go addInt(&wg, &count)
	}
	wg.Wait()
	fmt.Println("data:", count)
}

func addInt(wg *sync.WaitGroup, data *int64) {
	defer wg.Done()
	atomic.AddInt64(data, 1)
	//协程暂停
	runtime.Gosched()
}
