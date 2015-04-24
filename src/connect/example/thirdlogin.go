package main

// Open url in browser:
// http://localhost8080/app

import (
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/RangelReale/osin/example"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body>"))

		w.Write([]byte(fmt.Sprintf("<a href=\"https://connect.funzhou.cn/oauth2/authorize?response_type=code&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Code</a><br/>", url.QueryEscape("http://localhost8080/appauth/code"))))
		w.Write([]byte(fmt.Sprintf("<a href=\"https://connect.funzhou.cn/oauth2/authorize?response_type=token&client_id=1234&state=xyz&scope=everything&redirect_uri=%s\">Implict</a><br/>", url.QueryEscape("http://localhost8080/appauth/token"))))
		w.Write([]byte(fmt.Sprintf("<a href=\"/appauth/password\">Password</a><br/>")))
		w.Write([]byte(fmt.Sprintf("<a href=\"/appauth/client_credentials\">Client Credentials</a><br/>")))
		w.Write([]byte(fmt.Sprintf("<a href=\"/appauth/assertion\">Assertion</a><br/>")))

		w.Write([]byte("</body></html>"))
	})

	// Application destination - CODE
	http.HandleFunc("/appauth/code", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		code := r.Form.Get("code")

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - CODE<br/>"))
		defer w.Write([]byte("</body></html>"))

		if code == "" {
			w.Write([]byte("Nothing to do"))
			return
		}

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("https://connect.funzhou.cn/oauth2/token?grant_type=authorization_code&client_id=1234&state=xyz&redirect_uri=%s&code=%s",
			url.QueryEscape("http://localhost8080/appauth/code"), url.QueryEscape(code))

		// if parse, download and parse json
		if r.Form.Get("doparse") == "1" {
			err := example.DownloadAccessToken(fmt.Sprintf("http://localhost8080%s", aurl),
				&osin.BasicAuth{"1234", "aabbccdd"}, jr)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.Write([]byte("<br/>"))
			}
		}

		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		// output links
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Goto Token URL</a><br/>", aurl)))

		cururl := *r.URL
		curq := cururl.Query()
		curq.Add("doparse", "1")
		cururl.RawQuery = curq.Encode()
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Download Token</a><br/>", cururl.String())))

		if rt, ok := jr["refresh_token"]; ok {
			rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/appauth/info?code=%s", at)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
		}
	})

	// Application destination - TOKEN
	http.HandleFunc("/appauth/token", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - TOKEN<br/>"))

		w.Write([]byte("Response data in fragment - not acessible via server - Nothing to do"))

		w.Write([]byte("</body></html>"))
	})

	// Application destination - PASSWORD
	http.HandleFunc("/appauth/password", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - PASSWORD<br/>"))

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("https://connect.funzhou.cn/oauth2/token?grant_type=password&scope=everything&username=%s&password=%s",
			"test", "test")

		// download token
		err := example.DownloadAccessToken(fmt.Sprintf("http://localhost8080%s", aurl),
			&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}

		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		if rt, ok := jr["refresh_token"]; ok {
			rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/appauth/info?code=%s", at)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
		}

		w.Write([]byte("</body></html>"))
	})

	// Application destination - CLIENT_CREDENTIALS
	http.HandleFunc("/appauth/client_credentials", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - CLIENT CREDENTIALS<br/>"))

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("https://connect.funzhou.cn/oauth2/token?grant_type=client_credentials")

		// download token
		err := example.DownloadAccessToken(fmt.Sprintf("http://localhost8080%s", aurl),
			&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}

		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		if rt, ok := jr["refresh_token"]; ok {
			rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/appauth/info?code=%s", at)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
		}

		w.Write([]byte("</body></html>"))
	})

	// Application destination - ASSERTION
	http.HandleFunc("/appauth/assertion", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - ASSERTION<br/>"))

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("https://connect.funzhou.cn/oauth2/token?grant_type=assertion&assertion_type=urn:osin.example.complete&assertion=osin.data")

		// download token
		err := example.DownloadAccessToken(fmt.Sprintf("http://localhost8080%s", aurl),
			&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}

		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		if rt, ok := jr["refresh_token"]; ok {
			rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/appauth/info?code=%s", at)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
		}

		w.Write([]byte("</body></html>"))
	})

	// Application destination - REFRESH
	http.HandleFunc("/appauth/refresh", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - REFRESH<br/>"))
		defer w.Write([]byte("</body></html>"))

		code := r.Form.Get("code")

		if code == "" {
			w.Write([]byte("Nothing to do"))
			return
		}

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("https://connect.funzhou.cn/oauth2/token?grant_type=refresh_token&refresh_token=%s", url.QueryEscape(code))

		// download token
		err := example.DownloadAccessToken(fmt.Sprintf("http://localhost:8080%s", aurl),
			&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}

		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		if rt, ok := jr["refresh_token"]; ok {
			rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
		}

		if at, ok := jr["access_token"]; ok {
			rurl := fmt.Sprintf("/appauth/info?code=%s", at)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Info</a><br/>", rurl)))
		}
	})

	// Application destination - INFO
	http.HandleFunc("/appauth/info", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		w.Write([]byte("<html><body>"))
		w.Write([]byte("APP AUTH - INFO<br/>"))
		defer w.Write([]byte("</body></html>"))

		code := r.Form.Get("code")

		if code == "" {
			w.Write([]byte("Nothing to do"))
			return
		}

		jr := make(map[string]interface{})

		// build access code url
		aurl := fmt.Sprintf("/info?code=%s", url.QueryEscape(code))

		// download token
		err := example.DownloadAccessToken(fmt.Sprintf("http://localhost8080%s", aurl),
			&osin.BasicAuth{Username: "1234", Password: "aabbccdd"}, jr)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("<br/>"))
		}

		// show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))

		if rt, ok := jr["refresh_token"]; ok {
			rurl := fmt.Sprintf("/appauth/refresh?code=%s", rt)
			w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Refresh Token</a><br/>", rurl)))
		}
	})
	http.ListenAndServe(":8080", nil)
}
