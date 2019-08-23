package facebook

import (
	"net/http"

	"github.com/agilasolutions/livechat-server-go/util/requests"
)

var hostURL = "https://graph.facebook.com/v2.9"

// GetUserInfo : get facebook user info
func GetUserInfo(userID string, reqFieds string,
	accessToken string) *http.Response {
	URL := hostURL + "/" + userID
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	qmap := map[string]interface{}{
		"fields":       reqFieds,
		"access_token": accessToken,
	}
	response := requests.HTTPGet(URL, qmap, headers)
	return response
}

// SendUserTextMsg : Send txt msg to user
func SendUserTextMsg(userID string, msg string,
	accessToken string) *http.Response {
	URL := hostURL + "/me/messages?access_token=" + accessToken
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	payload := map[string]interface{}{
		"recipient": map[string]string{
			"id": userID,
		},
		"message": map[string]string{
			"text": msg,
		},
	}

	resp := requests.HTTPPost(URL, payload, headers)
	return resp
}
