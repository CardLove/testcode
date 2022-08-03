package main

import (
	"fmt"
	"sync"
)

type stu struct {
	Name string
	Age  int
}

func main() {
	//a()
	//b()
	//c()
	//d()
	e()
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
func e() {
	user := new(UserAges)
	user.Add("whx", 23)
	fmt.Println(user.Get("whx"))
}
func d() {
	int_chan := make(chan int)
	string_chan := make(chan string)
	int_chan <- 1
	string_chan <- "ssss"
	select {
	case v := <-int_chan:
		fmt.Println(v)
	case v := <-string_chan:
		panic(v)

	}
}
func c() {
	wg := sync.WaitGroup{}
	wg.Add(30)
	for i := 0; i < 10; i++ {
		wg.Done()
	}
	wg.Wait()
}
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func a() {
	m := make(map[string]*stu)
	stus := []stu{
		{Name: "sd",
			Age: 32,
		},
		{Name: "23333",
			Age: 11,
		},
	}
	for _, s := range stus {
		fmt.Println(s)
		m[s.Name] = &s
	}
	fmt.Println("%+v", m)
}
