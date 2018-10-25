package admin

import (
	"blog/core"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
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
		return c.Fail(203, "missing param tick")
	}
	q.Set("Aid", "2070777383")
	q.Set("AppSecretKey", "0KIXIBVzzzzimk1KWeO8ycw**")
	q.Set("Ticket", ticket)
	log.Println("ticket:", ticket)
	randStr := c.DefaultQuery("randstr", "")
	if randStr == "" {
		return c.Fail(203, "missing param randstr")
	}
	q.Set("RandStr", randStr)
	ip := GetIntranetIp()
	if ip == "" {
		return c.Fail(204, "get ip failed")
	}
	q.Set("UserIP", ip)
	log.Println("Ip:", ip)
	resp, err := http.Get(u.String())
	if err != nil {
		return c.Fail(201, err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Fail(203, err)
	}
	f := map[string]interface{}{}
	json.Unmarshal(result, &f)
	log.Println("ffff!!!", f)
	if f["response"] != 1 {
		return c.Fail(205, f["err_msg"])
	}
	resp.Body.Close()
	return c.Success("验证通过！")
}

func GetIntranetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
