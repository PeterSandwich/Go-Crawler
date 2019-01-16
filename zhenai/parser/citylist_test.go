package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents,err := ioutil.ReadFile("test_citylist.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents,"")

	const resiltSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	//expectedCities :=[]string{
	//	"阿坝","阿克苏","阿拉善盟",
	//}

	if len(result.Requests) != resiltSize {
		t.Errorf("result should have %d  requests;but had %d",resiltSize,len(result.Requests))
	}

	for i,url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s",i,url,result.Requests[i].Url)
		}
	}

	//if len(result.Items) != resiltSize {
	//	t.Errorf("result should have %d  requests;but had %d",resiltSize,len(result.Items))
	//}

	//for i,city := range expectedCities {
	//	if result.Items[i].(string) != city {
	//		t.Errorf("expected city #%d: %s;but was %s",i,city,result.Items[i])
	//	}
	//}

}