package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentType_String(t *testing.T) {
	assert.Equal(t, "application/x-www-form-urlencoded", ApplicationFormURLEncoded.String())
}

func TestContentType_String_WithoutCharset(t *testing.T) {
	contentType := ContentType{
		Type: "application/x-www-form-urlencoded",
	}

	assert.Equal(t, "application/x-www-form-urlencoded", contentType.String())
}
