package main

import (
	"flag"
	"fmt"
	"strings"
)

var opt Option

type Option struct {
	SetPluginMem mapFlag
}

type value interface {
	String() string
	Set(string) error
}
type sliceFlag []string
type mapFlag map[string]string

func (f *sliceFlag) String() string {
	return fmt.Sprintf("%v", []string(*f))
}
func (f *sliceFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}
func (f mapFlag) String() string {
	return fmt.Sprintf("%v", map[string]string(f))
}
func (f mapFlag) Set(value string) error {
	split := strings.SplitN(value, "=", 2)
	fmt.Println("split:", split)
	f[split[0]] = split[1]
	return nil
}
func main() {
	var hostsFlag sliceFlag
	opt.SetPluginMem = mapFlag{}
	flag.Var(&hostsFlag, "host", "Application hosts, for example: -host=a.com -host=b.com")
	flag.Var(&opt.SetPluginMem, "mem", "ssss")
	flag.Parse()
	// fmt.Println(hostsFlag.String())
	fmt.Println(opt.SetPluginMem)
	// if v, ok := opt.SetPluginMem["1"]; ok {
	// 	fmt.Println("v", v)
	// }

	var err error

	if err == nil {
		fmt.Println("err == nil ok ")
	}

}
