package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaders(t *testing.T) {
	req := New()
	url := "https://httpbin.org/get"
	req.SetHeader("Foo", "Bar")
	resp, _ := req.Get(url)
	j, _ := resp.JSON()
	defer resp.Body.Close()
	v, _ := j.Get("headers").MustMap()["Foo"]
	assert.Equal(t, v, "Bar")
	v, _ = j.Get("headers").MustMap()["User-Agent"]
	assert.Equal(t, v, DefaultUserAgent)

	req.SetHeader("User-Agent", "Foobar")
	resp, _ = req.Get(url)
	j, _ = resp.JSON()
	defer resp.Body.Close()
	v, _ = j.Get("headers").MustMap()["User-Agent"]
	assert.Equal(t, v, "Foobar")
}
