package request

import (
	"io"
	"net/http"
	"net/url"
)

// FileField 上传文件的结构体
type FileField struct {
	FieldName string
	FileName  string
	File      io.Reader
}

// BasicAuth http basic auth 结构体
type BasicAuth struct {
	Username string
	Password string
}

// Args 发送http请求所使用的参数结构
type Args struct {
	Client  *http.Client
	Headers http.Header
	Cookies map[string]string

	Body     io.Reader
	Params   url.Values
	FormData url.Values
	Files    []FileField
	JSON     interface{}

	BasicAuth BasicAuth

	Hooks []Hook
}

// NewArgs 返回一个新的 *Args
func NewArgs() *Args {
	return &Args{
		Client: &http.Client{
			Transport:     DefaultClient().Transport,
			CheckRedirect: DefaultClient().CheckRedirect,
			Jar:           newCookieJar(),
			Timeout:       DefaultClient().Timeout,
		},
		Headers:  newDefaultHeaders(),
		Cookies:  map[string]string{},
		Params:   url.Values{},
		FormData: url.Values{},
	}
}

// Reset 重置所有的参数
func (a *Args) Reset() {
	a.Headers = newDefaultHeaders()
	a.Cookies = map[string]string{}
	a.Params = url.Values{}
	a.FormData = url.Values{}
	a.BasicAuth = BasicAuth{}
	a.Files = nil
	a.Hooks = nil
	a.JSON = nil
	a.Body = nil
}

// setCookie 设置http cookie
func (a *Args) setCookie(key, value string) {
	a.Cookies[key] = value
}

// setBasicAuth 设置 http BasicAuth 需要的用户名和密码
func (a *Args) setBasicAuth(username, password string) {
	a.BasicAuth = BasicAuth{username, password}
}

// addFile 添加需要上传的文件
func (a *Args) addFile(file FileField) {
	a.Files = append(a.Files, file)
}

// addHook 添加请求处理钩子
func (a *Args) addHook(hook Hook) {
	a.Hooks = append(a.Hooks, hook)
}

// SetContentType 设置Content-Type header
// 如果不设置，默认使用 DefaultContentType
func (a *Args) setContentType(contentType IContentType) {
	a.Headers.Set("Content-Type", contentType.String())
}

// buildBodyReader 生成http请求所需要的 bodyReader
func (a *Args) buildBodyReader() (body io.Reader, err error) {
	if a.Body != nil {
		return a.Body, nil
	}
	if a.JSON != nil {
		return a.jsonReader()
	}
	if a.FormData != nil || a.Files != nil {
		return a.formReader()
	}

	return nil, nil
}

// buildRequest 根据Args生成 *http.Request
func (a *Args) buildRequest(method string, u string) (*http.Request, error) {
	if _, ok := a.Headers["Content-Type"]; !ok {
		a.setContentType(DefaultContentType)
	}

	body, err := a.buildBodyReader()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, buildURL(u, a.Params), body)
	if err != nil {
		return nil, err
	}

	req.Header = a.Headers
	a.writeCookies(req)

	if a.BasicAuth.Username != "" {
		req.SetBasicAuth(a.BasicAuth.Username, a.BasicAuth.Password)
	}
	return req, nil
}

// request 发送http请求
func (a *Args) request(method string, url string) (*Response, error) {
	req, err := a.buildRequest(method, url)
	if err != nil {
		return nil, err
	}

	// 调用 BeforeRequest hook
	res, err := a.applyBeforeReqHooks(req)
	if err != nil {
		return nil, err
	} else if res != nil {
		return &Response{res, nil}, err
	}

	res, err = a.Client.Do(req)

	// 调用 AfterRequest hook
	newResp, newErr := a.applyAfterReqHooks(req, res, err)
	if newErr != nil {
		err = newErr
	}
	if newResp != nil {
		res = newResp
	}

	return &Response{res, nil}, err
}
