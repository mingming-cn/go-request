package request

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	req := New()
	req.FormData.Add("a", "A")
	req.FormData.Add("foo", "bar")
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	assert.Equal(t, d.Get("url").MustString(), url)
	assert.Equal(t, d.Get("form").MustMap(),
		map[string]interface{}{
			"a":   "A",
			"foo": "bar",
		}, true)
}

func TestPostFiles(t *testing.T) {
	req := New()
	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	f := []byte{
		'a',
		'b',
		'c',
		'd',
	}
	_, _ = w.Write(f)
	w.Flush()
	f2, _ := os.Open("doc.go")
	defer f2.Close()
	req.FormData.Add("key", "value")
	req.FormData.Add("a", "123")
	req.Files = []FileField{
		{"abc", "abc.txt", b},
		{"test", "test.txt", f2},
	}
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)
	d, _ := resp.JSON()
	defer resp.Body.Close()

	assert.Equal(t, resp.Ok(), true)
	v := map[string]interface{}{
		"key": "value",
		"a":   "123",
	}
	assert.Equal(t, d.Get("form").MustMap(), v)
	_, x := d.Get("files").CheckGet("abc")
	assert.Equal(t, x, true)
	_, x = d.Get("files").CheckGet("test")
	assert.Equal(t, x, true)
}

func TestPostRawBody(t *testing.T) {
	req := New()
	req.Body = strings.NewReader("a=1&b=2")
	req.SetContentType(DefaultContentType)
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)
	defer resp.Body.Close()

	j, _ := resp.JSON()
	assert.Equal(t, j.Get("form").MustMap(),
		map[string]interface{}{
			"a": "1",
			"b": "2",
		}, true)
}

func TestPostXML(t *testing.T) {
	req := New()
	xml := "<xml><a>abc</a></xml"
	req.Body = strings.NewReader(xml)
	req.SetContentType(ApplicationXML)
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)

	j, _ := resp.JSON()
	data, _ := j.Get("data").String()
	assert.Equal(t, data, xml)
}

func TestPostFormIO(t *testing.T) {
	req := New()
	req.Body = strings.NewReader("a=1&b=2")
	url := "https://httpbin.org/post"
	resp, err := req.Post(url)
	assert.Nil(t, err)

	j, err := resp.JSON()
	assert.Nil(t, err)
	a := j.Get("form").MustMap()
	assert.Equal(t, a,
		map[string]interface{}{
			"a": "1",
			"b": "2",
		}, true)
}

func TestPostFormFile(t *testing.T) {
	req := New()
	b := &bytes.Buffer{}
	w := bufio.NewWriter(b)
	f := []byte{'a', 'b', 'c', 'd'}
	_, _ = w.Write(f)
	w.Flush()

	req.FormData.Add("key", "value")
	req.FormData.Add("a", "123")
	req.Files = []FileField{
		{"abc", "abc.txt", b},
	}
	url := "https://httpbin.org/post"
	resp, _ := req.Post(url)
	d, _ := resp.JSON()

	v := map[string]interface{}{
		"key": "value",
		"a":   "123",
	}
	assert.Equal(t, d.Get("form").MustMap(), v)
	_, x := d.Get("files").CheckGet("abc")
	assert.Equal(t, x, true)
}
