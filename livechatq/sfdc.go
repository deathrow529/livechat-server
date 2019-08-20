package livechatq

import (
	"fmt"
	"net/http"
	"net/url"
)

// SfdcConfig : SFDC App Configuration
type SfdcConfig struct {
	authorizerURL     string
	liveChatURL       string
	basicAuthEmail    string
	basicAuthPassword string
	clientID          string
	clientSecret      string
}

// QueryStringBuilder : Build query string in url values
func queryStringBuilder(qs *url.Values, qm map[string]string) {
	// Iterate query map
	for key, value := range qm {
		(*qs).Add(key, value)
	}
}

// Authorize SFDC Requests
func authorize(config SfdcConfig) *http.Response {
	URL := config.authorizerURL + "/services/oauth2/token"
	req, _ := http.NewRequest("GET", URL, nil)

	queryString := req.URL.Query()
	qmap := map[string]string{
		"grant_type":    "password",
		"client_id":     config.clientID,
		"client_secret": config.clientSecret,
		"username":      config.basicAuthEmail,
		"password":      config.basicAuthPassword,
	}

	queryStringBuilder(&queryString, qmap)
	req.URL.RawQuery = queryString.Encode()
	resp, err := http.Get(req.URL.String())
	if err != nil {
		fmt.Println("Authorize Fail")
	}
	return resp
}

// CreateSession : create livechat session
func CreateSession(config SfdcConfig) {
	// URL := config.liveChatURL + "/chat/rest/System/SessionId"
	// payload := map[string]string{
	// 	"grant_type":    "password",
	// 	"client_id":     config.clientID,
	// 	"client_secret": config.clientSecret,
	// 	"username":      config.basicAuthEmail,
	// 	"password":      config.basicAuthPassword,
	// }
	// jsonValue, _ := json.Marshal(payload)
}
