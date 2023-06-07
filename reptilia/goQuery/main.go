package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// func main() {
// 	// 发起 HTTP 请求获取网页内容
// 	resp, err := http.Get("https://example.com")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	// 解析 HTML
// 	doc, err := goquery.NewDocumentFromReader(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 使用选择器提取数据
// 	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
// 		fmt.Println(s.Text())
// 	})

// 	doc.Find("a").Each(func(i int, s *goquery.Selection) {
// 		link, _ := s.Attr("href")
// 		fmt.Println(link)
// 	})
// }

func main() {
	// 发起 HTTP 请求获取网页内容
	resp, err := http.Get("https://developer.huaweicloud.com/euleros/cvedetail.html?cveId=CVE-2023-0286")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 使用选择器提取数据
	doc.Find(".content ").Each(func(i int, s *goquery.Selection) { //.wrapper .content-block .news .memoryfirst
		fmt.Println(s.Text())

		content := s.Find(".title.ng-binding p").Text()
		fmt.Println(content)
	})

	// doc.Find("a").Each(func(i int, s *goquery.Selection) {
	// 	link, _ := s.Attr("href")
	// 	fmt.Println(link)
	// })
}
