package request

import (
	"io"
	"net/http"
	"time"
)

// Version 版本号
const Version = "0.0.1"

var (
	// DefaultTimeout 默认超时时间
	DefaultTimeout = time.Second * 10

	// DefaultTransport 默认的 transport
	DefaultTransport http.RoundTripper = transport()

	// DefaultClient 默认的HTTP Client
	DefaultClient = newDefaultClient
)

func newDefaultClient() *http.Client {
	return &http.Client{
		Timeout:       DefaultTimeout,
		Transport:     DefaultTransport,
		Jar:           newCookieJar(),
		CheckRedirect: defaultCheckRedirect,
	}
}

// Request Args的别名，主要对外提供调用方法
type Request struct {
	*Args
}

// New 返回*Request
func New() *Request {
	return &Request{NewArgs()}
}

// Send 发送HTTP请求到指定的URL
func Send(method, url string, args *Args) (*Response, error) {
	if args == nil {
		args = NewArgs()
	}

	return args.request(method, url)
}

// Get 发送HTTP GET 请求到指定的URL
func Get(url string, args *Args) (*Response, error) {
	return Send("GET", url, args)
}

// Get 发送HTTP GET 请求到指定的URL
func (req *Request) Get(url string) (*Response, error) {
	return req.request("GET", url)
}

// Head 发送HTTP HEAD 请求到指定的URL
func Head(url string, args *Args) (*Response, error) {
	return Send("HEAD", url, args)
}

// Head 发送HTTP HEAD 请求到指定的URL
func (req *Request) Head(url string) (*Response, error) {
	return req.request("HEAD", url)
}

// Post 发送HTTP POST 请求到指定的URL
func Post(url string, args *Args) (*Response, error) {
	return Send("POST", url, args)
}

// Post 发送HTTP POST 请求到指定的URL
func (req *Request) Post(url string) (*Response, error) {
	return req.request("POST", url)
}

// Put 发送HTTP PUT 请求到指定的URL
func Put(url string, args *Args) (*Response, error) {
	return Send("PUT", url, args)
}

// Put 发送HTTP PUT 请求到指定的URL
func (req *Request) Put(url string) (*Response, error) {
	return req.request("PUT", url)
}

// Patch 发送HTTP PATCH 请求到指定的URL
func Patch(url string, args *Args) (*Response, error) {
	return Send("PATCH", url, args)
}

// Patch 发送HTTP PATCH 请求到指定的URL
func (req *Request) Patch(url string) (*Response, error) {
	return req.request("PATCH", url)
}

// Delete 发送HTTP DELETE 请求到指定的URL
func Delete(url string, args *Args) (*Response, error) {
	return Send("DELETE", url, args)
}

// Delete 发送HTTP DELETE 请求到指定的URL
func (req *Request) Delete(url string) (*Response, error) {
	return req.request("DELETE", url)
}

// Options 发送HTTP OPTIONS 请求到指定的URL
func Options(url string, args *Args) (*Response, error) {
	return Send("OPTIONS", url, args)
}

// Options 发送HTTP OPTIONS 请求到指定的URL
func (req *Request) Options(url string) (*Response, error) {
	return req.request("OPTIONS", url)
}

// SetTimeout 设置HTTP请求超时时间
func (req *Request) SetTimeout(t time.Duration) *Request {
	req.Client.Timeout = t
	return req
}

// Reset 重置所有参数
func (req *Request) Reset() *Request {
	req.Args.Reset()
	return req
}

// SetBasicAuth 设置 http BasicAuth需要的用户名和密码
func (req *Request) SetBasicAuth(username, password string) *Request {
	req.Args.setBasicAuth(username, password)
	return req
}

// AddFile 添加需要上传的文件
func (req *Request) AddFile(file FileField) *Request {
	req.Args.addFile(file)
	return req
}

// AddHook 添加请求处理钩子
func (req *Request) AddHook(hook Hook) *Request {
	req.Args.addHook(hook)
	return req
}

// AddHooks 批量添加请求处理钩子
func (req *Request) AddHooks(hooks ...Hook) *Request {
	for _, hook := range hooks {
		req.Args.addHook(hook)
	}
	return req
}

// SetContentType 设置 Content-Type
// 如果不设置，默认使用 DefaultContentType
func (req *Request) SetContentType(contentType IContentType) *Request {
	req.Args.setContentType(contentType)
	return req
}

// SetCookie 设置 Cookie
func (req *Request) SetCookie(key, value string) *Request {
	req.Args.setCookie(key, value)
	return req
}

// SetHeader 设置 Header
// 如果不设置，将只使用DefaultHeaders中定义的header
func (req *Request) SetHeader(key, value string) *Request {
	req.Args.Headers.Set(key, value)
	return req
}

// SetBody 设置 Body Reader
func (req *Request) SetBody(body io.Reader) *Request {
	req.Args.Body = body
	return req
}

// SetJSON 设置 JSON Object
func (req *Request) SetJSON(json interface{}) *Request {
	req.Args.JSON = json
	return req
}

// AddParam 添加查询参数
// 多次添加相同的参数，将在发送请求中发送多个相同的参数
func (req *Request) AddParam(key, value string) *Request {
	req.Args.Params.Add(key, value)
	return req
}

// SetParam 设置查询参数
// 如果key已经存在将会覆盖旧的数据，如果key不存在将会添加一对新的key value
func (req *Request) SetParam(key, value string) *Request {
	req.Args.Params.Set(key, value)
	return req
}

// AddFormData 添加Form参数
// 多次添加相同的参数，将在发送请求中发送多个相同的参数
func (req *Request) AddFormData(key, value string) *Request {
	req.Args.FormData.Add(key, value)
	return req
}

// SetFormData 设置Form参数
// 如果key已经存在将会覆盖旧的数据，如果key不存在将会添加一对新的key value
func (req *Request) SetFormData(key, value string) *Request {
	req.Args.FormData.Set(key, value)
	return req
}
