package MateBotSDKGo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func GetStatus(config SDKConfig) (Status, error) {
	request, err := http.NewRequest("GET", config.BaseUrl+"/v1/status", bytes.NewBuffer([]byte{}))
	request.Header.Set("Authorization", "Bearer "+config.AccessToken)
	fmt.Println("request", request)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error fetching API status", err)
		return Status{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Unexpected error while closing response buffer:", err)
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Read all of response body failed:", err)
		return Status{}, err
	}

	if response.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("response code %d for status request", response.StatusCode))
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
