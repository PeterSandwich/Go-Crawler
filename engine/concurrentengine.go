package engine

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}
type ConncurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

func (e *ConncurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i:=0;i<e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _,r :=range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _,item := range result.Items {
			go func(){e.ItemChan <- item }()
		}

		for _,request := range result.Requests {
			if isDUplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult,s ReadyNotifier){

	go func(){
		for{
			s.WorkerReady(in)
			r := <-in
			result,err := worker(r)
			if err!= nil{
				continue
			}
			out <- result
		}
	}()
}
var isExit = make(map[string]bool)
func isDUplicate(url string) bool {
	if _,ok := isExit[url];ok {
		return true
	}
	isExit[url]=true
	return false
}