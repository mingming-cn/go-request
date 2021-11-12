go-request
------------
Advanced HTTP client for golang.


Installation
------------

```
go get github.com/mingming-cn/go-request
```


Usage
-------

```go
import (
    "github.com/mingming-cn/go-request"
)
```

**GET**:

```go
req := request.New()
resp, err := req.Get("https://httpbin.org/get")
j, err := resp.JSON()
```

**POST**:

```go
req := request.New()
req.FormData.Set("key", "value")
req.FormData.Set("a", "123")
resp, err := req.Post("https://httpbin.org/post")
```

**Cookies**:

```go
req := request.New()
req.Cookies.Set("key", "value")
req.Cookies.Set("a", "123")
resp, err := req.Get("https://httpbin.org/cookies")
```

**Headers**:

```go
req := request.New()
req.Headers.Set("Accept-Encoding", "gzip,deflate,sdch",)
req.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",)
resp, err := req.Get("https://httpbin.org/get")
```

**Files**:

```go
req := request.New()
f, err := os.Open("test.txt")
req.AddFile(request.FileField{"file", "test.txt", f})
resp, err := req.Post("https://httpbin.org/post")
```

**Json**:

```go
req := request.New()
req.SetJSON(map[string]string{
    "a": "A",
    "b": "B",
})
resp, err := req.Post("https://httpbin.org/post")
req.SetJSON([]int{1, 2, 3})
resp, err = req.Post("https://httpbin.org/post")
```


**HTTP Basic Authentication**:
```go
req := request.New()
req.SetBasicAuth("user", "passwd")
resp, err := req.Get("https://httpbin.org/basic-auth/user/passwd")
```