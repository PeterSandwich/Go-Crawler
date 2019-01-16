package persist

import (
	"testing"
	"crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
	"encoding/json"
	"crawler/engine"
)

func TestSave(t *testing.T){
	expect_profile:=engine.Item{
		Url: "http://album.zhenai.com/u/15151515",
		Type: "zhenai",
		Id: "12154555",
		Payload: model.Profile{
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



	// Todo: Try to start up elasctic search
	// here using docker go client
	client,err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}

	const index = "dating_test"
	err =save(client,index,expect_profile)
	if err!= nil {
		panic(err)
	}

	resp,err:=client.Get().
		Index("dating_profile").Type(expect_profile.Type).
		Id(expect_profile.Id).Do(context.Background())
	if err!= nil{
		panic(err)
	}
	t.Logf("%s",*resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source,&actual)
	if err != nil {
		panic(err)
	}

	actualProfile,err := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expect_profile {
		t.Errorf("got %v expected %v\n",actual,expect_profile)
	}
}
