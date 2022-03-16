package main

import "fmt"

func modify(array [5]int) {
	array[0] = 10
}

func main() {
	// 传入数组，注意数组与slice的区别
	array := [5]int{1, 2, 3, 4, 5}

	modify(array)
	fmt.Println(array) //添加注释

}
