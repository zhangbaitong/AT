package action

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Me struct {
}

func NewMe() *Me {
	return new(Me)
}

func (m *Me) Get(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//获取当前登陆用户的openID,同http://wiki.connect.qq.com/%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7openid_oauth2-0
	//ps.ByName("access_token")
	//1.access_token转换出 client_id->rsid
	//2.cookie 提出出acname->acid
	//rsid+acid 去查询 openid 不存在则新建立，最后返回

	cookie, err := req.Cookie(COOKIENAME)
	login := authcookie.Login(cookie.Value, []byte(KEY))
	if err != nil {
		fmt.Println("me.go cookie decryption error")
	}
	acid := getAcid(login)
	client_id := "TEST"
	openId := getOpenId(client_id, acid)
	result := "client_id=" + client_id + "&openid=" + openId
	w.Write([]byte(result))
}

func getAcId(acName string) int {
	strSQL := fmt.Sprintf("select acid from account_tab where ac_name='%s' limit 1", acName)
	rows, err := common.GetDB().Query(strSQL)
	defer rows.Close()
	if err != nil {
		return ""
	} else {
		var acid int
		for rows.Next() {
			rows.Scan(&acid)
		}

		return acid
	}
}

func getOpenId(clientId string, acid int) string {
	strSQL := fmt.Sprintf("select openid from openid_tab where res_id='%s' and acid=%d limit 1", clientId, acid)
	rows, err := common.GetDB().Query(strSQL)
	defer rows.Close()
	if err != nil {
		return ""
	} else {
		var openId int
		for rows.Next() {
			rows.Scan(&openId)
		}

		return openId
	}
}
