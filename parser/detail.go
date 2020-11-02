package parser

import (
	"bytes"
	"encoding/binary"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math"
	"regexp"
	"spider/engine"
	"strconv"
)



/**
 * 详情页解析器：提取汽车详细字段
 */
func ParseDetail(content []byte, name string) engine.ParserResult {
	result := engine.ParserResult{}
	// 获取dom对象
	reader := bytes.NewReader(content)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Printf("Goquery detail new reader error :%v\n", err)
	}

	// 提取汽车详细字段
	tdObj := doc.Find(".w237")

	detail := engine.Details{
		Name:          name,
		Rank:          getTdHtml(tdObj, 0),
		Engine:        getTdHtml(tdObj, 1),
		PowerType:     getTdHtml(tdObj, 2),
		Gearbox:       getTdHtml(tdObj, 3),
		Size:          getTdHtml(tdObj, 4),
		BodyStructure: getTdHtml(tdObj, 5),
		ListedTime:    getTdHtml(tdObj, 6),
		OilWear:       getTdHtml(tdObj, 9),
	}
	result.Items = append(result.Items, detail)

	// 由于详情页没有需要爬的url所以result.Requests就不用管了返回result.Items即可
	return result
}

// 提取函数，返回字符串
func getTdHtml(tdObj *goquery.Selection, index int) string {
	html, err := tdObj.Eq(index).Html()
	if err != nil {
		log.Printf("Goquery detail string error: %v\n", err)
	}
	return html
}

// 提取函数，返回浮点型
func getTdHtmlFloat64(tdObj *goquery.Selection, index int) float64 {
	html, err := tdObj.Eq(index).Html()
	if err != nil {
		log.Printf("Goquery detail float64 error: %v\n", err)
	}
	newHtml, err := strconv.ParseFloat(html, 64)
	if err != nil {
		log.Printf("Strconv string to float64 error: %v\n", err)
	}
	return newHtml
}

// 正则提取函数，返回字符串
func matchContent(content []byte, re string) string {
	res := regexp.MustCompile(re)
	match := res.FindSubmatch(content)
	return string(match[1])
}

// 正则提取函数，返回浮点型
func matchContentFloat(content []byte, re string) float64 {
	res := regexp.MustCompile(re)
	match := res.FindSubmatch(content)
	bits := binary.LittleEndian.Uint64(match[1])
	return math.Float64frombits(bits)
}
