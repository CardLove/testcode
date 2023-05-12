package actualcombat

type MpvBuilder struct {
	Seatstype  string
	EngineType string
	Number     int
}

func NewMpvBuilder() *MpvBuilder {
	return &MpvBuilder{}
}

func (m *MpvBuilder) SetSeatsType() {
	m.Seatstype = "太空座椅"
}

func (m *MpvBuilder) SetEngineType() {
	m.EngineType = "v8"
}

func (m *MpvBuilder) SetNumber() {
	m.Number = 10
}

func (m *MpvBuilder) GetCar() Car {
	return Car{
		SetSeatsType:  m.EngineType,
		SetEngineType: m.EngineType,
		Number:        m.Number,
	}
}
