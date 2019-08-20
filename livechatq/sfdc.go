package livechatq

import (
	"fmt"
	"net/http"
	"reflect"
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

// Hello : prints hello
func Hello() {
	fmt.Println("hello")
}

/*
	Query String builder:
	qs: queryString
	qm: queryMap
*/
func queryStringBuilder(qs *interface{}, qm map[string]string) {
	// Iterate query map
	// for key, value := range qm {

	// }
}

// Authorize SFDC Requests
func authorize(config SfdcConfig) *http.Response {
	URL := config.authorizerURL + "/services/oauth2/token"
	req, _ := http.NewRequest("GET", URL, nil)

	fmt.Println(reflect.TypeOf(req))

	queryString := req.URL.Query()
	queryString.Add("grant_type", "password")
	queryString.Add("client_id", config.clientID)
	queryString.Add("client_secret", config.clientSecret)
	queryString.Add("username", config.basicAuthEmail)
	queryString.Add("password", config.basicAuthPassword)

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
