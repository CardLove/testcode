package Builder

import (
	"fmt"
	"testing"
)

func TestConcreteBuilder_Build(t *testing.T) {
	builder := NewConcreteBuilder()
	director := NewDirector(&builder)
	director.Construct()
	product := builder.GetResult()
	fmt.Println(product.Build)
}
