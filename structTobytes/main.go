package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	House    int
	Street1  string
	Town     string
	PostCode PostCode
}

type PostCode struct {
	Value string
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func Compress(s []byte) []byte {

	zipbuf := bytes.Buffer{}
	zipped := gzip.NewWriter(&zipbuf)
	zipped.Write(s)
	zipped.Close()
	fmt.Println("compressed size (bytes): ", len(zipbuf.Bytes()))
	return zipbuf.Bytes()
}

func Decompress(s []byte) []byte {

	rdr, _ := gzip.NewReader(bytes.NewReader(s))
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		log.Fatal(err)
	}
	rdr.Close()
	fmt.Println("uncompressed size (bytes): ", len(data))
	return data
}

func DecodeToPerson(s []byte) Person {

	p := Person{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func WriteToFile(s []byte, file string) {

	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}

	f.Write(s)
}

func ReadFromFile(path string) []byte {

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {

	p := Person{
		Name: "Joe Bloggs",
		Address: Address{
			House:   1,
			Street1: "The Lane",
			Town:    "Blackburn",
			PostCode: PostCode{
				Value: "BB2 5LB",
			},
		},
	}
	fmt.Printf("p:%+v 变量地址：%p\n", p, &p)
	p1 := p
	fmt.Printf("p1:%+v 变量地址：%p\n", p1, &p1)
	p1.Name = "aa"
	fmt.Printf("p:%+v 变量地址：%p\n", p, &p)
	fmt.Printf("p1:%+v 变量地址：%p\n", p1, &p1)

	// p1 := Person{
	// 	Name: "Joe Bloggs",
	// 	Address: Address{
	// 		House:   1,
	// 		Street1: "The Lane",
	// 		Town:    "Blackburn",
	// 		PostCode: PostCode{
	// 			Value: "BB2 5LB",
	// 		},
	// 	},
	// }

	// data := EncodeToBytes(p)
	// data1 := EncodeToBytes(p1)

	// if bytes.Compare(data, data1) == 0 {
	// 	log.Println("ok")
	// }
	// if bytes.Equal(data, data1) {
	// 	log.Println("equal ok ")
	// }

	// dataOut = Compress(dataOut)
	// WriteToFile(dataOut, "person.dat")

	// dataIn := ReadFromFile("person.dat")
	// dataIn = Decompress(dataIn)
	// p2 := DecodeToPerson(dataIn)

	// fmt.Println(p2)
}
