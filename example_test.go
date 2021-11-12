package request_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/mingming-cn/go-request"
)

func ExampleRequest_Get() {
	req := request.New()
	url := "https://httpbin.org/get"
	resp, _ := req.Get(url)
	d, _ := resp.JSON()
	fmt.Println(resp.Ok())
	fmt.Println(d.Get("url").MustString())
	// Output:
	// true
	// https://httpbin.org/get
}

func ExampleRequest_Get_params() {
	req := request.New()
	req.Params.Add("a", "1")
	req.Params.Add("b", "2")
	url := "https://httpbin.org/get"
	resp, _ := req.Get(url)
	d, _ := resp.JSON()
	fmt.Println(d.Get("url").MustString())
	// Output:
	// https://httpbin.org/get?a=1&b=2
}

func ExampleRequest_Get_customHeaders() {
	req := request.New()
	req.Headers.Set("X-Abc", "abc")
	req.Headers.Set("User-Agent", "go-Send-test")
	url := "https://httpbin.org/get"
	resp, _ := req.Get(url)
	d, _ := resp.JSON()
	fmt.Println(d.Get("headers").Get("User-Agent").MustString())
	fmt.Println(d.Get("headers").Get("X-Abc").MustString())
	// Output:
	// go-Send-test
	// abc
}

func ExampleRequest_Post() {
	req := request.New()
	req.FormData.Add("a", "1")
	req.FormData.Add("b", "2")
	url := "https://httpbin.org/post"
	_, _ = req.Post(url)
}

func ExampleRequest_Get_cookies() {
	req := request.New()
	req.Cookies = map[string]string{
		"name": "value",
		"foo":  "bar",
	}
	url := "https://httpbin.org/cookies"
	_, _ = req.Get(url)
}

func ExampleRequest_Post_files() {
	req := request.New()
	f, _ := os.Open("test.txt")
	defer f.Close()
	req.Files = []request.FileField{
		{FieldName: "abc", FileName: "abc.txt", File: f},
	}
	url := "https://httpbin.org/post"
	_, _ = req.Post(url)
}

func ExampleRequest_Post_rawBody() {
	req := request.New()
	req.Body = strings.NewReader("a=1&b=2&foo=bar")
	req.SetContentType(request.DefaultContentType)
	url := "https://httpbin.org/post"
	_, _ = req.Post(url)
}
