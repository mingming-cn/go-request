package main

import (
	"fmt"
	"time"

	"github.com/mingming-cn/go-request"
)

func diff(req *request.Request) {
	url := "http://example.com:12345"
	start := time.Now()

	_, _ = req.Get(url)

	diff := time.Now().Sub(start)
	fmt.Println(diff.Seconds())
}

func main() {
	req := request.New()

	fmt.Println("default timeout")
	diff(req)

	timeout := time.Duration(1 * time.Second)
	req.SetTimeout(timeout)
	fmt.Printf("set timeout = %f seconds\n", timeout.Seconds())
	diff(req)

	// Or use req.Client
	req = request.New()
	req.Client.Timeout = timeout
	fmt.Printf("set timeout = %f seconds\n", timeout.Seconds())
	diff(req)
}
