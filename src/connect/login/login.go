package action

import (
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

type Register struct {
	viewRender *render.Render
}

func NewLogin() *Login {
	l := Login{
		viewRender: render.New(),
	}
	return &l
}

func (l *Login) GetLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	l.viewRender.HTML(w, http.StatusOK, "login", nil)
}
func (l *Login) PostLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//l.viewRender.HTML(w, http.StatusOK, "login", nil)
	//l.viewRender.Data(w, http.StatusOK, []byte("Login ok."))
	/*cookie, err := req.Cookie("id")
	if err == nil {
	    fmt.Fprintln(w, "Domain:", cookie.Domain)
	    fmt.Fprintln(w, "Expires:", cookie.Expires)
	    fmt.Fprintln(w, "Name:", cookie.Name)
	    fmt.Fprintln(w, "Value:", cookie.Value)
	}*/
	log.Println(req.FormValue("p"))
	log.Println(req.FormValue("u"))
	w.Write([]byte("Login ok."))

}
