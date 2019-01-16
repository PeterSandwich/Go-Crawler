package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
	"crawler/model"
	"github.com/gpmgo/gopm/modules/log"
)

var basicInfoRe= regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)
//var nickNameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>([^<]+)</h1>`)
var imageRe = regexp.MustCompile(`<div class="logo f-fl" style="background-image:url\(([^?\)]*)[?\)]`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)


func ParseProFile(contents []byte, url string,name string,sex string)engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name
	 match:=extractString(contents,basicInfoRe)
	 if len(match)<6 {
	 	log.Error("basicInfoRe extractString error")
	 }else {
		 str := strings.Split(string(match), "|")
		 profile.City = strings.Trim(str[0], " ")
		 profile.Age = strings.Trim(str[1], " ")
		 profile.Education = strings.Trim(str[2], " ")
		 profile.Marriage = strings.Trim(str[3], " ")
		 profile.Height = strings.Trim(str[4], " ")
		 profile.Income = strings.Trim(str[5], " ")
		 profile.Sex = sex
	 }


	 match = extractString(contents,imageRe)
	 profile.ImageUrl = match
	 if len(match)==0{
	 	log.Error("imageRe basicInfoRe extractString error")
	 }

	//fmt.Println("====>",profile)
	 return engine.ParseResult{Requests:nil,Items:[]engine.Item{
	 	{
	 		Url: url,
	 		Type: "zhenai",
	 		Id: extractString([]byte(url),idUrlRe),
	 		Payload:profile,
	 	},
	 }}

}

func extractString(contens []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contens)
	if len(match)>=2 {
		return string(match[1])
	}else {
		return ""
	}
}