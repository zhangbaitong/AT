package main

import (
	"connect/login"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello"))
}
func main() {
	l := login.NewLogin()
	router := httprouter.New()
	router.GET("/", SayHello)
	router.NotFound = http.FileServer(http.Dir("./static/public")).ServeHTTP

	router.GET("/login", l.GetLogin)
	router.POST("/login", l.PostLogin)
	router.GET("/register", SayHello)
	router.GET("/oauth2", SayHello)

	log.Fatal(http.ListenAndServeTLS(":443", "./static/pem/servercert.pem", "./static/pem/serverkey.pem", router))
}
