package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	var tomlExample = []byte(`
	title = ""

	[owner]
	organization = "MongoDB"
	Bio = "MongoDB Chief Developer Advocate & Hacker at Large"
	dob = 1979-05-27T07:32:00Z # First class dates? Why not?`)
	v.SetConfigType("toml")
	v.ReadConfig(bytes.NewBuffer(tomlExample))
	fmt.Println(v.GetString("title"))
	fmt.Println(v.Get("title"))
	fmt.Println(v.GetString("owner.organization"))
	fmt.Println(v.Get("owner.organization"))

	fmt.Println(v.IsSet("title"))

}
