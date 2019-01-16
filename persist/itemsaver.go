package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
	"crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (chan engine.Item,error){
	out := make(chan engine.Item)
	client,err:=elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil,err
	}

	go func(client *elastic.Client){
		itemCount := 0
		for{
			item := <- out
			log.Printf("【itemSaver】#%d: %v\n",itemCount,item)
			itemCount++
			err:=save(client,index,item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v:%v",item,err)
			}
		}
	}(client)
	return out,nil
}

func save(client *elastic.Client,index string,item engine.Item)(err error){

	if item.Type==""{
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != ""{
		indexService.Id(item.Id)
	}

	_,err = indexService.
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
