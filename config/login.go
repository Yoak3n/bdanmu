package config

import (
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
	"time"
)

func getLoginUrl() (string, string) {
	loginKey, loginUrl := getLoginKeyAndLoginUrl()

	return loginUrl, loginKey
}

func IsLogin(cookie ...string) (bool, gjson.Result, string, string) {
	uri := "https://api.bilibili.com/x/web-interface/nav"
	csrf := getCsrf()
	cookieStr := Conf.Auth.Cookie
	client := http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", userAgent)
	if len(cookie) > 0 {
		req.Header.Set("Cookie", cookie[0])
	} else {
		req.Header.Set("Cookie", cookieStr)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	data := gjson.ParseBytes(body)
	return data.Get("code").Int() == 0, data, cookieStr, csrf
}

func VerifyLogin(loginKey string) {
	for {
		uri := "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
		client := http.Client{}
		uri += "?" + "qrcode_key=" + loginKey
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Set("User-Agent", userAgent)
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		data := gjson.ParseBytes(body)
		if data.Get("data.url").String() != "" {
			var cookieContent []byte
			cookie := make(map[string]string)
			for _, v := range resp.Header["Set-Cookie"] {
				kv := strings.Split(v, ";")[0]
				kvArr := strings.Split(kv, "=")
				cookie[kvArr[0]] = kvArr[1]
			}
			cookieContent = []byte(`DedeUserID=` + cookie["DedeUserID"] + `;DedeUserID__ckMd5=` + cookie["DedeUserID__ckMd5"] + `;Expires=` + cookie["Expires"] + `;SESSDATA=` + cookie["SESSDATA"] + `;bili_jct=` + cookie["bili_jct"] + `;`)
			Conf.Auth.Cookie = string(cookieContent)
			Conf.Auth.RefreshToken = data.Get("data.refresh_token").String()
			break
		}
		time.Sleep(time.Second * 3)
	}
}
