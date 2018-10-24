package admin

import (
	"blog/core"
	"blog/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var Captcha core.HandlerFunc = func(c *core.Context) core.Response {
	u, _ := url.Parse("https://ssl.captcha.qq.com/ticket/verify")
	q := u.Query()
	q.Set("Aid", "2070777383")
	q.Set("AppSecretKey", "0KIXIBVzzzzimk1KWeO8ycw**")
	q.Set("Ticket", c.Query("ticket"))
	q.Set("RandStr", util.GetRandomString(18))
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return c.Fail(201, err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	f := map[string]interface{}{}
	json.Unmarshal(result, &f)
	log.Println("ffff!!!", f)
	if f["response"] != 1 {
		return c.Fail(205, err)
	}
	resp.Body.Close()
	if err != nil {
		return c.Fail(203, err)
	}
	return c.Success("验证通过！")
}
