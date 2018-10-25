package admin

import (
	"blog/core"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type CaptchaRequestJson struct {
	Ticket string `json:"ticket"`
}

var Captcha core.HandlerFunc = func(c *core.Context) core.Response {
	u, _ := url.Parse("https://ssl.captcha.qq.com/ticket/verify")
	q := u.Query()
	ticket := c.DefaultQuery("ticket", "")
	if ticket == "" {
		return c.Fail(201, "missing param tick")
	}
	q.Set("aid", "2070777383")
	q.Set("AppSecretKey", "0KIXIBVzzzzimk1KWeO8ycw**")
	q.Set("Ticket", ticket)
	log.Println("ticket:", ticket)
	randStr := c.DefaultQuery("randstr", "")
	if randStr == "" {
		return c.Fail(202, "missing param randstr")
	}
	q.Set("RandStr", randStr)
	q.Set("UserIP", c.ClientIP())
	log.Println("Ip:", c.ClientIP())
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	log.Println("url::", u.String())
	if err != nil {
		return c.Fail(204, err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Fail(205, err)
	}
	f := map[string]interface{}{}
	json.Unmarshal(result, &f)
	log.Println("ffff!!!", f)
	if f["response"] != 1 {
		return c.Fail(206, f["err_msg"])
	}
	resp.Body.Close()
	return c.Success("验证通过！")
}
