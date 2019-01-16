package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)

var rateLimiter = time.Tick(100*time.Millisecond)
// 根据URL获取网页（转码为UTF-8编码）
func Fetch(url string)([]byte,error) {
	<-rateLimiter
	request,err:=http.NewRequest(http.MethodGet,url,nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	resp,err :=http.DefaultClient.Do(request)
	//resp,err := http.Get(url)
	if err != nil {
		return []byte(""),err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("error status code: %d\n",resp.StatusCode)
	}

	// 文本转换 手动
	//utf8Reader:=transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	// all ,err := ioutil.ReadAll(utf8Reader)

	// 文本转换 自动
	bodyReader := bufio.NewReader(resp.Body)
	e := determicneEncoding(bodyReader)
	utf8Reader:=transform.NewReader(resp.Body,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

// 确定编码格式
func determicneEncoding(r *bufio.Reader) encoding.Encoding {
	bytes,err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8
	}
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	return e
}
