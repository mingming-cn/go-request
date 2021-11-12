package request

import "net/http"

// RoundTripFunc convert func to RoundTripFunc of implement RoundTripper interface
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip  implement RoundTripper interface
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
