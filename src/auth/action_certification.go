package action

import (
	"common"
	"github.com/dchest/authcookie"
	"net/http"
	"time"
)

type Account struct {
	Acid        int
	Ac_name     string
	Ac_password string
	Email       string
	Mobile      string
	Status      int
	Create_time int
}

//帐号注册
func RegisterAccount(w http.ResponseWriter, res http.Response, request common.RequestData) (code int, result string) {
	//把前台参数转换成结构体
	var account Account
	err := json.Unmarshal([]byte(request.Params), &account)
	if err != nil {
		logger.Println("json data decode faild :", err_json)
		return 1, "json data decode faild"
	}

	//参数校验
	if account.Ac_name == "" {
		logger.Println("action_certification：ac_name can't be empty")
		return 1, "ac_name can't be empty"
	}
	if account.Ac_password == "" {
		logger.Println("action_certification：ac_password can't be empty")
		return 1, "ac_password can't be empty"
	}
	if account.Email == "" {
		logger.Println("action_certification：email can't be empty")
		return 1, "ac_email can't be empty"
	}
	if account.Mobile == "" {
		logger.Println("action_certification：mobile can't be empty")
		return 1, "mobile can't be empty"
	}

	//校验账户、邮箱、手机号码是否已存在
	if true == isFieldExist("ac_name", account.Ac_name) {
		return 1, "ac_name is already exist"
	}
	if true == isFieldExist("email", account.Email) {
		return 1, "email is already exist"
	}
	if true == isFieldExist("mobile", account.Mobile) {
		return 1, "mobile is already exist"
	}

}

//登录
func Login(w http.ResponseWriter, res http.Response, request common.RequestData) (code int, result string) {

	//把前台参数转换成结构体
	var account Account
	err := json.Unmarshal([]byte(request.Params), &account)
	if err != nil {
		logger.Println("json data decode faild :", err_json)
		return 1, "json data decode faild"
	}

	//参数校验
	if account.Ac_name == "" {
		logger.Println("action_certification：ac_name can't be empty")
		return 1, "ac_name can't be empty"
	}
	if account.Ac_password == "" {
		logger.Println("action_certification：ac_password can't be empty")
		return 1, "ac_password can't be empty"
	}

	//生成cookie，放到reponse对象中
	secret := []byte("secret")
	cookie := authcookie.NewSinceNow(account.Ac_name, 24*time.Hour, secret)
	http.SetCookie(w, cookie)
}

//查询账户是否存在
func isFieldExist(name string, value string) boolean {

}
