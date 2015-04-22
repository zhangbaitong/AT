package main

import (
	"common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	post_data := common.RequestData{Version: "1.0", Method: "res/reg", Params: ""}
	strPostData, _ := json.Marshal(post_data)
	strTemp := "request=" + string(strPostData)
	fmt.Println(strTemp)
	resp, err := http.Post("http://127.0.0.1:8080/res/create",
		"application/x-www-form-urlencoded", strings.NewReader(strTemp))
	// "application/json", strings.NewReader(strTemp))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}
