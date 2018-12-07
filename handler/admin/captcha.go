package admin

import (
	"blog/config"
	"blog/core"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CaptchaRequestJson struct {
	Ticket string `json:"ticket"`
}

var Captcha core.HandlerFunc = func(c *core.Context) core.Response {
	u, _ := url.Parse(config.App.Captcha.Url)
	q := u.Query()
	ticket := c.DefaultQuery("ticket", "")
	if ticket == "" {
		return c.Fail(201, "missing param tick")
	}
	q.Set("aid", config.App.Captcha.Aid)
	q.Set("AppSecretKey", config.App.Captcha.AppSecretKey)
	q.Set("Ticket", ticket)
	randStr := c.DefaultQuery("randstr", "")
	if randStr == "" {
		return c.Fail(202, "missing param randstr")
	}
	q.Set("Randstr", randStr)
	q.Set("UserIP", c.ClientIP())
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return c.Fail(204, err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Fail(205, err)
	}
	f := map[string]interface{}{}
	json.Unmarshal(result, &f)
	if f["response"] != "1" {
		return c.Fail(206, f["err_msg"])
	}
	resp.Body.Close()
	return c.Success("验证通过！")
}
