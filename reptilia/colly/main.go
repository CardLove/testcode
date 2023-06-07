package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	// 创建一个新的 Colly 收集器
	c := colly.NewCollector()

	// 定义访问的网址
	url := "https://example.com"

	// 在访问之前执行的操作
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// 在访问成功之后执行的操作
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response received:", r.StatusCode)
	})

	// 定义要查找的数据
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
	})

	// 开始爬取数据
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
