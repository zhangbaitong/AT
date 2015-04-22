package main

import (
	"common"
	"net/http"
	"rcenter/resource"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func startServer() {
	common.Log().Println("Server is start at ", time.Now().String(), " , on port 8080")
	http.HandleFunc("/", resource.ResouceHandler)
	// err := http.ListenAndServeTLS(":8080", "../connect/static/pem/servercert.pem", "../connect/static/pem/serverkey.pem", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		common.Log().Println("Server start faild error:", err)
	}
}

func main() {
	startServer()
}
