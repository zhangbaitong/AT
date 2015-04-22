package login

import (
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"net/http"
)

type Login struct {
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
	l.viewRender.HTML(w, http.StatusOK, "login", nil)
}
