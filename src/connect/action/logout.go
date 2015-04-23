package action

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Logout struct {
}

func NewLogout() *Logout {
	return new(Logout)
}

func (l *Logout) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cookie := http.Cookie{Name: COOKIENAME, Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
}
