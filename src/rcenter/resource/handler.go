package resource

import (
	"common"
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	common.Log().Println("This is resource handler ... ")
}

func ResouceHandler(w http.ResponseWriter, r *http.Request) {

	strMethod := r.URL.Path[len("/res/"):]
	if err := r.ParseForm(); err != nil {
		common.Log().Println("Server internal error:", err)
	}

	var request common.RequestData
	var res Resource

	if r.Method == "GET" {
	} else {
		strPostData := r.FormValue("request")
		common.Log().Println("strPostData :", strPostData)

		err := json.Unmarshal([]byte(strPostData), &request)
		if err != nil {
			common.Log().Println("json data decode faild :", err)
		}

		//get post resource object
		err_json := json.Unmarshal([]byte(request.Params), &res)
		if err_json != nil {
			common.Log().Println("json data decode faild :", err_json)
		}
		common.Log().Println("RES Handler : ", res)
	}

	var ret string = "ok"
	var code int = 0
	switch strMethod {
	case "reg":
		{
			Insert(res.res_name, res.owner_acid, res.operator_acid)
		}
	case "stop":
		{
			Update(res.res_id, res.res_name, res.owner_acid, res.operator_acid, 0)
		}
	case "get":
		{
			sqlstr := fmt.Sprintf(SELECT_BY_ID, res.res_id)
			resRet, _ := Query(sqlstr)
			resRetJSON, _ := json.Marshal(resRet)
			ret = string(resRetJSON)
		}
	case "update":
		{
			Update(res.res_id, res.res_name, res.owner_acid, res.operator_acid, res.status)
		}
	}
	ResMessage := common.Response{Method: strMethod, Code: code, Messgae: "ok", Data: ret}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body, err1 := json.Marshal(ResMessage)
	if err1 != nil {
		common.Log().Println(err1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
