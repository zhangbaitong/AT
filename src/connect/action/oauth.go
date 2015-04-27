package action

import (
	"common"
	"errors"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/dchest/authcookie"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"net/http"
	"oauth"
	"time"
)

type (
	OAuth struct {
		Server *osin.Server
		View   *render.Render
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

func NewOAuth() *OAuth {

	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
	sconfig.AllowGetAccessRequest = true
	sconfig.AllowClientSecretInParams = true

	oauth := OAuth{
		Server: osin.NewServer(sconfig, oauth.NewATStorage()),
		View:   render.New(),
	}
	return &oauth
}

func (oauth *OAuth) GetAuthorize(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	acname := oauth.Logged(w, r)
	if acname != "" {
		//已经登录，则返回页面，出现 授权按钮+权限列表
		oauth.View.HTML(w, http.StatusOK, "oauth", map[string]string{"AuthorizeDisplay": "block", "LoginDisplay": "none", "RequestURI": r.RequestURI})

	} else {
		//未登录，则返回页面，出现 用户名密码框+授权并登陆按钮+权限列表
		oauth.View.HTML(w, http.StatusOK, "oauth", map[string]string{"AuthorizeDisplay": "none", "LoginDisplay": "block", "RequestURI": r.RequestURI})
	}

}

func (oauth *OAuth) PostAuthorize(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	acname := oauth.Logged(w, r)
	if acname == "" {
		//使用提交的表单登陆
		acname, _ := oauth.Login(w, r)
		//登陆失败
		if acname == "" {
			//返回页面，出现 登陆失败提示，用户名密码框+授权并登陆按钮+权限列表
			oauth.View.HTML(w, http.StatusOK, "oauth", nil)
			return
		}
	}

	//用户登陆成功，并确认授权，则进行下一步,根据请求,发放code 或token
	resp := oauth.Server.NewResponse()
	defer resp.Close()
	ar := oauth.Server.HandleAuthorizeRequest(resp, r)
	if ar != nil {
		//发放code 或token ,附加到redirect_uri后，并跳转
		//存储acname，acid,rsid,clientid,clientSecret等必要信息
		ar.UserData = struct{ Acname string }{Acname: acname}
		ar.Authorized = true
		oauth.Server.FinishAuthorizeRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)
}

func (oauth *OAuth) Token(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := oauth.Server.NewResponse()
	defer resp.Close()

	if ar := oauth.Server.HandleAccessRequest(resp, r); ar != nil {
		switch ar.Type {
		case osin.AUTHORIZATION_CODE:
			ar.Authorized = true
		case osin.REFRESH_TOKEN:
			ar.Authorized = true
		case osin.PASSWORD:
			user := User{Acname: ar.Username, Password: ar.Password}
			ok := oauth.LoginQuery(&user)
			if ok {
				oauth.GenerateCookie(w, r, user.Acname, 1)
				ar.Authorized = true
			} else {
				//通过redirect_uri 返回错误约定 并跳转到改redirect_uri
			}
		case osin.CLIENT_CREDENTIALS:
			ar.Authorized = true
		case osin.ASSERTION:
			if ar.AssertionType == "urn:osin.example.complete" && ar.Assertion == "osin.data" {
				ar.Authorized = true
			}
		}
		oauth.Server.FinishAccessRequest(resp, r, ar)
	}
	osin.OutputJSON(resp, w, r)
}

func (oauth *OAuth) Logged(w http.ResponseWriter, req *http.Request) string {
	cookie, err := req.Cookie(COOKIENAME)
	if err == nil {
		return authcookie.Login(cookie.Value, []byte(KEY))
	}
	return ""
}

func (oauth *OAuth) Login(w http.ResponseWriter, req *http.Request) (string, error) {
	acname := req.FormValue("acname")
	password := req.FormValue("password")
	if acname == "" || password == "" {
		return "", errors.New("未输入用户名和密码！")
	}
	user := User{Acname: acname, Password: password}
	ok := oauth.LoginQuery(&user)
	if ok {
		oauth.GenerateCookie(w, req, user.Acname, 1)
		return acname, nil
	} else {
		return "", errors.New("用户名或密码错误！")
	}
}

//登录插入
func (oauth *OAuth) LoginQuery(user *User) bool {
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
func (oauth *OAuth) GenerateCookie(w http.ResponseWriter, r *http.Request, userNmae string, number int) {
	timeLength := 24 * time.Hour
	cookieValue := authcookie.NewSinceNow(userNmae, timeLength, []byte(KEY))
	expire := time.Now().Add(timeLength)
	cookie := http.Cookie{Name: COOKIENAME, Value: cookieValue, Path: "/", Expires: expire, MaxAge: 86400}
	http.SetCookie(w, &cookie)
}
