package livechatq

import (
	"net/http"

	"github.com/agila/livechat-server-go/util"
)

/*
payload := map[string]string{
	"grant_type":    "password",
	"client_id":     config.clientID,
	"client_secret": config.clientSecret,
	"username":      config.basicAuthEmail,
	"password":      config.basicAuthPassword,
}
jsonValue, _ := json.Marshal(payload)
*/

// SfdcConfig : SFDC App Configuration
type SfdcConfig struct {
	authorizerURL     string
	liveChatURL       string
	basicAuthEmail    string
	basicAuthPassword string
	clientID          string
	clientSecret      string
}

// NewSfdcConfig : Create SFDC Config Object
func NewSfdcConfig(authorizerURL string, liveChatURL string,
	basicAuthEmail string, basicAuthPassword, clientID string,
	clientSecret string) SfdcConfig {
	return SfdcConfig{
		authorizerURL,
		liveChatURL,
		basicAuthEmail,
		basicAuthPassword,
		clientID,
		clientSecret,
	}
}

// Authorize SFDC Requests
func authorize(config SfdcConfig) *http.Response {
	URL := config.authorizerURL + "/services/oauth2/token"
	qmap := map[string]string{
		"grant_type":    "password",
		"client_id":     config.clientID,
		"client_secret": config.clientSecret,
		"username":      config.basicAuthEmail,
		"password":      config.basicAuthPassword,
	}
	resp := util.Get(URL, qmap, nil)
	return resp
}

// CreateSession : create livechat session
func CreateSession(config SfdcConfig) *http.Response {
	URL := config.liveChatURL + "/chat/rest/System/SessionId"
}
