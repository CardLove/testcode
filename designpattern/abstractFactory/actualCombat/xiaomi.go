package actualcombat

type XiaomiFactory struct {
}

func (p *XiaomiFactory) MakePhone() AbstractPhone {
	return &XiaomiPhone{
		Phone{
			Color: "white",
			Size:  12,
		},
	}

}
func (p *XiaomiFactory) MakeComputer() AbstractComputer {
	return &XiaomiComputer{
		Computer{
			Color: "balck",
			Size:  15,
		},
	}
}
