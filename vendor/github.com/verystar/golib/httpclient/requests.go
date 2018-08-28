package httpclient

import (
	"io"
	"net/http"
	"time"
)

const (
	FORM_URLENCODED = "application/x-www-form-urlencoded"
	JSON            = "application/json"
	XML             = "application/xml"
)

type IHook interface {
	Befor(req *http.Request)
	After(req *http.Request,rep *http.Response,err error)
}

type HttpClient struct {
	*http.Client
	Hook IHook
}

func NewClient(options ...func(*HttpClient)) *HttpClient {
	httpclient := &HttpClient{
		&http.Client{},
		nil,
	}
	httpclient.Client.Timeout = 5 * time.Second
	// Apply options in the parameters to request.
	for _, option := range options {
		option(httpclient)
	}

	return httpclient
}

func (h *HttpClient) Post(url string, contentType string, body io.Reader, options ... func(r *http.Request)) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)

	for _, option := range options {
		option(req)
	}
	return h.Do(req)
}

func (h *HttpClient) Get(url string, options ... func(r *http.Request)) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		option(req)
	}
	return h.Do(req)
}

func (h *HttpClient) Do(req *http.Request) (resp *http.Response, err error) {
	if h.Hook != nil{
		h.Hook.Befor(req)
	}

	rep, err := h.Client.Do(req)

	if h.Hook != nil {
		h.Hook.After(req, rep, err)
	}

	return rep, err
}