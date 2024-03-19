package config

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func login() (string, string) {
	for {
		loginKey, loginUrl := getLoginKeyAndLoginUrl()
		fmt.Println(loginKey, loginUrl)
		fp, err := os.OpenFile("qrcode.png", os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		var png []byte
		png, err = qrcode.Encode(loginUrl, qrcode.Medium, 256)
		if err != nil {
			panic(err)
		}
		_, err = fp.Write(png)
		if err != nil {
			panic(err)
		}
		fp.Close()
		verifyLogin(loginKey)
		logged, data, cookieStr, csrf := isLogin()
		if logged {
			_ = os.Remove("qrcode.png")
			uname := data.Get("data.uname").String()
			log.Println(uname + "已登录")
			return cookieStr, csrf
		}

	}
}
func isLogin() (bool, gjson.Result, string, string) {
	uri := "https://api.bilibili.com/x/web-interface/nav"
	csrf := getCsrf()
	cookieStr := Conf.Auth.Cookie
	client := http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", cookieStr)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	data := gjson.ParseBytes(body)
	return data.Get("code").Int() == 0, data, cookieStr, csrf
}

func verifyLogin(loginKey string) {
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
