package coroutines

import "testing"

func TestCoroutine(t *testing.T) {
	//线程安全，同步锁
	//GoSyncLock()
	//线程的阻塞
	//GoSyncWait()
	//协程安全的做加法
	//GoSyncAtomic()
	//通道的同步并发设计
	GoChannelWait()
}
