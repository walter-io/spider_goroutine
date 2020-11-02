package parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
	"spider/config"
	"spider/engine"
)

/**
 * 首页解析器
 */
func ParseIndex(content []byte) engine.ParserResult {
	// 获取dom对象
	result := engine.ParserResult{}
	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Printf("Goquery index new reader error: %v\n", err)
	}

	// 提取url
	doc.Find(".car_col2").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Find("a").Eq(0).Attr("href")
		// 这部分逻辑是：给url，并且指定该页面用哪个解析器解析
		result.Items = append(result.Items, engine.Details{})
		result.Requests = append(result.Requests, engine.Request{
			Url:        config.DomainName + url,
			ParserFunc: ParseLists,
		})
	})
	return result
}
