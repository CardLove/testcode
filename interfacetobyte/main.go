package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type player struct {
	name string
	uid  int64
}

func main() {
	p := player{
		name: "whx",
		uid:  11,
	}
	fmt.Println(GetBytes(p))
}
