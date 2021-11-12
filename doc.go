// Package request is a developer-friendly HTTP Send library for Gopher.
//
// GET Request:
//
// 	resp, err := Send.New().Get("https://httpbin.org/get")
// 	j, err := resp.JSON()
//
// POST Request:
//
// 	req = Send.New()
//	req.SetFormData("key","value")
//	req.SetFormData("a","123")
//	resp, err := req.Post("https://httpbin.org/post")
//
// Custom Cookies:
//
// 	req = Send.New()
//	req.SetCookie("key": "value")
//	req.SetCookie("a": "123")
//	resp, err := req.Get("https://httpbin.org/cookies")
//
//
// Custom Headers:
//
// 	req = Send.New()
//	req.SetHeader("Accept-Encoding", "gzip,deflate,sdch")
//	req.SetHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
//	resp, err := req.Get("https://httpbin.org/get")
//
// Upload Files:
//
// 	req = Send.New()
//	f, err := os.Open("test.txt")
//	req.AddFile(Send.FileField{"file", "test.txt", f})
//	resp, err := req.Post("https://httpbin.org/post")
//
// JSON Body:
//
// 	req = Send.New()
//	req.SetJSON(map[string]string{
//		"a": "A",
//		"b": "B",
//	})
//	resp, err := req.Post("https://httpbin.org/post")
//	req.SetJSON([]int{1, 2, 3})
//	resp, err = req.Post("https://httpbin.org/post")
//
// others Body:
//
// 	req = Send.New()
//	req.SetBody(strings.NewReader("<xml><a>abc</a></xml"))
//  req.SetContentType(Send.ApplicationXML)
//	resp, err := req.Post("https://httpbin.org/post")
//
//	// form
// 	req = Send.New()
//  // Default Content-Type is "application/x-www-form-urlencoded"
//	req.Body = strings.NewReader("a=1&b=2")
//	resp, err = req.Post("https://httpbin.org/post")
//
// HTTP Basic Authentication:
//
// 	req = Send.New()
//	req.SetBasicAuth("user", "passwd")
//	resp, err := req.Get("https://httpbin.org/basic-auth/user/passwd")
//
// Set Timeout:
//
// req = Send.New()
// req.SetTimeout(5 * time.Second)
// resp, err := req.Get("http://example.com:12345")
//
// Need more control?
//
// You can setup req.Client(you know, it's an &http.Client),
// for example: set timeout
package request
