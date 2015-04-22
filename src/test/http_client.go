package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestData struct {
	Version  string
	ServerIP string
	Port     int
	Method   string
	Params   string
}

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8080/v1/version")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func httpPost() {
	//var data map[string]string
	//data["imageName"]="centos:latest"
	//data["containerName"]="tomzhao"
	//post_data := RequestData{Version: "1.0", Method: "/auth/logout", Params: "{\"ac_name\":\"tomzhao\",\"ac_password\":\"111111\"}"}
	post_data := RequestData{Version: "1.0", Method: "/auth/getacid", Params: "{\"openid\":\"tomzhao\"}"}
	strPostData, _ := json.Marshal(post_data)
	strTemp := "request=" + string(strPostData)
	resp, err := http.Post("http://127.0.0.1:8080/auth/getacid",
		"application/x-www-form-urlencoded", strings.NewReader(strTemp))
	//"application/json",strings.NewReader(strTemp))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1:8080/v1/version",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/v1/version", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

//

func main() {
	//strImage,strTag:=GetImage("10.122.75.228:5000/centostest:latest")
	//fmt.Println("strImage=", strImage)
	//fmt.Println("strTag=", strTag)
	//return
	//httpGet()
	httpPost()
	//httpPostForm()
	//httpDo()
}
