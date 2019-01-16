package main

import (
	"net/http"
	"crawler/frontend/controller"
)

func main(){
	http.Handle("/",http.FileServer(
		http.Dir(`E:\GoProjects\src\crawler\frontend\view`)))
	http.Handle("/search",controller.CreateSearchResultHandler(`E:\GoProjects\src\crawler\frontend\view\template.html`))
	err:=http.ListenAndServe(":9999",nil)
	if err != nil {
		panic(err)
	}
}