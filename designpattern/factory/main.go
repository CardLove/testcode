package main

import (
	"fmt"
	"testcode/designpattern/factory/anta"
	"testcode/designpattern/factory/clother"
)

func MakeClother(clotherType string) (clother.Iclother, error) {
	if clotherType == "anta" {
		return anta.NewAnta(), nil
	}
	return nil, fmt.Errorf("Wrong clothes type passed")
}
func main() {
	fmt.Println("factory")
	Anta, err := MakeClother("anta")
	if err == nil {
		PrintDetails(Anta)
	}

}

func PrintDetails(c clother.Iclother) {
	fmt.Println(c.GetName())
	fmt.Println(c.GetSize())
}
