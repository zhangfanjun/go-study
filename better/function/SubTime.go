package function

import (
	"fmt"
	"time"
)

/**
* 两种方法进行时间的统计，
* 第一种是采用since
* 另一种是采用time.Now().Sub
 */
func SubTimeTest() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		time.Sleep(2)
	}
	since := time.Since(start)
	sub := time.Now().Sub(start)
	fmt.Println(since)
	fmt.Println(sub)
}
