package core

import (
	"github.com/gin-gonic/gin"
	"github.com/verystar/golib/convert"
)

func getHttpStatus(c *Context, status int) int {
	if c.HttpStatus == 0 {
		return status
	}
	return c.HttpStatus
}

type Context struct {
	*gin.Context
	HttpStatus  int
	SessionData *SessionData
	Response    Response
}

func (c *Context) Status(status int) {
	c.HttpStatus = status
}

func (c *Context) Fail(code int, msg interface{}) Response {
	var message string
	if m, ok := msg.(error); ok {
		message = m.Error()
	} else {
		message = convert.ToStr(msg)
	}

	return &ApiResponse{
		HttpStatus: getHttpStatus(c, 200),
		Context:    c.Context,
		Retcode:    code,
		Msg:        message,
	}
}

func (c *Context) Success(data interface{}) Response {
	return &ApiResponse{
		HttpStatus: getHttpStatus(c, 200),
		Context:    c.Context,
		Retcode:    0,
		Data:       data,
		Msg:        "ok",
	}
}

func (c *Context) GetToken() string {
	return c.Request.Header.Get("Access-Token")
}
