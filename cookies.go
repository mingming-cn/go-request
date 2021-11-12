package request

import (
	"net/http"
	"net/http/cookiejar"
)

// writeCookies 把Args中的cookies写入到http.Request中
func (a *Args) writeCookies(req *http.Request) {
	if a.Cookies == nil {
		return
	}
	cookies := a.Client.Jar.Cookies(req.URL)
	for k, v := range a.Cookies {
		cookies = append(cookies, &http.Cookie{Name: k, Value: v})
	}
	a.Client.Jar.SetCookies(req.URL, cookies)
}

// newCookieJar 新建一个空的cookieJar
func newCookieJar() *cookiejar.Jar {
	jar, _ := cookiejar.New(nil)
	return jar
}
