package main

import (
	"fmt"
	"sync"
)

func main() {

	m := &sync.Map{}
	m.Store("1", "")
	fmt.Println(m.Load("1"))

	m.Delete("1")
	fmt.Println(m.Load("1"))
	m.Store("1", "")
	fmt.Println(m.Load("1"))

	fmt.Println("sdsd")

}
