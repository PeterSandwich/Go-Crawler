package parser

import (
	"crawler/engine"
	"regexp"
)


var cityRe = regexp.MustCompile(`<a href="([^"]+)" target="_blank">([^<]+)</a>`)
var nextPageRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)`)
var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
// 解析单个城市页面内的数据
func ParseCity(contents []byte, _ string) engine.ParseResult {

	result := engine.ParseResult{}

	sexMatch := sexRe.FindAllSubmatch(contents,-1)
	match := cityRe.FindAllSubmatch(contents,-1)
	//log.Println(len(sexMatch),len(match))
	for i,m:= range match {
		url:= string(m[1])
		name := string(m[2])
		//result.Items = append(result.Items,string(m[2]))
		//log.Println("Get User: ",string(m[2]),string(m[1]))
		sex := "未知"
		if i< len(sexMatch){
			sex = string(sexMatch[i][1])
			//log.Println(sex)
		}
		result.Requests = append(result.Requests,engine.Request{
			Url:url,
			ParserFunc: ProfileParser(name,sex),
		})
	}

	match = nextPageRe.FindAllSubmatch(contents,-1)
	for _,m:= range match {
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc: ParseCity,
		})
	}


	return result
}


func ProfileParser(name,sex string) engine.ParserFunc {
	return  func(c []byte,url string) engine.ParseResult{
		return ParseProFile(c,url,name,sex)
	}
}