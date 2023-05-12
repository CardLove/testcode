package actualcombat

type AbstractPhone interface {
	SetSize(size int)
	SetColor(color string)
	GetSize() int
	GetColor() string
}
type Phone struct {
	Size  int
	Color string
}

func (p *Phone) SetSize(size int) {
	p.Size = size
}
func (p *Phone) SetColor(color string) {
	p.Color = color
}
func (p *Phone) GetSize() int {
	return p.Size
}
func (p *Phone) GetColor() string {
	return p.Color
}
