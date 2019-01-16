package parser

import (
	"crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
func ParseCityList(contens []byte,_ string) engine.ParseResult {

	result := engine.ParseResult{}
	match := cityListRe.FindAllSubmatch(contens,-1)

	for _,m := range match {
		//result.Items = append(result.Items,string(m[2]))

		result.Requests = append(result.Requests,
			engine.Request{
				Url:		string(m[1]),
				ParserFunc: ParseCity},
		)
	}


	return result
}
