package MateBotSDKGo

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/hopfenspace/MateBotSDKGo/internal"
)

type SDK struct {
	BaseUrl       string
	Username      string
	Password      string
	ApplicationID uint
	AccessToken   string
	Callbacks     []Callback
	APIVersion    uint
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

func (sdk *SDK) GetSettings() (Settings, error) {
	settings := Settings{}
	_, body, err := Get("/v1/settings", nil, sdk)
	if err != nil {
		return settings, err
	}

	if err = json.Unmarshal(body, &settings); err != nil {
		log.Println("No valid JSON body:", err)
		return settings, err
	}
	return settings, err
}

func (sdk *SDK) GetAliases(filter map[string]string) ([]Alias, error) {
	_, body, err := Get("/v1/aliases", filter, sdk)
	if err != nil {
		return nil, err
	}

	var aliases []Alias
	if err := json.Unmarshal(body, &aliases); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return aliases, err
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

func (sdk *SDK) GetCallbacks(filter map[string]string) ([]Callback, error) {
	_, body, err := Get("/v1/callbacks", filter, sdk)
	if err != nil {
		return nil, err
	}

	var callbacks []Callback
	if err := json.Unmarshal(body, &callbacks); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return callbacks, err
}

func (sdk *SDK) GetCommunisms(filter map[string]string) ([]Communism, error) {
	_, body, err := Get("/v1/communisms", filter, sdk)
	if err != nil {
		return nil, err
	}

	var communisms []Communism
	if err := json.Unmarshal(body, &communisms); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return communisms, err
}

func (sdk *SDK) GetConsumables(filter map[string]string) ([]Consumable, error) {
	_, body, err := Get("/v1/consumables", filter, sdk)
	if err != nil {
		return nil, err
	}

	var consumables []Consumable
	if err := json.Unmarshal(body, &consumables); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return consumables, err
}

func (sdk *SDK) GetPolls(filter map[string]string) ([]Poll, error) {
	_, body, err := Get("/v1/polls", filter, sdk)
	if err != nil {
		return nil, err
	}

	var polls []Poll
	if err := json.Unmarshal(body, &polls); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return polls, err
}

func (sdk *SDK) GetRefunds(filter map[string]string) ([]Refund, error) {
	_, body, err := Get("/v1/refunds", filter, sdk)
	if err != nil {
		return nil, err
	}

	var refunds []Refund
	if err := json.Unmarshal(body, &refunds); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return refunds, err
}

func (sdk *SDK) GetTransactions(filter map[string]string) ([]Transaction, error) {
	_, body, err := Get("/v1/transactions", filter, sdk)
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	if err := json.Unmarshal(body, &transactions); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return transactions, err
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

func (sdk *SDK) GetVotes(filter map[string]string) ([]Vote, error) {
	_, body, err := Get("/v1/votes", filter, sdk)
	if err != nil {
		return nil, err
	}

	var votes []Vote
	if err := json.Unmarshal(body, &votes); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return votes, err
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

func (sdk *SDK) DropInternalPrivilege(user any, issuer any) (User, error) {
	result := User{}
	err := checkStrOrPosInt(user, false)
	if err != nil {
		return result, err
	}
	err = checkStrOrPosInt(issuer, true)
	if err != nil {
		return result, err
	}

	content, err := json.Marshal(internal.UserPrivilegeDrop{
		User:   user,
		Issuer: issuer,
	})
	if err != nil {
		return result, err
	}

	_, body, err := Post("/v1/users/dropInternal", content, sdk)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("No valid JSON body:", err)
		return result, err
	}
	return result, err
}

func (sdk *SDK) DropPermissionPrivilege(user any, issuer any) (User, error) {
	result := User{}
	err := checkStrOrPosInt(user, false)
	if err != nil {
		return result, err
	}
	err = checkStrOrPosInt(issuer, true)
	if err != nil {
		return result, err
	}

	content, err := json.Marshal(internal.UserPrivilegeDrop{
		User:   user,
		Issuer: issuer,
	})
	if err != nil {
		return result, err
	}

	_, body, err := Post("/v1/users/dropPermission", content, sdk)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("No valid JSON body:", err)
		return result, err
	}
	return result, err
}

func (sdk *SDK) SetVoucher(debtor any, voucher any) (VoucherUpdate, error) {
	update := VoucherUpdate{}
	err := checkStrOrPosInt(debtor, false)
	if err != nil {
		return update, err
	}
	err = checkStrOrPosInt(voucher, true)
	if err != nil {
		return update, err
	}

	content, err := json.Marshal(internal.VoucherUpdateRequest{
		Debtor:  debtor,
		Voucher: voucher,
	})
	if err != nil {
		return update, err
	}

	_, body, err := Post("/v1/users/setVoucher", content, sdk)
	if err != nil {
		return update, err
	}

	if err = json.Unmarshal(body, &update); err != nil {
		log.Println("No valid JSON body:", err)
		return update, err
	}
	return update, err
}

func (sdk *SDK) DeleteUser(userID uint, issuer any) (User, error) {
	user := User{}
	err := checkStrOrPosInt(issuer, false)
	if err != nil {
		return user, err
	}

	content, err := json.Marshal(internal.IssuerIdBody{
		Id:     userID,
		Issuer: issuer,
	})
	if err != nil {
		return user, err
	}

	_, body, err := Post("/v1/users/delete", content, sdk)
	if err != nil {
		return user, err
	}

	if err = json.Unmarshal(body, &user); err != nil {
		log.Println("No valid JSON body:", err)
		return user, err
	}
	return user, err
}

func (sdk *SDK) NewAlias(userID uint, username string) (Alias, error) {
	alias := Alias{}
	content, err := json.Marshal(internal.NewAlias{
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

	users, err := sdk.GetUsers(map[string]string{"id": strconv.Itoa(int(user.Id))})
	if err != nil {
		return user, err
	}
	return users[0], err
}

func (sdk *SDK) ConfirmAlias(aliasID uint, issuer any) (Alias, error) {
	alias := Alias{}
	err := checkStrOrPosInt(issuer, false)
	if err != nil {
		return alias, err
	}

	content, err := json.Marshal(internal.IssuerIdBody{
		Id:     aliasID,
		Issuer: issuer,
	})
	if err != nil {
		return alias, err
	}

	_, body, err := Post("/v1/aliases/confirm", content, sdk)
	if err != nil {
		return alias, err
	}

	if err = json.Unmarshal(body, &alias); err != nil {
		log.Println("No valid JSON body:", err)
		return alias, err
	}
	return alias, err
}

func (sdk *SDK) DeleteAlias(aliasID uint, issuer any) (AliasDeletion, error) {
	deletion := AliasDeletion{}
	err := checkStrOrPosInt(issuer, false)
	if err != nil {
		return deletion, err
	}

	content, err := json.Marshal(internal.IssuerIdBody{
		Id:     aliasID,
		Issuer: issuer,
	})
	if err != nil {
		return deletion, err
	}

	_, body, err := Post("/v1/aliases/delete", content, sdk)
	if err != nil {
		return deletion, err
	}

	if err = json.Unmarshal(body, &deletion); err != nil {
		log.Println("No valid JSON body:", err)
		return deletion, err
	}
	return deletion, err
}

func (sdk *SDK) SendTransaction(sender any, receiver any, amount uint, reason string) (Transaction, error) {
	transaction := Transaction{}
	err := checkStrOrPosInt(sender, false)
	if err != nil {
		return transaction, err
	}
	err = checkStrOrPosInt(receiver, false)
	if err != nil {
		return transaction, err
	}

	content, err := json.Marshal(internal.NewTransaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Reason:   reason,
	})
	if err != nil {
		return transaction, err
	}

	_, body, err := Post("/v1/transactions/send", content, sdk)
	if err != nil {
		return transaction, err
	}

	if err = json.Unmarshal(body, &transaction); err != nil {
		log.Println("No valid JSON body:", err)
		return transaction, err
	}
	return transaction, err
}

func (sdk *SDK) ConsumeTransaction(consumer any, amount uint, consumable string) (Transaction, error) {
	transaction := Transaction{}
	err := checkStrOrPosInt(consumer, false)
	if err != nil {
		return transaction, err
	}

	content, err := json.Marshal(internal.NewConsumption{
		User:       consumer,
		Amount:     amount,
		Consumable: consumable,
	})
	if err != nil {
		return transaction, err
	}

	_, body, err := Post("/v1/transactions/consume", content, sdk)
	if err != nil {
		return transaction, err
	}

	if err = json.Unmarshal(body, &transaction); err != nil {
		log.Println("No valid JSON body:", err)
		return transaction, err
	}
	return transaction, err
}

func (sdk *SDK) NewCallback(url string, applicationID uint, sharedSecret string) (Callback, error) {
	callback := Callback{}
	content, err := json.Marshal(internal.NewCallback{
		Url:           url,
		ApplicationId: applicationID,
		SharedSecret:  sharedSecret,
	})
	if err != nil {
		return callback, err
	}

	_, body, err := Post("/v1/callbacks", content, sdk)
	if err != nil {
		return callback, err
	}

	if err = json.Unmarshal(body, &callback); err != nil {
		log.Println("No valid JSON body:", err)
		return callback, err
	}
	return callback, err
}

func (sdk *SDK) DeleteCallback(id uint) (bool, error) {
	content, err := json.Marshal(internal.IdBody{Id: id})
	if err != nil {
		return false, err
	}

	code, _, err := Delete("/v1/callbacks", content, sdk)
	if err != nil {
		return code == 204, err
	}
	return code == 204, err
}
