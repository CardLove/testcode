package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	now := time.Now()
	h, _ := time.ParseDuration("24h")
	fmt.Println(int32(now.Add(2 * h).Unix()))
	var num int32 = 23
	var num1 int32 = 2

	//sdsdds   ddd反反复复妮妮的
	fmt.Println(reflect.TypeOf(num*num1 + int32(3)))

}
