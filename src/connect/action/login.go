package action

import (
	"common"
	"fmt"
	"github.com/dchest/authcookie"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"net/http"
	"time"
)

type (
	Login struct {
		View *render.Render
	}

	User struct {
		Acname   string
		Password string
	}
)

const (
	//cookie加密、解密使用
	KEY        string = "QAZWERT4556"
	COOKIENAME string = "MNBVCXZ"
)

func NewLogin() *Login {
	l := Login{
		View: render.New(),
	}
	return &l
}

func (l *Login) Get(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cookie, err := req.Cookie(COOKIENAME)
	if err == nil {
		login := authcookie.Login(cookie.Value, []byte(KEY))
		fmt.Println(login)
	}
	l.View.HTML(w, http.StatusOK, "login", nil)

}

func (l *Login) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	acname := req.FormValue("acname")
	password := req.FormValue("password")
	user := User{Acname: acname, Password: password}
	ok := login_query(&user)
	if ok {
		generateCookie(w, req, user.Acname, 1)
		w.Write([]byte("0"))
	} else {
		w.Write([]byte("-1"))
	}
}

//登录插入
func login_query(user *User) bool {
	strSQL := fmt.Sprintf("select count(ac_name) from account_tab where (ac_name='%s' or email='%s' or mobile='%s') and ac_password='%s'", user.Acname, user.Acname, user.Acname, user.Password)
	rows, err := common.GetDB().Query(strSQL)
	defer rows.Close()
	if err != nil {
		return false
	} else {
		var nCount int
		for rows.Next() {
			rows.Scan(&nCount)
		}

		if nCount == 0 {
			return false
		}
	}
	return true
}

//生成cookie，放到reponse对象
func generateCookie(w http.ResponseWriter, r *http.Request, userNmae string, number int) {
	timeLength := 24 * time.Hour
	cookieValue := authcookie.NewSinceNow(userNmae, timeLength, []byte(KEY))
	expire := time.Now().Add(timeLength)
	cookie := http.Cookie{Name: COOKIENAME, Value: cookieValue, Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(w, &cookie)
}
