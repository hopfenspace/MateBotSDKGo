package MateBotSDKGo

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetLoginToken(username string, password string, baseURL string) (Token, error) {
	token := Token{}
	for baseURL[len(baseURL)-1] == '/' {
		baseURL = baseURL[:len(baseURL)-1]
	}

	response, err := http.PostForm(
		baseURL+"/v1/login",
		url.Values{"username": {username}, "password": {password}, "grant_type": {"password"}, "scope": {""}, "client_id": {""}, "client_secret": {""}})
	if err != nil {
		log.Println("Failed to connect to API server:", err)
		return token, err
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
		return token, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Read all of response body failed:", err)
		return token, err
	}

	if response.StatusCode == 200 {
		if err = json.Unmarshal(body, &token); err != nil {
			log.Println("No valid JSON body:", err)
			return token, err
		}
		return token, err
	} else {
		e := Error{}
		if err = json.Unmarshal(body, &e); err != nil {
			log.Println("No valid JSON body: err")
			return token, err
		}
		logError(e)
		return token, e
	}
}
