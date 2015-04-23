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
	login := action.NewLogin()
	logout := action.NewLogout()
	me := action.NewMe()

	router := httprouter.New()
	router.GET("/", SayHello)
	router.NotFound = http.FileServer(http.Dir("./static/public")).ServeHTTP

	router.POST("/register", register.Post)
	router.GET("/login", login.Get)
	router.POST("/login", login.Post)
	router.POST("/logout", logout.Post)
	router.GET("/me", me.Get)

	router.GET("/oauth2/authorize", oauth.Authorize)
	router.GET("/oauth2/Token", oauth.Token)
	router.GET("/oauth2/Info", oauth.Info)

	log.Fatal(http.ListenAndServeTLS(":443", "./static/pem/servercert.pem", "./static/pem/serverkey.pem", router))
}
