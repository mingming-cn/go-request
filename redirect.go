package request

import (
	"errors"
	"net/http"
)

// DefaultRedirectLimit 最大重定向次数
var DefaultRedirectLimit = 10

// ErrMaxRedirect 当重定向次数大于DefaultRedirectLimit是将返回这个错误
var ErrMaxRedirect = errors.New("exceeded max redirects")

func defaultCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) > DefaultRedirectLimit {
		return ErrMaxRedirect
	}
	if len(via) == 0 {
		return nil
	}
	// Redirect requests with the first Header
	for key, val := range via[0].Header {
		// Don't copy Referer Header
		if key != "Referer" {
			req.Header[key] = val
		}
	}
	return nil
}
