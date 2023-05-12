package actualcombat

type SuvBuilder struct {
	EngineType string
	Seatstype  string
	Number     int
}

func NewSuvBuilder() *SuvBuilder {
	return &SuvBuilder{}
}
func (s *SuvBuilder) SetEngineType() {
	s.EngineType = "V10"
}
func (s *SuvBuilder) SetSeatsType() {
	s.Seatstype = "手动座椅"
}
func (s *SuvBuilder) SetNumber() {
	s.Number = 1
}

func (s *SuvBuilder) Get() Car {
	return Car{
		SetSeatsType:  s.Seatstype,
		SetEngineType: s.EngineType,
		Number:        s.Number,
	}
}
