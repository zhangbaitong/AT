package action

import (
	"common"
	_"github.com/dchest/authcookie"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"encoding/json"
	"log"
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

func getParams( r *http.Request) (params string) {
	strPostData := r.FormValue("request")
	fmt.Println("strPostData :", strPostData)
	var request common.RequestData

	err := json.Unmarshal([]byte(strPostData), &request)
	if err != nil {
		fmt.Println("json data decode faild :", err)
	}
	fmt.Println("request.Params :", request.Params)
	return request.Params
}

func setParams(strMethod string, code int,strMessgae string,strData string)(strbody []byte,err error){
	v1 := common.Response{Method: strMethod, Code: code, Messgae: strMessgae, Data: strData}	
	body, err := json.Marshal(v1)
	if err != nil {
		fmt.Println(err)
		return body, err
	}
	return body, nil
}

var logger *log.Logger

func init() {
	if logger==nil{
		logger = common.Log()
	}
}
//帐号注册
func RegisterAccount(w http.ResponseWriter, res http.Response, request common.RequestData) (code int, result string) {
	//把前台参数转换成结构体
	var account Account
	err := json.Unmarshal([]byte(request.Params), &account)
	if err != nil {
		logger.Println("json data decode faild :", err)
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
	return 0,"OK"
}

//登录
func Login(w http.ResponseWriter, res http.Response, request common.RequestData) (code int, result string) {

	//把前台参数转换成结构体
	var account Account
	err := json.Unmarshal([]byte(request.Params), &account)
	if err != nil {
		logger.Println("json data decode faild :", err)
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
	/*
	secret := []byte("secret")
	cookie := authcookie.NewSinceNow(account.Ac_name, 24*time.Hour, secret)
	http.SetCookie(w, cookie)
	*/
	return 0,"OK"
}

//查询账户是否存在
func isFieldExist(name string, value string) bool {
	return true
}

func Logout(w http.ResponseWriter, r *http.Request,ps httprouter.Params){
	strParams:=getParams(r);
	fmt.Fprint(w, "%s BYE BYE !\n",strParams)
}

func GetAcidByOpenid(w http.ResponseWriter, r *http.Request,ps httprouter.Params){
	strParams:=getParams(r);
	var openvalue map[string]interface{}
	err := json.Unmarshal([]byte(strParams), &openvalue)
	if err != nil {
		logger.Println("json data decode faild :", err)
		return 
	}
	common.DisplayJson(openvalue)
	strOpenid, ok := openvalue["openid"].(string)
	if !ok {
		fmt.Fprint(w, "openid error !\n")
		return 
	}

	strSQL:=fmt.Sprintf("select acid from openid_tab where openid='%s'",strOpenid)

	rows, err := common.GetDB().Query(strSQL)
	defer rows.Close()
	var strBody []byte
	if err != nil {
		strBody,_=setParams("/auth/getacid",1,"database error !","")
	} else {
		var nAcid int
		for rows.Next() {
			rows.Scan(&nAcid)
		}
		if nAcid==0 {
			strBody,_=setParams("/auth/getacid",1,"user acid not exist!","")
		} else {
			strData:=fmt.Sprintf( "{\"acid\":\"%d\"}",nAcid)
			strBody,_=setParams("/auth/getacid",0,"ok",strData)			
		}
	}	
	w.Write(strBody)	
}
