package request

import (
	"bdanmu/config"
	"bdanmu/package/logger"
	"io"
	"net/http"
	"net/url"
)

func Get(urlStr string, args ...string) ([]byte, error) {
	params := "?"
	if l := len(args); l > 0 {
		for i := 0; i < l; i++ {
			params += args[i]
		}
	} else {
		params = ""
	}
	client := &http.Client{}
	parsed, err := url.Parse(config.Conf.Proxy)
	if err == nil && parsed.String() != "" {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(parsed)}
	}
	logger.Logger.Println("request url:", urlStr+params)
	req, _ := http.NewRequest(http.MethodGet, urlStr+params, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	req.Header.Set("Cookie", config.Conf.Auth.Cookie)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func GetWithWbi(urlStr string) ([]byte, error) {
	newUrlStr, err := signAndGenerateURL(urlStr)
	if err != nil {
		logger.Logger.Errorln(err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, newUrlStr, nil)

	if err != nil {
		logger.Logger.Errorln(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0")
	req.Header.Set("Cookie", config.Conf.Auth.Cookie)
	client := &http.Client{}
	parsed, err := url.Parse(config.Conf.Proxy)
	if err == nil && parsed.String() != "" {
		client.Transport = &http.Transport{Proxy: http.ProxyURL(parsed)}
	}
	response, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorln(err)
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Logger.Errorln(err)
		return nil, err
	}
	response.Body.Close()
	return body, err

}
