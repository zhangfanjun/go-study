package function

import "fmt"

/**
* 使用goto功能，从循环中跳到任意的地方，灵活性很强，但是考虑return的问题
 */
func ManyForAndGoTo() {
	for {
		for i := 1; i < 20; i++ {
			fmt.Println(i)
			if i == 15 {
				goto print
			}
		}
	}
print:
	fmt.Println("跳出两个for循环")
}
