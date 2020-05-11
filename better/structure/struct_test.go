package structure

import (
	"fmt"
	"testing"
)

func TestStructure(t *testing.T) {
	san := GetZhangSan()
	fmt.Println("结构体的构造方法", san)
	fatherAndSunTest()
	NiMingStructTest()
}
