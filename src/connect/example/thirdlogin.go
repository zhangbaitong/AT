package main

// Open url in browser:
// http://localhost:8080/

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RangelReale/osin"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><head><meta charset=\"utf-8\" /><title>第三方应用首页</title></head><body>"))
		w.Write([]byte("<h2>第三方登陆第一步</h2>"))
		w.Write([]byte("<h4>在应用顶部导航栏加入登陆链接（请求平台authorize接口）,并传入应用标识client_id和redirect_uri（此URL必须 在 “应用登记的域名”下）</h4>"))
		w.Write([]byte(fmt.Sprintf("<a href=\"https://connect.funzhou.cn/oauth2/authorize?response_type=code&client_id=1234&state=xyz1&scope=everything&redirect_uri=%s\">登陆（code模式）</a><br/>", url.QueryEscape("http://localhost:8080/callback"))))
		w.Write([]byte("以下是应用的其他正常内容...<br/></body></html>"))
	})

	// Application destination - CODE
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		w.Write([]byte("<html><body>"))
		w.Write([]byte("<h2>第三方登陆第二步</h2>"))
		w.Write([]byte("<h4>平台在redirect_uri后附加授权code，并返回HTTP302</h4>"))
		code := r.Form.Get("code")
		w.Write([]byte(fmt.Sprintf("<h4>由浏览器跳转到当前页并把code=%s随请求一起传递给此页的http.Request对象</h4>", code)))
		w.Write([]byte("<h4>此页的通过http.Request对象，取得授权code后，在后台访问平台Token接口(需使用HTTPBasicAuth协议传入client_id和client_secret)</h4>"))
		// build access code url
		aurl := fmt.Sprintf("https://connect.funzhou.cn/oauth2/token?grant_type=authorization_code&client_id=1234&state=xyz1&redirect_uri=%s&code=%s",
			url.QueryEscape("http://localhost:8080/callback"), url.QueryEscape(code))
		w.Write([]byte("<h4>" + aurl + "</h4>"))
		w.Write([]byte("<h4>取得如下信息</h4>"))
		jr := make(map[string]interface{})
		err := DownloadAccessToken(aurl, &osin.BasicAuth{"1234", "aabbccdd"}, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}
		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}
		// show json access token
		token, ok := jr["access_token"]
		if ok {
			w.Write([]byte(fmt.Sprintf("授权令牌access_token: %s<br/>\n", token)))
		}
		if at, ok := jr["expires_in"]; ok {
			w.Write([]byte(fmt.Sprintf("有效期expires_in: %f秒<br/>\n", at)))
		}
		if at, ok := jr["refresh_token"]; ok {
			w.Write([]byte(fmt.Sprintf("续期刷新令牌refresh_token: %s<br/>\n", at)))
		}
		if at, ok := jr["scope"]; ok {
			w.Write([]byte(fmt.Sprintf("正式授予的权限scope: %s<br/>\n", at)))
		}
		w.Write([]byte(fmt.Sprintf("<br/><a href=./bindingAndlogin?access_token=%s>进入第三步（获取OpenID登陆)可和第二步合并</a><br/>", token)))
		defer w.Write([]byte("</body></html>"))
	})
	http.HandleFunc("/bindingAndlogin", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body>"))
		w.Write([]byte("<h2>第三方登陆第三步</h2>"))
		accessToken := r.Form.Get("access_token")
		w.Write([]byte(fmt.Sprintf("<h4>第三方使用access_token去平台请求https://connect.funzhou.cn/oauth2/me?access_token=%s</h4>", accessToken)))

		w.Write([]byte("<h4>获取当前登陆用户的openid（每个应用不一致）,创建（存在就不创建）并绑定第三方自己账号并登陆</h4>"))

		jr := make(map[string]interface{})
		err := CallInterface("https://connect.funzhou.cn/oauth2/me", accessToken, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}
		// show json access token
		clientId, _ := jr["client_id"]
		openId, _ := jr["openid"]
		w.Write([]byte(fmt.Sprintf("<h4>相当于登陆openId:%s的用户名是openid%s", clientId, openId)))
		w.Write([]byte(fmt.Sprintf("<h4>相当于登陆openId:%s的密码名是access_token：%s", clientId, accessToken)))

		defer w.Write([]byte("</body></html>"))
	})
	http.ListenAndServe(":8080", nil)
}

func DownloadAccessToken(url string, auth *osin.BasicAuth, output map[string]interface{}) error {
	// download access token
	preq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if auth != nil {
		preq.SetBasicAuth(auth.Username, auth.Password)
	}

	pclient := &http.Client{}
	presp, err := pclient.Do(preq)
	if err != nil {
		return err
	}

	if presp.StatusCode != 200 {
		return errors.New("Invalid status code")
	}
	jdec := json.NewDecoder(presp.Body)
	err = jdec.Decode(&output)
	return err
}

func CallInterface(url string, accessToken string, output map[string]interface{}) error {
	// download access token
	preq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	preq.Header.Add("access_token", accessToken)

	pclient := &http.Client{}
	presp, err := pclient.Do(preq)
	if err != nil {
		return err
	}

	if presp.StatusCode != 200 {
		return errors.New("Invalid status code")
	}
	jdec := json.NewDecoder(presp.Body)
	err = jdec.Decode(&output)
	return err
}
