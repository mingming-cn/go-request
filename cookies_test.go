package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCookies(t *testing.T) {
	req := New()
	req.Cookies = map[string]string{
		"key": "value",
		"a":   "123",
	}
	resp, _ := req.Get("https://httpbin.org/cookies")
	d, _ := resp.JSON()
	defer resp.Body.Close()

	v := map[string]interface{}{
		"key": "value",
		"a":   "123",
	}
	assert.Equal(t, d.Get("cookies").MustMap(), v)
}
