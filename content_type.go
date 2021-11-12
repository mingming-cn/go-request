package request

import (
	"bytes"
)

type (
	// IContentType ContentType接口
	IContentType interface {
		String() string
		GetCharset() string
		GetBoundary() string
	}

	// ContentType IContentType实现
	ContentType struct {
		Type     string
		Charset  string
		Boundary string
	}
)

// NewContentType 生成ContentType
func NewContentType(contentType string) IContentType {
	return ContentType{Type: contentType}
}

// 常用的 Content-Type
var (
	TextXML                   = NewContentType("text/xml")
	TextHTML                  = NewContentType("text/html")
	TextPlain                 = NewContentType("text/plain")
	ApplicationXML            = NewContentType("application/xml")
	ApplicationJSON           = NewContentType("application/json")
	MultipartFormData         = NewContentType("multipart/form-data")
	ApplicationOctetStream    = NewContentType("application/octet-stream")
	ApplicationFormURLEncoded = NewContentType("application/x-www-form-urlencoded")
	DefaultContentType        = ApplicationFormURLEncoded
)

// String 实现Stringer接口，转为string类型
func (c ContentType) String() string {
	if c.Type == "" {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString(c.Type)
	if c.Charset != "" {
		buf.WriteString("; charset=")
		buf.WriteString(c.Charset)
	}
	if c.Boundary != "" {
		buf.WriteString("; boundary=")
		buf.WriteString(c.Boundary)
	}

	return buf.String()
}

// GetCharset 获取charset
func (c ContentType) GetCharset() string {
	return c.Charset
}

// GetBoundary 获取boundary
func (c ContentType) GetBoundary() string {
	return c.Boundary
}
