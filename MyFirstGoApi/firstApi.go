package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
)

func main() {
	webservice := new(restful.WebService)
	webservice.Route(webservice.GET("show").To(showData))
	webservice.Route(webservice.GET("test").To(test))
	restful.Add(webservice)

	err := http.ListenAndServe(":8006", nil)
	if err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}

func showData(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "Jacky show data")
}

func test(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "It's a test")
}
