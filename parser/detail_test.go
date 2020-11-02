package parser

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"spider/engine"

	"testing"
)

func TestParseDetail(t *testing.T) {
	url := "https://newcar.xcar.com.cn/m25935/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 转为utf8
	reader := bufio.NewReader(resp.Body)
	encode := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, encode.NewDecoder())
	byte, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	// 造数据
	name := ""
	expected := engine.Details{
		Name: name,
		Rank: "SUV",
		Engine: "185kW(3.0L自然吸气)",
		PowerType: "汽油机",
		Gearbox: "5挡AT",
		Size: "4695×1815×1825",
		BodyStructure: "5门 7座 SUV",
		ListedTime: "2015",
		OilWear: 10.8,
	}

	//把拿到的和自己造的对比
	res := ParseDetail(byte, name)
	if err != nil {
		panic(err)
	}
	got  := res.Items[0]
	if expected != got {
		t.Errorf("Expected: %v, but got %v\n", expected, got)
	}
	t.Logf("Result: %v\n", res)
}

// 获取编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024) // 先将resp.Body转换bufio.NewReader，再Peek，然后读取bufio.NewReader，所以指针并没有移动，读取的内容是全的
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}