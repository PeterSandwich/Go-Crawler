package view

import (
	"testing"
	"os"
	"crawler/engine"
	commom "crawler/model"
	"crawler/frontend/model"
)

func TestTemplate(t *testing.T){
	view:= CreateSearchResultView("template.html")
	out,err := os.Create("tempalte.test.html")
	expect_profile:=engine.Item{
		Url: "http://album.zhenai.com/u/74822048",
		Type: "zhenai",
		Id: "12154555",
		Payload: commom.Profile{
			Name: 		"test",
			Age: 		"26岁",
			Sex:		"男士",
			Height: 	"157cm",
			Income: 	"3001-5000元",
			Marriage: 	"未婚",
			Education: 	"大学本科",
			City:		"珠海",
			ImageUrl:	"https://photo.zastatic.com/images/photo/27375/109498461/16870242661675398.png",
		},
	}

	data := model.SearchReqult{
		Hits:123,
		Start: 0,
		Items: []engine.Item{expect_profile,expect_profile},
	}
	err = view.Render(out,data)
	if err != nil{
		panic(err)
	}
}
