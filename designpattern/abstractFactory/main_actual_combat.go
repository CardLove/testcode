package main

import (
	"fmt"
	actualcombat "testcode/designpattern/abstractFactory/actualCombat"
	"testcode/designpattern/abstractFactory/example"
)

func main() {
	xiaomiFactory, _ := actualcombat.GetElectronFactory("xiaomi")
	lenovpFactory, _ := actualcombat.GetElectronFactory("lenovo")
	xiaomiPhone := xiaomiFactory.MakeComputer()
	xiaomiComputer := xiaomiFactory.MakePhone()

	lenovoPhone := lenovpFactory.MakeComputer()
	lenovoComputer := lenovpFactory.MakeComputer()

	printPhoneDetails(lenovoPhone)
	printComputerDetails(lenovoComputer)

	printPhoneDetails(xiaomiPhone)
	printComputerDetails(xiaomiComputer)

	//测试2
	fmt.Println("******************")
	factory := example.NewConcreteFactory()
	product := factory.CreateProduct()
	product.GetName()

}

func printPhoneDetails(s actualcombat.AbstractPhone) {
	fmt.Printf("Color: %s", s.GetColor())
	fmt.Println()
	fmt.Printf("Size: %d inch", s.GetSize())
	fmt.Println()
}

func printComputerDetails(s actualcombat.AbstractComputer) {
	fmt.Printf("Color: %s", s.GetColor())
	fmt.Println()
	fmt.Printf("Size: %d inch", s.GetSize())
	fmt.Println()
}
