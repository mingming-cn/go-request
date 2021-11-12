package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRedirectNoRedirect(t *testing.T) {
	req := New()
	url := "https://httpbin.org/get"
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), url)
}

// 由于 httpbin.org 暂停了Redirect测试接口，所以以下测试用例暂时注释掉
// see: https://github.com/postmanlabs/httpbin/issues/617
//
// func TestCheckRedirectNumberLessThanDefault(t *testing.T) {
// 	req := New()
// 	url := "https://httpbin.org/redirect/3"
// 	resp, _ := req.Get(url)
// 	u, _ := resp.URL()
// 	assert.Equal(t, u.String(), "https://httpbin.org/get")
// }
//
// func TestCheckRedirectNumberGreatThanDefault(t *testing.T) {
// 	req := New()
// 	url := "https://httpbin.org/redirect/15"
// 	resp, err := req.Get(url)
// 	assert.NotEqual(t, err, ErrMaxRedirect)
// 	u, err := resp.URL()
// 	assert.NotEqual(t, err, ErrMaxRedirect)
// 	assert.Equal(t, u.String(), "https://httpbin.org/relative-redirect/4")
// }
//
// func TestCheckRedirectWithHeaders(t *testing.T) {
// 	req := New()
// 	url := "https://httpbin.org/redirect/2"
// 	req.Headers.Set("Referer", "http://example.com")
// 	req.Headers.Set("X-Test", "test")
//
// 	resp, _ := req.Get(url)
// 	u, _ := resp.URL()
// 	assert.Equal(t, u.String(), "https://httpbin.org/get")
// 	assert.Equal(t, resp.Request.Header.Get("X-Test"), req.Headers.Get("X-Test"))
// 	assert.NotEqual(t, resp.Request.Header.Get("Referer"), req.Headers.Get("Referer"))
// 	assert.Equal(t, resp.Request.Header.Get("User-Agent"), DefaultUserAgent)
// }
//
// func TestCheckRedirectCustom(t *testing.T) {
// 	url := "https://httpbin.org/redirect/12"
// 	req := New()
// 	req.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
// 		if len(via) > 16 {
// 			return errors.New("redirect")
// 		}
// 		return nil
// 	}
// 	resp, _ := req.Get(url)
// 	u, _ := resp.URL()
// 	assert.Equal(t, u.String(), "https://httpbin.org/get")
// }
//
// func TestCheckRedirectChangeDefaultLimit(t *testing.T) {
// 	url := "https://httpbin.org/redirect/12"
// 	req := New()
// 	origin := DefaultRedirectLimit
// 	DefaultRedirectLimit = 16
// 	resp, _ := req.Get(url)
// 	u, _ := resp.URL()
// 	assert.Equal(t, u.String(), "https://httpbin.org/get")
// 	DefaultRedirectLimit = origin
// }
