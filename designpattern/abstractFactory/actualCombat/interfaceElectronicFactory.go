package actualcombat

import "fmt"


type InterfaceElectronicFactory interface {
	MakePhone() AbstractPhone
	MakeComputer() AbstractComputer
}

func GetElectronFactory(brand string) (InterfaceElectronicFactory, error) {
	if brand == "xiaomi" {
		return &XiaomiFactory{}, nil
	}
	if brand == "lenovo" {
		return &LenovoFactory{}, nil
	}
	return nil,fmt.Errorf("brand type err")
}
	