package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type Response interface {
	Render()
}

type JSONResponse struct {
	HttpStatus int          `json:"-"`
	Context    *gin.Context `json:"-"`
	Data       interface{}  `json:"data"`
}

func (c *JSONResponse) Render() {
	c.Context.Render(c.HttpStatus, JSON{render.JSON{Data: c.Data}})
}

type ApiResponse struct {
	HttpStatus int          `json:"-"`
	Context    *gin.Context `json:"-"`
	Retcode    int          `json:"retcode"`
	Data       interface{}  `json:"data"`
	Msg        string       `json:"msg"`
}

func (c *ApiResponse) Render() {

	fmt.Println("------------------>")
	c.Context.Render(c.HttpStatus, JSON{render.JSON{Data: c}})
}

type RedirectResponse struct {
	HttpStatus int          `json:"-"`
	Context    *gin.Context `json:"-"`
	Location   string
}

func (c *RedirectResponse) Render() {
	c.Context.Redirect(c.HttpStatus, c.Location)
}

type StringResponse struct {
	HttpStatus int          `json:"-"`
	Context    *gin.Context `json:"-"`
	Name       string
	Data       []interface{}
}

func (c *StringResponse) Render() {
	c.Context.String(c.HttpStatus, c.Name, c.Data...)
}

type HTMLResponse struct {
	HttpStatus int          `json:"-"`
	Context    *gin.Context `json:"-"`
	Name       string
	Data       interface{}
}

func (c *HTMLResponse) Render() {
	c.Context.HTML(c.HttpStatus, c.Name, c.Data)
}
