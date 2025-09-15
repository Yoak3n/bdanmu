package config

import (
	"bdanmu/package/logger"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

const (
	userAgent    = `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36 Edg/97.0.1072.69`
	publicKeyPEM = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDLgd2OAkcGVtoE3ThUREbio0Eg
Uc/prcajMKXvkCKFCWhJYJcLkcM2DKKcSeFpD/j6Boy538YXnR6VhcuUJOhH2x71
nzPjfdTcqMz7djHum0qSZA0AyCBDABUqCrfNgCiJ00Ra7GmRj+YCK1NJEuewlb40
JNrRuoEUXpabUzGB8QIDAQAB
-----END PUBLIC KEY-----
`
)

func LoginFromFrontend() (string, string) {
	return getLoginUrl()
}

func checkCookieNeedRefresh() (bool, int64, error) {
	uri := "https://passport.bilibili.com/x/passport-login/web/cookie/info?csrf=" + getCsrf()
	client := http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", Conf.Auth.Cookie)
	res, err := client.Do(req)
	if err != nil {
		logger.Logger.Println("checkCookieNeedRefresh error:", err)
		return false, 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	// logger.Logger.Println(string(body))
	data := gjson.ParseBytes(body)
	if data.Get("code").Int() != 0 {
		return true, 0, err
	}
	if data.Get("data.refresh").Bool() {
		logger.Logger.Println("Need Refresh")
		return true, data.Get("data.timestamp").Int(), nil
	}
	return false, 0, nil
}

func CheckCookieValid() bool {
	logger.Logger.Println("检查cookie有效性")
	refresh, _, err := checkCookieNeedRefresh()
	logger.Logger.Debugln("checkCookieNeedRefresh", refresh, err)
	if err != nil {
		if err.Error() == "cookie有效" {
			return true
		} else {
			return false
		}
	}
	return !refresh
}

func RefreshCookie() error {
	// 获取 refresh_csrf
	refresh, _, err := checkCookieNeedRefresh()
	if err != nil {
		return err
	}
	if !refresh {
		return errors.New("cookie有效")
	}
	logger.Logger.Debugln("cookie过期，开始刷新")
	refreshCsrf, err := getRefreshCsrf()
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	// 获取新cookie
	newRefreshToken, err := refreshCookie(refreshCsrf)
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	// 确认更新
	logger.Logger.Debugln("commit cookie")
	err = commitCookie()
	if err != nil {
		return err
	}
	Conf.Auth.RefreshToken = newRefreshToken
	SetCookieRefresh()
	return nil
}

func getRefreshCsrf() (string, error) {
	uri := "https://www.bilibili.com/correspond/1/"
	correspondPath, err := getCorrespondPath(time.Now().UnixMilli())
	if err != nil {
		return "", err
	}
	uri += correspondPath
	client := http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", Conf.Auth.Cookie)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	refreshCsrf := dom.Find("#1-name").Text()
	return refreshCsrf, nil
}

func refreshCookie(refreshCsrf string) (string, error) {
	uri := "https://passport.bilibili.com/x/passport-login/web/cookie/refresh"
	postData := url.Values{}
	csrf := getCsrf()
	postData.Add("refresh_token", Conf.Auth.RefreshToken)
	postData.Add("source", "main_page")
	postData.Add("refresh_csrf", refreshCsrf)
	postData.Add("csrf", csrf)
	client := http.DefaultClient
	req, _ := http.NewRequest("POST", uri, strings.NewReader(postData.Encode()))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", Conf.Auth.Cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	data := gjson.ParseBytes(body)
	if data.Get("code").Int() != 0 {
		return "", errors.New(data.Get("message").String())
	}
	var cookieContent []byte
	cookie := make(map[string]string)
	for _, v := range res.Header["Set-Cookie"] {
		kv := strings.Split(v, ";")[0]
		kvArr := strings.Split(kv, "=")
		cookie[kvArr[0]] = kvArr[1]
	}
	cookieContent = []byte(`DedeUserID=` + cookie["DedeUserID"] + `;DedeUserID__ckMd5=` + cookie["DedeUserID__ckMd5"] + `;Expires=` + cookie["Expires"] + `;SESSDATA=` + cookie["SESSDATA"] + `;bili_jct=` + cookie["bili_jct"] + `;`)
	Conf.Auth.Cookie = string(cookieContent)
	newRefreshToken := data.Get("data.refresh_token").String()
	return newRefreshToken, nil
}

func getLoginKeyAndLoginUrl() (loginKey string, loginUrl string) {
	uri := "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
	client := http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	data := gjson.ParseBytes(body)
	loginKey = data.Get("data.qrcode_key").String()
	loginUrl = data.Get("data.url").String()
	return
}

func commitCookie() error {
	client := http.DefaultClient
	uri := "https://passport.bilibili.com/x/passport-login/web/confirm/refresh"
	postData := url.Values{}
	postData.Add("csrf", getCsrf())
	postData.Add("refresh_token", Conf.Auth.RefreshToken)
	req, _ := http.NewRequest("POST", uri, strings.NewReader(postData.Encode()))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Cookie", Conf.Auth.Cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	data := gjson.ParseBytes(body)
	if data.Get("code").Int() != 0 {
		return errors.New(data.Get("message").String())
	}
	return nil
}

func getCorrespondPath(ts int64) (string, error) {
	pubKeyBlock, _ := pem.Decode([]byte(publicKeyPEM))
	hash := sha256.New()
	random := rand.Reader
	msg := []byte(fmt.Sprintf("refresh_%d", ts))
	var pub *rsa.PublicKey
	pubInterface, parseErr := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if parseErr != nil {
		return "", parseErr
	}
	pub = pubInterface.(*rsa.PublicKey)
	encryptedData, encryptErr := rsa.EncryptOAEP(hash, random, pub, msg, nil)
	if encryptErr != nil {
		return "", encryptErr
	}
	return hex.EncodeToString(encryptedData), nil
}

func getCsrf() string {
	reg := regexp.MustCompile(`bili_jct=([0-9a-zA-Z]+);`)
	if result := reg.FindStringSubmatch(Conf.Auth.Cookie); len(result) > 1 {
		return result[1]
	}
	return ""
}
