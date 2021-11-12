package request

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
)

// Response ...
type Response struct {
	*http.Response
	content []byte
}

// JSON 返回 simplejson.JSON 类型的Body数据
func (resp *Response) JSON() (*simplejson.Json, error) {
	b, err := resp.Content()
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(b)
}

// JSONUnmarshal 解析JSON格式的数据到destPtr指针中
// destPtr 只能传指针
func (resp *Response) JSONUnmarshal(destPtr interface{}) error {
	reader, err := resp.Reader()
	if err != nil {
		return err
	}

	return jsoniter.NewDecoder(reader).Decode(destPtr)
}

// Reader 获取body reader
func (resp *Response) Reader() (reader io.ReadCloser, err error) {
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
	case "deflate":
		reader, err = zlib.NewReader(resp.Body)
	default:
		return resp.Body, nil
	}

	if err == io.EOF {
		return ioutil.NopCloser(bytes.NewReader([]byte{})), nil
	} else if err != nil {
		return nil, err
	}

	return reader, nil
}

// Content 返回[]byte格式的Body数据
func (resp *Response) Content() (b []byte, err error) {
	if resp.content != nil {
		return resp.content, nil
	}

	var reader io.ReadCloser
	if reader, err = resp.Reader(); err != nil {
		return nil, err
	}

	defer reader.Close()
	if b, err = ioutil.ReadAll(reader); err != nil {
		return nil, err
	}

	resp.content = b
	return b, err
}

// Text 获取文本格式的Body数据
func (resp *Response) Text() (string, error) {
	b, err := resp.Content()
	s := string(b)
	return s, err
}

// OK check Response StatusCode < 400 ?
func (resp *Response) OK() bool {
	return resp.StatusCode < 400
}

// Ok check Response StatusCode < 400 ?
func (resp *Response) Ok() bool {
	return resp.OK()
}

// Reason return Response Status
func (resp *Response) Reason() string {
	return resp.Status
}

// URL return finally Send url
func (resp *Response) URL() (*url.URL, error) {
	u := resp.Request.URL
	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound,
		http.StatusSeeOther, http.StatusTemporaryRedirect:
		location, err := resp.Location()
		if err != nil {
			return nil, err
		}
		u = u.ResolveReference(location)
	}
	return u, nil
}
