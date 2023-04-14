package clother

type Iclother interface {
	SetSize(size int32)
	SetName(name string)
	GetSize() int32
	GetName() string
}

type Clother struct {
	Size int32
	Name string
}

func (p *Clother) SetName(name string) {
	p.Name = name
}

func (p *Clother) SetSize(size int32) {
	p.Size = size
}

func (p *Clother) GetName() string {
	return p.Name
}

func (p *Clother) GetSize() int32 {
	return p.Size
}
