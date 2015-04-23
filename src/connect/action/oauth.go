package action

import (
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/RangelReale/osin/example"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"oauth"
)

type OAuth struct {
	Server *osin.Server
}

func NewOAuth() *OAuth {

	sconfig := osin.NewServerConfig()
	sconfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	sconfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
	sconfig.AllowGetAccessRequest = true
	sconfig.AllowClientSecretInParams = true

	oauth := OAuth{
		Server: osin.NewServer(sconfig, oauth.NewATStorage()),
	}
	return &oauth
}

func (oauth *OAuth) Authorize(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := oauth.Server.NewResponse()
	defer resp.Close()

	if ar := oauth.Server.HandleAuthorizeRequest(resp, r); ar != nil {
		if !example.HandleLoginPage(ar, w, r) {
			return
		}
		ar.UserData = struct{ Login string }{Login: "test"}
		ar.Authorized = true
		oauth.Server.FinishAuthorizeRequest(resp, r, ar)
	}
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["custom_parameter"] = 187723
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
			if ar.Username == "test" && ar.Password == "test" {
				ar.Authorized = true
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
	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	if !resp.IsError {
		resp.Output["custom_parameter"] = 19923
	}
	osin.OutputJSON(resp, w, r)
}

func (oauth *OAuth) Info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := oauth.Server.NewResponse()
	defer resp.Close()

	if ir := oauth.Server.HandleInfoRequest(resp, r); ir != nil {
		oauth.Server.FinishInfoRequest(resp, r, ir)
	}
	osin.OutputJSON(resp, w, r)
}
