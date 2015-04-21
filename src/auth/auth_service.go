package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

type HostSwitch map[string]http.Handler

func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if handler := hs[r.Host]; handler != nil {
        handler.ServeHTTP(w, r)
    } else {
        http.Error(w, "Forbidden", 403) // Or Redirect?
    }
}

func new_router()(*httprouter.Router) {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
    return router
}

func main() {
	fmt.Println("Server is start at ", time.Now().String(), " , on port 8080")
	router:=new_router();
	hs := make(HostSwitch)
	hs["127.0.0.1:8080"] = router

	// Use the HostSwitch to listen and serve on port 12345
	//log.Fatal(http.ListenAndServe(":8080", hs))
	log.Fatal(http.ListenAndServeTLS(":8080", "servercert.pem", "serverkey.pem", hs))
}
