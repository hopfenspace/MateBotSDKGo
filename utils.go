package MateBotSDKGo

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Config struct {
	BaseUrl        string
	Username       string
	Password       string
	CallbackURL    *string
	CallbackSecret *string
	Currency       Currency
}

type Currency struct {
	Digits uint
	Factor uint
	Symbol string
}

func New(conf *Config) (*SDK, error) {
	if (conf.CallbackURL != nil && conf.CallbackSecret == nil) || (conf.CallbackURL == nil && conf.CallbackSecret != nil) {
		return nil, errors.New("options 'CallbackURL' and 'CallbackSecret' must both be set or omitted")
	}

	baseUrl := conf.BaseUrl
	for baseUrl[len(baseUrl)-1] == '/' {
		baseUrl = baseUrl[:len(baseUrl)-1]
	}
	sdk := SDK{
		BaseUrl:  baseUrl,
		Username: conf.Username,
		Password: conf.Password,
		Currency: conf.Currency,
	}

	token, err := GetLoginToken(conf.Username, conf.Password, conf.BaseUrl)
	if err != nil {
		return nil, err
	}
	sdk.AccessToken = token.AccessToken

	status, err := sdk.GetStatus()
	if err != nil {
		return nil, err
	}
	sdk.APIVersion = status.ApiVersion
	sdk.ServerVersion = status.ProjectVersion

	apps, err := sdk.GetApplications(map[string]string{"name": sdk.Username})
	if err != nil {
		return nil, err
	} else if len(apps) != 1 {
		return nil, errors.New("not exactly 1 result from app lookup")
	}
	sdk.ApplicationID = apps[0].ID

	if conf.CallbackURL != nil {
		callbacks, err := sdk.GetCallbacks(map[string]string{"application_id": strconv.Itoa(int(sdk.ApplicationID))})
		if err != nil {
			return nil, err
		}
		for _, callback := range callbacks {
			if success, err := sdk.DeleteCallback(callback.ID); err != nil || !success {
				return nil, err
			}
		}
		if _, err := sdk.NewCallback(*conf.CallbackURL, sdk.ApplicationID, *conf.CallbackSecret); err != nil {
			return nil, err
		}
		callbacks, err = sdk.GetCallbacks(map[string]string{"application_id": strconv.Itoa(int(sdk.ApplicationID))})
		if err != nil {
			return nil, err
		}
		sdk.Callbacks = callbacks
	}

	communityUsers, err := sdk.GetUsers(map[string]string{"community": "true"})
	if err != nil {
		return nil, err
	}
	sdk.CommunityUserID = communityUsers[0].ID
	for _, alias := range communityUsers[0].Aliases {
		if alias.Confirmed && alias.ApplicationID == sdk.ApplicationID {
			sdk.CommunityUsername = &alias.Username
		}
	}

	return &sdk, err
}

func logError(error Error) {
	log.Println(fmt.Sprintf("Error %d '%s %s': %s; %s", error.Status, error.Method, error.Request, error.Message, error.Details))
}

func checkStrOrPosInt(value any, allowNil bool) error {
	switch value.(type) {
	case nil:
		if !allowNil {
			return errors.New(fmt.Sprintf("invalid data type '%T', expected a string or unsigned/positive integer (>= 16 bits)", value))
		}
		return nil
	case uint, uint16, uint32, uint64, string:
		return nil
	case int:
		if v := value.(int); v < 0 {
			return errors.New(fmt.Sprintf("invalid value of '%T', value %d must not be negative", value, value))
		}
		return nil
	case int16:
		if v := value.(int16); v < 0 {
			return errors.New(fmt.Sprintf("invalid value of '%T', value %d must not be negative", value, value))
		}
		return nil
	case int32:
		if v := value.(int32); v < 0 {
			return errors.New(fmt.Sprintf("invalid value of '%T', value %d must not be negative", value, value))
		}
		return nil
	case int64:
		if v := value.(int64); v < 0 {
			return errors.New(fmt.Sprintf("invalid value of '%T', value %d must not be negative", value, value))
		}
		return nil
	default:
		return errors.New(fmt.Sprintf("invalid data type '%T', expected a string or unsigned/positive integer (>= 16 bits)", value))
	}
}
