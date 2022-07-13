package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte{0, 1, 3, 2}
	b := []byte{0, 1, 3, 2}
	c := []byte{1, 1, 3, 2}

	fmt.Println(bytes.Equal(a, b))
	fmt.Println(bytes.Equal(a, c))
}
