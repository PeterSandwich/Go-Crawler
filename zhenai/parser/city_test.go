package parser

import (
	"testing"
	"io/ioutil"
	"log"
)

func TestParseCity(t *testing.T) {
	contents,err:=ioutil.ReadFile("test_city.html")
	if err != nil {
		panic(err)
	}
	result:=ParseCity(contents,"")
	for _,v:= range result.Requests {
		log.Println("URL: ",v.Url)
	}
	//for _,v:=range result.Items {
	//	log.Println("User: ",v)
	//}
}
