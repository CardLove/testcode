package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool

func read(name string, c *sync.Cond) {
	c.L.Lock()
	defer c.L.Unlock()
	for !done {
		c.Wait()
	}
	fmt.Println(name, "start reading")
}
func write(name string, c *sync.Cond) {
	fmt.Println(name, "start writing")
	time.Sleep(time.Second)
	done = true
	fmt.Println(name, "wakes all ")
	c.Broadcast()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)
	time.Sleep(time.Second * 4)
}
