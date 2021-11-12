package request

import (
	"net/url"
	"strings"
)

func buildURL(u string, params url.Values) string {
	if len(params) == 0 {
		return u
	}

	if strings.Contains(u, "?") {
		return u + "&" + params.Encode()
	}
	return u + "?" + params.Encode()
}
