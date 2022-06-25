package MateBotSDKGo

import (
	"encoding/json"
	"log"
)

type SDK struct {
	BaseUrl       string
	Username      string
	Password      string
	ApplicationID int
	AccessToken   string
	APIVersion    int
	ServerVersion VersionInfo
}

func (sdk SDK) GetStatus() (Status, error) {
	status := Status{}
	_, body, err := Get("/v1/status", nil, sdk)
	if err != nil {
		return status, err
	}
	if err = json.Unmarshal(body, &status); err != nil {
		log.Println("No valid JSON body:", err)
		return status, err
	}
	return status, err
}

func (sdk SDK) GetApplications(filter map[string]string) ([]Application, error) {
	_, body, err := Get("/v1/applications", filter, sdk)
	if err != nil {
		return nil, err
	}

	var apps []Application
	if err := json.Unmarshal(body, &apps); err != nil {
		log.Println("No valid JSON body:", err)
		return []Application{}, err
	}
	return apps, err
}
