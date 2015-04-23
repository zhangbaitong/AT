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
	router.GET("/me/:access_token", me.Get)

	log.Fatal(http.ListenAndServeTLS(":443", "./static/pem/servercert.pem", "./static/pem/serverkey.pem", router))
}
