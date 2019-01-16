package parser

import (
	"testing"
	"io/ioutil"
	"log"
	"fmt"
)

func TestParseProFile(t *testing.T) {
	body,_:=ioutil.ReadFile("test_profile.html")

	result:=ParseProFile(body,"http://album.zhenai.com/u/15151515","test","男")
	if result.Items == nil {
		fmt.Println("error")
		return
	}
	val :=  result.Items[0]
	log.Println(val)

}
