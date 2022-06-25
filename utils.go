package MateBotSDKGo

import (
	"errors"
	"fmt"
	"log"
)

func New(baseURL string, username string, password string) (SDK, error) {
	for baseURL[len(baseURL)-1] == '/' {
		baseURL = baseURL[:len(baseURL)-1]
	}
	sdk := SDK{
		BaseUrl:  baseURL,
		Username: username,
		Password: password,
	}
	token, err := GetLoginToken(username, password, baseURL)
	if err != nil {
		return SDK{}, err
	}
	sdk.AccessToken = token.AccessToken
	status, err := sdk.GetStatus()
	if err != nil {
		return SDK{}, err
	}
	sdk.APIVersion = status.ApiVersion
	sdk.ServerVersion = status.ProjectVersion
	apps, err := sdk.GetApplications(map[string]string{"name": sdk.Username})
	if err != nil {
		return SDK{}, err
	} else if len(apps) != 1 {
		return SDK{}, errors.New("not exactly 1 result from app lookup")
	}
	sdk.ApplicationID = apps[0].Id
	return sdk, err
}

func logError(error Error) {
	log.Println(fmt.Sprintf("Error %d '%s %s': %s; %s", error.Status, error.Method, error.Request, error.Message, error.Details))
}
