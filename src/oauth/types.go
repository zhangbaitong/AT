package oauth

import (
	"github.com/RangelReale/osin"
	"net/http"
	"time"
)

// AuthorizeRequestType is the type for OAuth param `response_type`
type AuthorizeRequestType string

const (
	CODE  AuthorizeRequestType = "code"
	TOKEN                      = "token"
)

// Authorize request information
type AuthorizeRequest struct {
	Type        AuthorizeRequestType
	Client      osin.DefaultClient
	Scope       string
	RedirectUri string
	State       string

	// Set if request is authorized
	Authorized bool

	// Token expiration in seconds. Change if different from default.
	// If type = TOKEN, this expiration will be for the ACCESS token.
	Expiration int32

	// Data to be passed to storage. Not used by the library.
	UserData interface{}

	// HttpRequest *http.Request for special use
	HttpRequest *http.Request
}

// Authorization data
type AuthorizeData struct {
	// Client information
	Client osin.DefaultClient

	// Authorization code
	Code string

	// Token expiration in seconds
	ExpiresIn int32

	// Requested scope
	Scope string

	// Redirect Uri from request
	RedirectUri string

	// State data from request
	State string

	// Date created
	CreatedAt time.Time

	// Data to be passed to storage. Not used by the library.
	UserData interface{}
}

func (old *AuthorizeData) transfer() *osin.AuthorizeData {
	var authorizeData osin.AuthorizeData = osin.AuthorizeData{}
	authorizeData.Client = &old.Client
	authorizeData.Code = old.Code
	authorizeData.ExpiresIn = old.ExpiresIn
	authorizeData.Scope = old.Scope
	authorizeData.RedirectUri = old.RedirectUri
	authorizeData.State = old.State
	authorizeData.CreatedAt = old.CreatedAt
	authorizeData.UserData = old.UserData
	return &authorizeData
}
