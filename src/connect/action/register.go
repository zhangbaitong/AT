package action

import (
	"common"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type (
	Register struct{}

	Account struct {
		Acid        int
		Ac_name     string
		Ac_password string
		Email       string
		Mobile      string
		Status      int
		Create_time int
	}
)

const (
	INSERT string = "insert into account_tab (ac_name,ac_password,email,mobile,status,create_time) values (?,?,?,?,?,?)"
)

func NewRegister() *Register {
	return new(Register)
}

func (register *Register) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	acname := req.FormValue("acname")
	password := req.FormValue("password")
	account := Account{Ac_name: acname, Ac_password: password,
		Email: acname, Mobile: acname}
	ok := register_insert(&account)
	if ok {
		w.Write([]byte("0"))
	} else {
		w.Write([]byte("-1"))
	}

}

func register_insert(ac *Account) (ok bool) {
	tx, err := common.GetDB().Begin()
	if err != nil {
		fmt.Println(err)
		return false
	}
	stmt, err := tx.Prepare(INSERT)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(ac.Ac_name, ac.Ac_password, ac.Email, ac.Mobile, 0, time.Now().Unix())

	if err != nil {
		fmt.Println(err)
		return false
	}
	tx.Commit()
	return true
}
