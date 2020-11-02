package main

import (
	"spider/engine"
	"spider/parser"
	"spider/scheduler"
)

/**
 * go mod init
 * go get github.com/PuerkitoBio/goquery
 * go get golang.org/x/net/html/charset
 * go get golang.org/x/text/encoding
 */


/**
 * 爬虫入口
 */
func main() {
	url := "https://newcar.xcar.com.cn/car/"
	// 指定调度器
	e := engine.Engine{
		Scheduler: &scheduler.Scheduler{},
	}
	// 开始运行
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseIndex,
	})
}
