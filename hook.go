package request

import "net/http"

// Hook ...
type Hook interface {
	// BeforeRequest
	// 发送HTTP请求之前将会调用 BeforeRequest,
	// 如果 resp != nil or err != nil
	// 将使用这里的 resp and err, 不再继续发送HTTP请求.
	BeforeRequest(req *http.Request) (resp *http.Response, err error)
	// AfterRequest
	// 获取到response之后将调用 AfterRequest
	// 如果 newResp != nil or newErr != nil
	// 将使用新的NewResp而不是原始响应.
	AfterRequest(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error)
}

// DefaultReqHooks 默认的Hook
// 比Args中的Hooks先执行
var DefaultReqHooks []Hook

func (a *Args) applyBeforeReqHooks(req *http.Request) (resp *http.Response, err error) {
	for _, hook := range DefaultReqHooks {
		resp, err = hook.BeforeRequest(req)
		if resp != nil || err != nil {
			return
		}
	}

	for _, hook := range a.Hooks {
		resp, err = hook.BeforeRequest(req)
		if resp != nil || err != nil {
			return
		}
	}
	return
}

func (a *Args) applyAfterReqHooks(req *http.Request, resp *http.Response, err error) (newResp *http.Response, newErr error) {
	for _, hook := range DefaultReqHooks {
		newResp, newErr = hook.AfterRequest(req, resp, err)
		if newResp != nil || newErr != nil {
			return
		}
	}

	for _, hook := range a.Hooks {
		newResp, newErr = hook.AfterRequest(req, resp, err)
		if newResp != nil || newErr != nil {
			return
		}
	}
	return
}
