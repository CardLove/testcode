package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ServiceData struct {
	ch   chan int // 用来 同步的channel
	data []int    // 存储数据的slice
}

func (s *ServiceData) Schedule() {
	// 从 channel 接收数据
	for i := range s.ch {
		s.data = append(s.data, i)
	}
}

func (s *ServiceData) Close() {
	// 最后关闭 channel
	close(s.ch)
}

func (s *ServiceData) AddData(v int) {
	s.ch <- v // 发送数据到 channel
}

func NewScheduleJob(size int, done func()) *ServiceData {
	s := &ServiceData{
		ch:   make(chan int, size),
		data: make([]int, 0),
	}

	go func() {
		// 并发地 append 数据到 slice
		s.Schedule()
		done()
	}()

	return s
}

func main() {
	arr := make([]int32, 0)

	go func() {
		for {
			arr = append(arr, rand.Int31n(444))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			for i, value := range arr {
				fmt.Println("1&:", i, value)
			}
			time.Sleep(1 * time.Second)
			fmt.Println("1 arr len ", len(arr))

		}
	}()

	go func() {
		for {
			for i, value := range arr {
				fmt.Println("2&:", i, value)
			}
			fmt.Println("2 arr len ", len(arr))
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(3 * time.Minute)
}
