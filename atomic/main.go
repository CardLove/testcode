package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var newValue int32 = 200
	var dst int32 = 100
	old := atomic.SwapInt32(&dst, newValue)
	fmt.Println("old value:", old, "new value :", dst)

}

/*
1.五种类型的原子操作
swap  compare-and-swap  add load store
2.有一种可以存储go 基本类型值的 value
*/
