package actualcombat

type AbstractComputer interface {
	SetSize(size int)
	SetColor(color string)
	GetSize() int
	GetColor() string
}

type Computer struct {
	Size  int
	Color string
}

func (p *Computer) SetSize(size int) {
	p.Size = size
}
func (p *Computer) SetColor(color string) {
	p.Color = color
}

func (p *Computer) GetSize() int {
	return p.Size
}
func (p *Computer) GetColor() string {
	return p.Color
}
