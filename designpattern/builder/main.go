package main

import (
	"fmt"
	actualcombat "testcode/designpattern/builder/actualCombat"
)

func mian() {
	// 声明mpv生成器对象
	MpvBuilder := actualcombat.GetBuilder("MPV")
	// 声明suv 生成对象
	SuvBuilder := actualcombat.GetBuilder("SUV")
	fmt.Println(MpvBuilder)
	fmt.Println(SuvBuilder)
}
