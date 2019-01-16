package engine



type SimpleEngine struct {}

func (e SimpleEngine)Run(seeds ...Request) {

	// 初始化种子请求
	var requests []Request
	for _,r := range seeds {
		requests = append(requests,r)
	}

	//对每个请求进行处理
	for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]

		parseResult,err:=worker(r)
		if err != nil {
			continue
		}
		requests=append(requests,parseResult.Requests...)
		//for _,item := range parseResult.Items {
			//log.Printf("Got item %s\n",item)
		//}
	}
}

