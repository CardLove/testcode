package main

import (
	"bytes"
	"fmt"
	"testing"

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

func TestXxx(t *testing.T) {

	v := viper.New()
	v.SetConfigType("json") // 设置配置文件的类型

	// 配置文件内容
	// var jsonExample = []byte(`{
	// 	"code": 0,
	// 	"data": {
	// 	  "total": 1,
	// 	  "list": [
	// 		{
	// 		  "id": 1,
	// 		  "name": "test",
	// 		  "type": "harbor",
	// 		  "url": "http://192.168.4.69",
	// 		  "auto_scan": false,
	// 		  "registry_status": 1,
	// 		  "version": "v2.0.2-e91b4ff1",
	// 		  "has_vul": true,
	// 		  "fail_result": ""
	// 		}
	// 	  ]
	// 	},
	// 	"msg": "操作成功",
	// 	"request_id": "8ad383e5f1620145130d"
	//   }`)
	var jsonExample = []byte(`{
		"objArray": [
		  {"name": "obj1", "value": 1},
		  {"name": "obj2", "value": 2}
		],
		"sliceArray": [1, 2, 3, 4, 5]
	  }`)

	//创建io.Reader
	v.ReadConfig(bytes.NewBuffer(jsonExample))
	fmt.Println("获取配置文件的port", v.IsSet("list"))
	fmt.Println("获取配置文件的port", v.Get("list"))

	info := v.GetStringMap("objArray")
	fmt.Println(info)
	for key, value := range info {
		fmt.Println(key, value)
	}

	sliceArray := v.GetStringMapStringSlice("sliceArray")
	for _, v := range sliceArray {
		fmt.Println(v)
	}

	// fmt.Println("获取配置文件的port", v.AllSettings())

}
