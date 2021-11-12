package request

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseURLNormal(t *testing.T) {
	req := New()
	url := "https://httpbin.org/get"
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), url)
}

func TestResponseURLWithRedirect(t *testing.T) {
	req := New()
	url := "https://httpbin.org/redirect/3"
	resp, _ := req.Get(url)
	u, _ := resp.URL()
	assert.Equal(t, u.String(), "https://httpbin.org/get")
}

func TestResponseURLWithDisableRedirect(t *testing.T) {
	req := New()
	req.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("redirect")
	}
	resp, _ := req.Get("https://httpbin.org/redirect/3")
	u, _ := resp.URL()
	assert.Equal(t, u.String(), "https://httpbin.org/relative-redirect/2")
}

func TestResponseGzip(t *testing.T) {
	req := New()
	url := "https://httpbin.org/gzip"
	resp, _ := req.Get(url)
	d, _ := resp.JSON()
	t2, _ := resp.Text()
	c2, _ := resp.Content()
	defer resp.Body.Close()

	assert.Equal(t, resp.Reason() != "", true)
	assert.Equal(t, resp.OK(), true)
	assert.Equal(t, t2 != "", true)
	assert.Equal(t, c2 != nil, true)
	assert.Equal(t, d.Get("gzipped").MustBool(), true)
}

func TestResponseDeflate(t *testing.T) {
	req := New()
	url := "https://httpbin.org/deflate"
	resp, _ := req.Get(url)
	d, _ := resp.JSON()
	t2, _ := resp.Text()
	c2, _ := resp.Content()
	defer resp.Body.Close()

	assert.Equal(t, resp.Reason() != "", true)
	assert.Equal(t, resp.OK(), true)
	assert.Equal(t, t2 != "", true)
	assert.Equal(t, c2 != nil, true)
	assert.Equal(t, d.Get("deflated").MustBool(), true)
}
