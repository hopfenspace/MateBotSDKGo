package MateBotSDKGo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func get(endpoint string, filter map[string]string, config SDKConfig, retry bool) (int, []byte, error) {
	uri := config.BaseUrl + endpoint
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
		return 0, []byte{}, err
	}
	request.Header.Set("Authorization", "Bearer "+config.AccessToken)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(fmt.Sprintf("Error performing 'GET %s' request:", uri), err)
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
		token, err := GetLoginToken(config.Username, config.Password, config.BaseUrl)
		if err != nil {
			return 401, []byte{}, err
		}
		config.AccessToken = token.AccessToken
		return get(endpoint, filter, config, false)
	}
	return response.StatusCode, body, err
}

func Get(endpoint string, filter map[string]string, config SDKConfig) (int, []byte, error) {
	return get(endpoint, filter, config, true)
}

func post(endpoint string, content []byte, config SDKConfig, retry bool) (int, []byte, error) {
	uri := config.BaseUrl + endpoint

	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(content))
	if err != nil {
		return 0, []byte{}, err
	}
	request.Header.Set("Authorization", "Bearer "+config.AccessToken)
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
		token, err := GetLoginToken(config.Username, config.Password, config.BaseUrl)
		if err != nil {
			return 401, []byte{}, err
		}
		config.AccessToken = token.AccessToken
		return post(endpoint, content, config, false)
	}
	return response.StatusCode, body, err
}

func Post(endpoint string, content []byte, config SDKConfig) (int, []byte, error) {
	return post(endpoint, content, config, true)
}

func GetStatus(config SDKConfig) (Status, error) {
	code, body, err := Get("/v1/status", nil, config)
	if code != 200 {
		err = errors.New(fmt.Sprintf("response code %d for status request", code))
		return Status{}, err
	}

	status := Status{}
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Println("No valid JSON body:", err)
		return Status{}, err
	}
	return status, err
}
