package synclock

import (
	"sync"
)

type MyData struct {
	lock sync.Mutex
	data int
	flag int
}

func (d *MyData) SyncWrite(data int) {
	//锁定
	d.lock.Lock()
	//退出函数的时候解除锁定
	defer d.lock.Unlock()
	if d.flag < 1000 {
		d.data += data
		d.flag += 1
	}
}

func (d *MyData) Write(data int) {
	if d.flag < 1000 {
		d.data += data
		d.flag += 1
	}
}
