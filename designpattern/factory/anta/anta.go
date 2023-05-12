package anta

import "testcode/designpattern/factory/clother"

type Anta struct {
	clother.Clother
}

func NewAnta() clother.Iclother {
	return &Anta{
		clother.Clother{
			Name: "anta",
			Size: 1,
		},
	}
}
