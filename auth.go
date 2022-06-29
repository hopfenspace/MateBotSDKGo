package MateBotSDKGo

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

func GetLoginToken(username string, password string, baseURL string) (*Token, error) {
	for baseURL[len(baseURL)-1] == '/' {
		baseURL = baseURL[:len(baseURL)-1]
	}

	response, err := http.PostForm(
		baseURL+"/v1/login",
		url.Values{"username": {username}, "password": {password}, "grant_type": {"password"}, "scope": {""}, "client_id": {""}, "client_secret": {""}})
	if err != nil {
		log.Println("Failed to connect to API server:", err)
		return nil, err
	}

	if response != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println("Closing response body failed:", err)
			}
		}(response.Body)
	} else {
		err = errors.New("response body is null")
		log.Println("Failed to access body:", err)
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Read all of response body failed:", err)
		return nil, err
	}

	if response.StatusCode == 200 {
		var token *Token
		if err = json.Unmarshal(body, &token); err != nil {
			log.Println("No valid JSON body:", err)
			return nil, err
		}
		return token, err
	} else {
		var e Error
		if err = json.Unmarshal(body, &e); err != nil {
			log.Println("No valid JSON body: err")
			return nil, err
		}
		logError(e)
		return nil, e
	}
}
