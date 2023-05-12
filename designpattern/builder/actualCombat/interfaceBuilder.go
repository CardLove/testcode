package actualcombat

type InterfaceBuilder interface {
	SetSeatsType()
	SetEngineType()
	SetNumber()
	GetCar() Car
}

func GetBuilder(BuilderType string) InterfaceBuilder {
	if BuilderType == "MPV" {

	}

	return nil
}
