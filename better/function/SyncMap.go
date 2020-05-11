package function

import (
	"fmt"
	"golang.org/x/sync/syncmap"
)

/**
* 线程安全的map
 */
func SyncMapTest() {
	fmt.Println("map的读写同时操作是不安全的，这个时候需要线程安全的syncMap，查看底层发现是用sync.Mutex达到线程安全")
	var myMap syncmap.Map
	fmt.Println("------------添加操作key1，value1--------------------")
	myMap.Store("key1", "value1")
	if value, ok := myMap.Load("key1"); ok {
		fmt.Println("------------获取key1的value值----------------")
		fmt.Println(value)
	}
	fmt.Println("------------删除key1--------------------")
	myMap.Delete("key1")
	if value, ok := myMap.Load("key1"); ok {
		fmt.Println("------------获取key1的value值----------------")
		fmt.Println(value)
	}
	myMap.Store("key1", "value1")
	myMap.Store("key2", "value2")
	myMap.Store("key3", "value3")
	myMap.Store("key3", "value4")
	myMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		//return false，返回false将不进行下一个值的遍历，测试发遍历是没有顺序的
		return true
	})
}
