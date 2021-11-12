package request

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	req := New()
	resp, _ := req.Get("https://httpbin.org/get")
	assert.Equal(t, resp.Ok(), true)

	url := "https://httpbin.org/get"
	resp, _ = req.Get(url)
	d, _ := resp.JSON()
	t2, _ := resp.Text()
	c2, _ := resp.Content()
	defer resp.Body.Close()

	assert.Equal(t, resp.Reason() != "", true)
	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, resp.OK(), true)
	assert.Equal(t, t2 != "", true)
	assert.Equal(t, c2 != nil, true)
	assert.Equal(t, d.Get("url").MustString(), url)

}

func TestGetParams(t *testing.T) {
	url := "https://httpbin.org/get"
	resp, _ := New().AddParam("foo", "bar").AddParam("a", "1").Get(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("args").MustMap(),
		map[string]interface{}{
			"foo": "bar",
			"a":   "1",
		})
}

func TestGetParams2(t *testing.T) {
	DefaultTransport = RoundTripFunc(func(req *http.Request) *http.Response {
		resp := httptest.NewRecorder()
		if req.Method != http.MethodGet {
			resp.Code = http.StatusBadRequest
		}
		req.ParseForm()
		reqForm := make(map[string]string)
		for k, v := range req.Form {
			if len(v) > 0 {
				reqForm[k] = v[0]
			}
		}
		reqFormBytes, _ := json.Marshal(map[string]interface{}{"args": reqForm})
		resp.Body = bytes.NewBuffer(reqFormBytes)

		return resp.Result()
	})

	url := "https://httpbin.org/get?ab=cd"
	resp, _ := New().AddParam("foo", "bar").AddParam("a", "1").Get(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("args").MustMap(),
		map[string]interface{}{
			"ab":  "cd",
			"foo": "bar",
			"a":   "1",
		})
}

func TestHead(t *testing.T) {
	DefaultTransport = RoundTripFunc(func(req *http.Request) *http.Response {
		resp := httptest.NewRecorder()
		if req.Method != http.MethodHead {
			resp.Code = http.StatusBadRequest
		}
		return resp.Result()
	})

	url := "https://httpbin.org/get"
	resp, err := New().Head(url)
	assert.Equal(t, nil, err)

	assert.Equal(t, resp.Ok(), true)
	content, err := resp.Content()
	assert.Equal(t, nil, err)
	assert.Equal(t, content, []byte{})
}

func TestPut(t *testing.T) {
	DefaultTransport = RoundTripFunc(func(req *http.Request) *http.Response {
		resp := httptest.NewRecorder()
		if req.Method != http.MethodPut {
			resp.Code = http.StatusBadRequest
		}
		resp.Body = bytes.NewBufferString(`{"url": "https://httpbin.org/put"}`)
		return resp.Result()
	})

	req := New()
	url := "https://httpbin.org/put"
	resp, _ := req.Put(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestDelete(t *testing.T) {
	DefaultTransport = RoundTripFunc(func(req *http.Request) *http.Response {
		resp := httptest.NewRecorder()
		if req.Method != http.MethodDelete {
			resp.Code = http.StatusBadRequest
		}
		resp.Body = bytes.NewBufferString(`{"url": "https://httpbin.org/delete"}`)
		return resp.Result()
	})

	url := "https://httpbin.org/delete"
	resp, _ := New().Delete(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestPatch(t *testing.T) {
	DefaultTransport = RoundTripFunc(func(req *http.Request) *http.Response {
		resp := httptest.NewRecorder()
		if req.Method != http.MethodPatch {
			resp.Code = http.StatusBadRequest
		}
		resp.Body = bytes.NewBufferString(`{"url": "https://httpbin.org/patch"}`)
		return resp.Result()
	})

	req := New()
	url := "https://httpbin.org/patch"
	resp, _ := req.Patch(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
}

func TestOptions(t *testing.T) {
	DefaultTransport = RoundTripFunc(func(req *http.Request) *http.Response {
		resp := httptest.NewRecorder()
		if req.Method != http.MethodOptions {
			resp.Code = http.StatusBadRequest
		}
		return resp.Result()
	})

	req := New()
	url := "https://httpbin.org/get"
	resp, _ := req.Options(url)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}

func TestPostJson(t *testing.T) {
	req := New()
	req.JSON = []int{1, 2, 3}
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v := []interface{}{
		json.Number("1"),
		json.Number("2"),
		json.Number("3"),
	}
	assert.Equal(t, d.Get("json").MustArray(), v)
}

func TestPostJson2(t *testing.T) {
	req := New()
	req.JSON = map[string]string{
		"a":   "b",
		"foo": "bar",
	}
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v := map[string]interface{}{
		"a":   "b",
		"foo": "bar",
	}
	assert.Equal(t, d.Get("json").MustMap(), v)
}

func TestPostJson3(t *testing.T) {
	req := New()
	type j struct {
		A string            `json:"a"`
		B map[string]string `json:"b"`
		C []string          `json:"c"`
		D []int             `json:"d"`
		E int               `json:"e"`
	}
	d := j{
		A: "hello",
		B: map[string]string{
			"a": "A",
			"b": "B",
			"c": "C",
		},
		C: []string{"lala", "aaaa"},
		D: []int{1, 2, 3},
		E: 5,
	}
	url := "https://httpbin.org/post"
	resp, _ := req.SetJSON(d).Post(url)
	j2, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v, _ := simplejson.NewJson([]byte(`{
		"a": "hello",
		"b": {"a": "A", "b":"B", "c":"C"},
		"c": ["lala", "aaaa"],
		"d": [1, 2, 3],
		"e": 5
	}`))
	assert.Equal(t, j2.Get("json"), v)
}

func TestBasicAuth(t *testing.T) {
	req := New()
	req.BasicAuth = BasicAuth{"user", "passwd"}
	url := "https://httpbin.org/basic-auth/user/passwd"
	resp, _ := req.Get(url)
	defer resp.Body.Close()
	assert.Equal(t, resp.OK(), true)

	req.BasicAuth = BasicAuth{
		Username: "user2",
		Password: "passwd2",
	}
	url = "https://httpbin.org/basic-auth/user2/passwd2"
	resp, _ = req.Get(url)
	defer resp.Body.Close()
	assert.Equal(t, resp.OK(), true)
}

func TestReset(t *testing.T) {
	req := New()
	req.BasicAuth = BasicAuth{"user", "passwd"}
	url := "https://httpbin.org"
	req.Get(url)

	req.Reset()
	assert.Equal(t, req.BasicAuth, BasicAuth{})
}

func TestGetArgsNil(t *testing.T) {
	url := "https://httpbin.org/get"
	resp, _ := New().Get(url)
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
}
