package MateBotSDKGo

import (
	"encoding/json"
	"log"
	"strconv"
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

func (sdk *SDK) GetStatus() (Status, error) {
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

func (sdk *SDK) GetApplications(filter map[string]string) ([]Application, error) {
	_, body, err := Get("/v1/applications", filter, sdk)
	if err != nil {
		return nil, err
	}

	var apps []Application
	if err := json.Unmarshal(body, &apps); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return apps, err
}

func (sdk *SDK) GetUsers(filter map[string]string) ([]User, error) {
	_, body, err := Get("/v1/users", filter, sdk)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(body, &users); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return users, err
}

func (sdk *SDK) NewUser() (User, error) {
	_, body, err := Post("/v1/users", nil, sdk)
	if err != nil {
		return User{}, err
	}

	user := User{}
	if err = json.Unmarshal(body, &user); err != nil {
		log.Println("No valid JSON body:", err)
		return User{}, err
	}
	return user, err
}

func (sdk *SDK) NewAlias(userID int, username string) (Alias, error) {
	alias := Alias{}
	content, err := json.Marshal(newAlias{
		UserId:        userID,
		ApplicationId: sdk.ApplicationID,
		Username:      username,
		Confirmed:     false,
	})
	if err != nil {
		return alias, err
	}

	_, body, err := Post("/v1/aliases", content, sdk)
	if err != nil {
		return alias, err
	}

	if err = json.Unmarshal(body, &alias); err != nil {
		log.Println("No valid JSON body:", err)
		return alias, err
	}
	return alias, err
}

func (sdk *SDK) NewUserWithAlias(username string) (User, error) {
	user, err := sdk.NewUser()
	if err != nil {
		return User{}, err
	}

	_, err = sdk.NewAlias(user.Id, username)
	if err != nil {
		return user, err
	}

	users, err := sdk.GetUsers(map[string]string{"id": strconv.Itoa(user.Id)})
	if err != nil {
		return user, err
	}
	return users[0], err
}
