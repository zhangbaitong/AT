package action

import (
	"github.com/julienschmidt/httprouter"
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

}
