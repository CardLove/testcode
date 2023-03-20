package main

import "fmt"

type Animal struct {
	Name   string
	Age    int
	Height int
	Weight int
	CanRun bool
	LegNum int
}
type OptoinFunc func(*Animal)

func NewAnimal(name string, opts ...OptoinFunc) *Animal {
	a := &Animal{Name: name, CanRun: true, LegNum: 4}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithName(name string) OptoinFunc {
	return func(a *Animal) { a.Name = name }
}
func WithAge(age int) OptoinFunc {
	return func(a *Animal) { a.Age = age }
}

func WithHeight(height int) OptoinFunc {
	return func(a *Animal) { a.Height = height }
}

func main() {
	a2 := NewAnimal("dog",WithAge(1),WithHeight(120))
	fmt.Printf("%+v",a2)
	a3 := NewAnimal("cat",WithHeight(30))
	fmt.Printf("%v",a3)
}
