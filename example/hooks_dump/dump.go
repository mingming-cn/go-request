package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/mingming-cn/go-request"
)

type dumpHook struct {
}

func (d *dumpHook) BeforeRequest(req *http.Request) (resp *http.Response, err error) {
	dump, _ := httputil.DumpRequestOut(req, true)
	fmt.Println(string(dump) + "\n")
	return
}
func (d *dumpHook) AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	dump, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(dump))
	return
}

func main() {
	hook := &dumpHook{}
	req := request.New()
	req.Hooks = []request.Hook{hook}
	req.JSON = []int{1, 2, 3}
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)
	defer func() {
		_ = resp.Body.Close()
	}()
}
