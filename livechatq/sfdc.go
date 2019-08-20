package livechatq

import (
	"net/http"

	"github.com/agila/livechat-server-go/util/requests"
)

// SfdcConfig : SFDC App Configuration
type SfdcConfig struct {
	authorizerURL     string
	liveChatURL       string
	basicAuthEmail    string
	basicAuthPassword string
	clientID          string
	clientSecret      string
	organizationID    string
	deploymentID      string
	buttonID          string
}

// NewSfdcConfig : Create SFDC Config Object
func NewSfdcConfig(authorizerURL string, liveChatURL string,
	basicAuthEmail string, basicAuthPassword, clientID string,
	clientSecret string, organizationID string, deploymentID,
	buttonID string) SfdcConfig {
	return SfdcConfig{
		authorizerURL,
		liveChatURL,
		basicAuthEmail,
		basicAuthPassword,
		clientID,
		clientSecret,
		organizationID,
		deploymentID,
		buttonID,
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
	resp := requests.HTTPGet(URL, qmap, nil)
	return resp
}

// CreateSession : create livechat session
func CreateSession(config SfdcConfig) *http.Response {
	URL := config.liveChatURL + "/chat/rest/System/SessionId"
	headers := map[string]string{
		"X-LIVEAGENT-AFFINITY":    "null",
		"X-LIVEAGENT-API-VERSION": "46",
	}

	resp := requests.HTTPGet(URL, nil, headers)
	return resp
}

// StartSession : start sfdc session
func StartSession(config SfdcConfig, username string, sessID string,
	sessKey string, affinityToken string) *http.Response {
	URL := config.liveChatURL + "/chat/rest/Chasitor/ChasitorInit"
	headers := map[string]string{
		"X-LIVEAGENT-AFFINITY":    affinityToken,
		"X-LIVEAGENT-API-VERSION": "46",
		"X-LIVEAGENT-SESSION-KEY": sessKey,
		"X-LIVEAGENT-SEQUENCE":    "1",
	}

	payload := map[string]string{
		"sessionId":           sessID,
		"organizationId":      config.organizationID,
		"deploymentId":        config.deploymentID,
		"buttonId":            config.buttonID,
		"userAgent":           "",
		"language":            "en-US",
		"screenResolution":    "1900x1080",
		"visitorName":         username,
		"receiveQueueUpdates": "true",
		"isPost":              "true",
	}

	resp := requests.HTTPPost(URL, payload, headers)
	return resp
}

// SendMsg : send message to live chat agent
func SendMsg(config SfdcConfig, sessID string, sessKey string,
	affinityToken string, msg string) *http.Response {
	URL := config.liveChatURL + "/chat/rest/Chasitor/ChatMessage"
	headers := map[string]string{
		"X-LIVEAGENT-AFFINITY":    affinityToken,
		"X-LIVEAGENT-API-VERSION": "41",
		"X-LIVEAGENT-SESSION-KEY": sessKey,
	}
	payload := map[string]string{
		"text": msg,
	}
	resp := requests.HTTPPost(URL, payload, headers)
	return resp
}

// Listen : listen to livechat messages updates
func Listen(config SfdcConfig, sessID string, sessKey string,
	affinityToken string) *http.Response {
	URL := config.liveChatURL + "/chat/rest/System/Messages"
	headers := map[string]string{
		"X-LIVEAGENT-AFFINITY":    affinityToken,
		"X-LIVEAGENT-API-VERSION": "41",
		"X-LIVEAGENT-SESSION-KEY": sessKey,
	}
	resp := requests.HTTPGet(URL, nil, headers)
	return resp
}
