package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)


//var timeLater = time.Tick(10 * time.Millisecond)
/**
 * 抓取器
 */
func Fetcher(url string) ([]byte, error) {
	//<-timeLater // 限定时间防止被对方网站限制

	// 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// 延后关闭
	defer resp.Body.Close()

	// 转为utf8编码
	bodyReader := bufio.NewReader(resp.Body)
	encode := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, encode.NewDecoder())

	// 获取页面内容
	//return utf8Reader, nil
	return ioutil.ReadAll(utf8Reader)
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
