package actualcombat

type LenovoFactory struct {
}

func (p *LenovoFactory) MakeComputer() AbstractComputer {
	return &LenovoComputer{
		Computer{
			Color: "red",
			Size:  1,
		},
	}
}
func (p *LenovoFactory) MakePhone() AbstractPhone {
	return &LenovoPhone{
		Phone{
			Color: "green",
			Size:  3},
	}
}
