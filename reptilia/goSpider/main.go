package main

// import (
// 	"fmt"

// 	"github.com/sundy-li/go-spider/core/common/page"
// 	"github.com/sundy-li/go-spider/core/common/request"
// 	"github.com/sundy-li/go-spider/core/spider"
// 	"github.com/sundy-li/go-spider/core/spider/pipeline"
// 	"github.com/sundy-li/go-spider/core/spider/scheduler"
// 	"github.com/sundy-li/go-spider/core/spider/spiderkit"
// )

// // 定义一个处理器结构体
// type MyProcessor struct {
// }

// // 实现 Process 方法，用于处理爬取到的页面数据
// func (p *MyProcessor) Process(pdc *page.Page) {
// 	// 解析页面数据，提取需要的信息
// 	title := pdc.GetHtmlParser().GetText("title")
// 	fmt.Println("Title:", title)
// }

// func main() {
// 	// 创建爬虫对象
// 	mySpider := spider.NewSpider(&MyProcessor{})

// 	// 创建一个初始请求
// 	startReq := request.NewRequest("https://example.com", "GET", "", "", nil, nil)

// 	// 设置爬虫配置
// 	mySpider.SetThreadnum(1)                         // 设置并发线程数
// 	mySpider.SetSleepTime("rand", 300, 1000)         // 设置随机睡眠时间，防止频繁访问
// 	mySpider.SetAcceptedDomains([]string{"example"}) // 设置可接受的域名

// 	// 设置请求过滤规则
// 	mySpider.SetRequestProcessors([]spider.RequestProcessor{
// 		spiderkit.NewDefaultResponseProcessor(nil), // 默认的响应处理器
// 		spiderkit.NewDefaultUrlFilter(),            // 默认的 URL 过滤器
// 	})

// 	// 设置页面处理管道
// 	mySpider.SetPipeline(pipeline.NewConsolePipeline()) // 输出到控制台

// 	// 创建调度器并启动爬虫
// 	sched := scheduler.NewScheduler()
// 	sched.Start()
// 	sched.SendReq(startReq)
// 	mySpider.SetScheduler(sched)
// 	mySpider.Run()
// }
