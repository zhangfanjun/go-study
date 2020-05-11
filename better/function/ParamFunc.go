package function

import (
	"fmt"
)

/**
* 入参为一个func，必须实现的匿名函数，可以达到java的接口的效果，用于分布式锁
 */
func ParamFunc() {
	showData(func() {
		fmt.Println("这是一个匿名的方法")
	})
	showData(myShow)
}

func showData(s show) {
	fmt.Println("------------------------")
	s()
	fmt.Println("这里使用方法作为类型，定义之后，s就是匿名方法的方法名，而show的func后面括号就是对应的入参，s（）就是调用匿名的方法")

}

type show func()

func myShow() {
	fmt.Println("------myShow()采用相同的方法入参")
}
