package main

import (
	"connect/action"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello"))
}
func main() {

	oauth := action.NewOAuth()
	register := action.NewRegister()
	logout := action.NewLogout()
	me := action.NewMe()

	router := httprouter.New()
	router.GET("/", SayHello)
	router.NotFound = http.FileServer(http.Dir("./static/public")).ServeHTTP

	//Step1：获取Authorization Code
	router.GET("/oauth2/authorize", oauth.GetAuthorize)
	router.POST("/oauth2/authorize", oauth.PostAuthorize)
	//Step2：通过Authorization Code获取Access Token
	//Step3：（可选）权限自动续期，获取Access Token
	router.GET("/oauth2/token", oauth.Token)
	//Step4:通过Access Token获取用户OpenID
	router.GET("/oauth2/me", me.Get)

	router.GET("/oauth2/logout", logout.Get)
	router.POST("/oauth2/logout", logout.Get)
	router.POST("/oauth2/register", register.Post)

	log.Fatal(http.ListenAndServeTLS(":443", "./static/pem/servercert.pem", "./static/pem/serverkey.pem", router))
}
