package Builder

type Builder interface {
	Build()
}

type Director struct {
	Builder Builder
}

func NewDirector(b Builder) Director {
	return Director{Builder: b}
}

func (d *Director) Construct() {
	d.Builder.Build()
}

type ConcreteBuilder struct {
	build bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{false}
}

func (b *ConcreteBuilder) Build() {
	b.build = true
}

type Product struct {
	Build bool
}

func (b *ConcreteBuilder) GetResult() Product {
	return Product{b.build}
}
