package request

import "net/http"

// DefaultUserAgent 定义默认的 User-Agent
var DefaultUserAgent = "go-request/" + Version

// DefaultHeaders 定义默认的 headers
var DefaultHeaders = map[string]string{
	"Connection":      "keep-alive",
	"Accept-Encoding": "gzip, deflate",
	"Accept":          "*/*",
	"User-Agent":      DefaultUserAgent,
}

// newDefaultHeaders 定义默认的 http headers
func newDefaultHeaders() http.Header {
	headers := http.Header{}
	for k, v := range DefaultHeaders {
		headers.Set(k, v)
	}

	return headers
}
