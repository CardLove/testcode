package main

import (
	"fmt"
)

func main() {
	arr := make([]int32, 4)
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)

	arr = append(arr, int32(12))
	arr = append(arr, int32(12))
	arr = append(arr, int32(12))
	arr = append(arr, int32(12))
	arr = append(arr, int32(12))

	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)

}
