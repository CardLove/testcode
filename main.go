package main

import "fmt"

var max int64
var str []string

func main() {
	CreateNum(4)
	fmt.Println("max", max)
	fmt.Println("end")
}

func CreateNum(number int32) {
	seedNumber := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	//seedNumber := []string{"0", "1", "2", "3", "4", "5"}

	ret := []string{}
	base := seedNumber
	for j := int32(0); j < number-1; j++ {
		for _, v := range base {
			for _, v1 := range seedNumber {
				newNum := v + v1
				ret = append(ret, newNum)
				if int32(len(newNum)) == number {
					max++
					//fmt.Println(newNum)
				}
			}
		}
		base = ret
	}
	//fmt.Println(max)
}
