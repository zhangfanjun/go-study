package coroutines

import "testing"

func TestCoroutine(t *testing.T) {
	//线程安全，同步锁
	//GoSyncLock()
	//线程的阻塞
	//GoSyncWait()
	//协程安全的做加法
	//GoSyncAtomic()
	//通道的并发同步设计
	//GoChannelWait()
	//遍历通道
	//GoChannelRange()
	//生产者和消费者
	Factory()
}
