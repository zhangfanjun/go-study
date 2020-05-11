package function

import "testing"

/**
* 在命名文件时文件名必须以_test.go结尾
* 入参*testing.T表示单元测试，方法以Test开头
* 入参*testing.B表示性能测试，方法以Benchmark开头
 */
/*单元测试*/
func TestMyTest(t *testing.T) {
	sum := myTest(11, 22)
	if sum != 33 {
		t.Error("加法的结果不等于33")
	} else {
		t.Log("加法的结果为33")
	}
}
func myTest(a, b int) int {
	return a + b
}

/*性能测试*/
func BenchmarkMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myTest(11, 22)
	}
}
