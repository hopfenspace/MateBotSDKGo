package MateBotSDKGo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func get(endpoint string, filter map[string]string, sdk *SDK, retry bool) (int, []byte, error) {
	uri := sdk.BaseUrl + endpoint
	query := url.Values(map[string][]string{})
	if filter != nil {
		for k, v := range filter {
			if k != "" {
				query.Add(k, v)
			}
		}
		uri += "?" + query.Encode()
	}

	request, err := http.NewRequest("GET", uri, bytes.NewBuffer([]byte{}))
	if err != nil {
		return 0, nil, err
	}
	request.Header.Set("Authorization", "Bearer "+sdk.AccessToken)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(fmt.Sprintf("Error performing 'GET %s' request:", uri), err)
		return 0, nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Unexpected error while closing response buffer:", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Read all of response body failed:", err)
		return response.StatusCode, nil, err
	}

	if response.StatusCode == 401 && retry {
		log.Println("Invalid login token, trying to refresh...")
		token, err := GetLoginToken(sdk.Username, sdk.Password, sdk.BaseUrl)
		if err != nil {
			return 401, nil, err
		}
		sdk.AccessToken = token.AccessToken
		return get(endpoint, filter, sdk, false)
	}

	if response.StatusCode >= 400 {
		var e Error
		if err := json.Unmarshal(body, &e); err != nil {
			log.Println("No valid JSON body:", err)
			return response.StatusCode, nil, err
		}
		logError(e)
		return response.StatusCode, nil, e
	}
	return response.StatusCode, body, err
}

func Get(endpoint string, filter map[string]string, sdk *SDK) (int, []byte, error) {
	return get(endpoint, filter, sdk, true)
}

func post(endpoint string, content []byte, sdk *SDK, retry bool) (int, []byte, error) {
	uri := sdk.BaseUrl + endpoint

	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(content))
	if err != nil {
		return 0, []byte{}, err
	}
	request.Header.Set("Authorization", "Bearer "+sdk.AccessToken)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(fmt.Sprintf("Error performing 'POST %s' request:", uri), err)
		return 0, []byte{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Unexpected error while closing response buffer:", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Read all of response body failed:", err)
		return response.StatusCode, []byte{}, err
	}

	if response.StatusCode == 401 && retry {
		log.Println("Invalid login token, trying to refresh...")
		token, err := GetLoginToken(sdk.Username, sdk.Password, sdk.BaseUrl)
		if err != nil {
			return 401, []byte{}, err
		}
		sdk.AccessToken = token.AccessToken
		return post(endpoint, content, sdk, false)
	}

	if response.StatusCode >= 400 {
		var e Error
		if err := json.Unmarshal(body, &e); err != nil {
			log.Println("No valid JSON body:", err)
			return response.StatusCode, nil, err
		}
		logError(e)
		return response.StatusCode, nil, e
	}
	return response.StatusCode, body, err
}

func Post(endpoint string, content []byte, sdk *SDK) (int, []byte, error) {
	return post(endpoint, content, sdk, true)
}

type newAlias struct {
	UserId        int    `json:"user_id"`
	ApplicationId int    `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}
