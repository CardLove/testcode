package actualcombat

type Director struct {
	Builder InterfaceBuilder
}

func NewDirector(b InterfaceBuilder) *Director {
	return &Director{Builder: b}
}

func (d *Director) SetBuilder(b InterfaceBuilder) {
	d.Builder = b
}

func (d *Director) BuilderCar() Car {

	d.Builder.SetEngineType()
	d.Builder.SetNumber()
	d.Builder.SetSeatsType()
	return d.Builder.GetCar()
}
