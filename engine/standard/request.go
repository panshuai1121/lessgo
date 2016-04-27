package standard

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/lessgo/lessgo/engine"
	"github.com/lessgo/lessgo/logs"
)

type (
	// Request implements `engine.Request`.
	Request struct {
		*http.Request
		url    engine.URL
		header engine.Header
		logger logs.Logger
	}
)

var _ engine.Request = new(Request)

// IsTLS implements `engine.Request#TLS` function.
func (r *Request) IsTLS() bool {
	return r.Request.TLS != nil
}

// Scheme implements `engine.Request#Scheme` function.
func (r *Request) Scheme() string {
	if r.IsTLS() {
		return "https"
	}
	return "http"
}

// Host implements `engine.Request#Host` function.
func (r *Request) Host() string {
	return r.Request.Host
}

// URL implements `engine.Request#URL` function.
func (r *Request) URL() engine.URL {
	return r.url
}

// Header implements `engine.Request#URL` function.
func (r *Request) Header() engine.Header {
	return r.header
}

// Cookies parses and returns the HTTP cookies sent with the request.
func (r *Request) Cookies() []*http.Cookie {
	return r.Request.Cookies()
}

// Cookie returns the named cookie provided in the request or
// ErrNoCookie if not found.
func (r *Request) Cookie(name string) (*http.Cookie, error) {
	return r.Request.Cookie(name)
}

// AddCookie adds a cookie to the request.  Per RFC 6265 section 5.4,
// AddCookie does not attach more than one Cookie header field.  That
// means all cookies, if any, are written into the same line,
// separated by semicolon.
func (r *Request) AddCookie(c *http.Cookie) {
	r.Request.AddCookie(c)
}

// func Proto() string {
// 	return r.request.Proto()
// }
//
// func ProtoMajor() int {
// 	return r.request.ProtoMajor()
// }
//
// func ProtoMinor() int {
// 	return r.request.ProtoMinor()
// }

// ContentLength implements `engine.Request#ContentLength` function.
func (r *Request) ContentLength() int {
	return int(r.Request.ContentLength)
}

// UserAgent implements `engine.Request#UserAgent` function.
func (r *Request) UserAgent() string {
	return r.Request.UserAgent()
}

// RemoteAddress implements `engine.Request#RemoteAddress` function.
func (r *Request) RemoteAddress() string {
	return r.RemoteAddr
}

// Method implements `engine.Request#Method` function.
func (r *Request) Method() string {
	return r.Request.Method
}

// SetMethod implements `engine.Request#SetMethod` function.
func (r *Request) SetMethod(method string) {
	r.Request.Method = method
}

// URI implements `engine.Request#URI` function.
func (r *Request) URI() string {
	return r.RequestURI
}

// SetURI implements `engine.Request#SetURI` function.
func (r *Request) SetURI(uri string) {
	r.RequestURI = uri
}

// Body implements `engine.Request#Body` function.
func (r *Request) Body() io.Reader {
	return r.Request.Body
}

// FormValue implements `engine.Request#FormValue` function.
func (r *Request) FormValue(name string) string {
	return r.Request.FormValue(name)
}

// FormParams implements `engine.Request#FormParams` function.
func (r *Request) FormParams() map[string][]string {
	if err := r.ParseForm(); err != nil {
		r.logger.Error("%v", err)
	}
	return map[string][]string(r.Request.PostForm)
}

// FormFile implements `engine.Request#FormFile` function.
func (r *Request) FormFile(name string) (*multipart.FileHeader, error) {
	_, fh, err := r.Request.FormFile(name)
	return fh, err
}

// MultipartForm implements `engine.Request#MultipartForm` function.
func (r *Request) MultipartForm() (*multipart.Form, error) {
	err := r.ParseMultipartForm(engine.MaxMemory)
	return r.Request.MultipartForm, err
}

func (r *Request) reset(rq *http.Request, h engine.Header, u engine.URL) {
	r.Request = rq
	r.header = h
	r.url = u
}
