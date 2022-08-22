package MateBotSDKGo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
)

type SDK struct {
	BaseUrl           string
	Username          string
	Password          string
	ApplicationID     uint
	AccessToken       string
	Callbacks         []*Callback
	APIVersion        uint
	ServerVersion     VersionInfo
	Currency          Currency
	CommunityUserID   uint
	CommunityUsername *string
}

func (sdk *SDK) FormatUsername(user *User, findUsername *func(uint) (string, error)) (string, error) {
	if user == nil {
		return "Unknown user", errors.New("invalid user")
	}
	if findUsername != nil {
		username, err := (*findUsername)(user.ID)
		if err == nil {
			return username, nil
		}
	}
	for _, alias := range user.Aliases {
		if alias.ApplicationID == sdk.ApplicationID && alias.Confirmed {
			return alias.Username, nil
		}
	}
	if user.ID == sdk.CommunityUserID {
		return "Community", nil
	}
	return fmt.Sprintf("User %d", user.ID), errors.New("no results")
}

func (sdk *SDK) FormatBalance(balance int) string {
	v := float64(balance) / float64(sdk.Currency.Factor)
	return fmt.Sprintf("%."+strconv.Itoa(int(sdk.Currency.Digits))+"f%s", v, sdk.Currency.Symbol)
}

func (sdk *SDK) GetStatus() (*Status, error) {
	_, body, err := Get("/v1/status", nil, sdk)
	if err != nil {
		return nil, err
	}

	var status *Status
	if err = json.Unmarshal(body, &status); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return status, err
}

func (sdk *SDK) GetSettings() (*Settings, error) {
	_, body, err := Get("/v1/settings", nil, sdk)
	if err != nil {
		return nil, err
	}

	var settings *Settings
	if err = json.Unmarshal(body, &settings); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return settings, err
}

func (sdk *SDK) GetAliases(filter map[string]string) ([]*Alias, error) {
	_, body, err := Get("/v1/aliases", filter, sdk)
	if err != nil {
		return nil, err
	}

	var aliases []*Alias
	if err := json.Unmarshal(body, &aliases); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return aliases, err
}

func (sdk *SDK) GetApplications(filter map[string]string) ([]*Application, error) {
	_, body, err := Get("/v1/applications", filter, sdk)
	if err != nil {
		return nil, err
	}

	var apps []*Application
	if err := json.Unmarshal(body, &apps); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return apps, err
}

func (sdk *SDK) GetCallbacks(filter map[string]string) ([]*Callback, error) {
	_, body, err := Get("/v1/callbacks", filter, sdk)
	if err != nil {
		return nil, err
	}

	var callbacks []*Callback
	if err := json.Unmarshal(body, &callbacks); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return callbacks, err
}

func (sdk *SDK) GetCommunisms(filter map[string]string) ([]*Communism, error) {
	_, body, err := Get("/v1/communisms", filter, sdk)
	if err != nil {
		return nil, err
	}

	var communisms []*Communism
	if err := json.Unmarshal(body, &communisms); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return communisms, err
}

func (sdk *SDK) GetConsumables(filter map[string]string) ([]*Consumable, error) {
	_, body, err := Get("/v1/consumables", filter, sdk)
	if err != nil {
		return nil, err
	}

	var consumables []*Consumable
	if err := json.Unmarshal(body, &consumables); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return consumables, err
}

func (sdk *SDK) GetPolls(filter map[string]string) ([]*Poll, error) {
	_, body, err := Get("/v1/polls", filter, sdk)
	if err != nil {
		return nil, err
	}

	var polls []*Poll
	if err := json.Unmarshal(body, &polls); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return polls, err
}

func (sdk *SDK) GetRefunds(filter map[string]string) ([]*Refund, error) {
	_, body, err := Get("/v1/refunds", filter, sdk)
	if err != nil {
		return nil, err
	}

	var refunds []*Refund
	if err := json.Unmarshal(body, &refunds); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return refunds, err
}

func (sdk *SDK) GetTransactions(filter map[string]string) ([]*Transaction, error) {
	_, body, err := Get("/v1/transactions", filter, sdk)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	if err := json.Unmarshal(body, &transactions); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return transactions, err
}

func (sdk *SDK) GetUsers(filter map[string]string) ([]*User, error) {
	_, body, err := Get("/v1/users", filter, sdk)
	if err != nil {
		return nil, err
	}

	var users []*User
	if err := json.Unmarshal(body, &users); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return users, err
}

func (sdk *SDK) GetVotes(filter map[string]string) ([]*Vote, error) {
	_, body, err := Get("/v1/votes", filter, sdk)
	if err != nil {
		return nil, err
	}

	var votes []*Vote
	if err := json.Unmarshal(body, &votes); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return votes, err
}

func (sdk *SDK) GetUser(userIdOrUsername any, extendedFilter *map[string]string) (*User, error) {
	err := checkStrOrPosInt(userIdOrUsername, false)
	if err != nil {
		return nil, err
	}

	var userID int
	switch userIdOrUsername.(type) {
	case uint, uint16, uint32, uint64:
		userID = int(userIdOrUsername.(uint))
	case int, int16, int32, int64:
		userID = userIdOrUsername.(int)
	case string:
		filter := map[string]string{
			"active":               "true",
			"alias_confirmed":      "true",
			"alias_username":       userIdOrUsername.(string),
			"alias_application_id": strconv.Itoa(int(sdk.ApplicationID)),
		}
		if extendedFilter != nil {
			for k, v := range *extendedFilter {
				filter[k] = v
			}
		}
		users, err := sdk.GetUsers(filter)
		if err != nil {
			return nil, err
		} else if len(users) < 1 {
			return nil, errors.New("no user found")
		} else if len(users) > 1 {
			return nil, errors.New("ambiguous username")
		}
		return users[0], nil
	default:
		return nil, errors.New("invalid")
	}

	filter := map[string]string{"active": "true", "id": strconv.Itoa(userID)}
	if extendedFilter != nil {
		for k, v := range *extendedFilter {
			filter[k] = v
		}
	}
	users, err := sdk.GetUsers(filter)
	if err != nil {
		return nil, err
	} else if len(users) < 1 {
		return nil, errors.New("no user found")
	} else if len(users) > 1 {
		return nil, errors.New("ambiguous username")
	}
	return users[0], nil
}

func (sdk *SDK) FindSponsoringUser(issuer *User) (*User, error) {
	if issuer == nil {
		return nil, errors.New("invalid user account")
	} else if !issuer.Active {
		return nil, errors.New("this user account has been disabled")
	} else if issuer.External {
		return nil, errors.New("you don't have the permission to request this information")
	}

	users, err := sdk.GetUsers(map[string]string{"community": "false", "active": "true"})
	if err != nil {
		return nil, err
	}
	sort.Slice(users, func(i int, j int) bool { return users[i].Balance < users[j].Balance })
	if users[0].Balance >= 0 {
		return nil, nil
	}
	return users[0], nil
}

func (sdk *SDK) GetCommunityBalance(issuer *User) (int, error) {
	if issuer == nil {
		return 0, errors.New("invalid user account")
	} else if !issuer.Active {
		return 0, errors.New("this user account has been disabled")
	} else if issuer.External {
		return 0, errors.New("you don't have the permission to request this information")
	}
	community, err := sdk.GetUser(sdk.CommunityUserID, &map[string]string{"community": "true"})
	if err != nil {
		return 0, err
	}
	return community.Balance, nil
}

func (sdk *SDK) abortSomething(obj uint, issuer any, endpoint string) ([]byte, error) {
	err := checkStrOrPosInt(issuer, false)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(IssuerIDBody{
		ID:     obj,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, sdk)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (sdk *SDK) NewUser() (*User, error) {
	_, body, err := Post("/v1/users", nil, sdk)
	if err != nil {
		return nil, err
	}

	var user *User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return user, err
}

func (sdk *SDK) dropPrivilege(user any, issuer any, endpoint string) (*User, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(issuer, true); err != nil {
		return nil, err
	}

	content, err := json.Marshal(UserPrivilegeDrop{
		User:   user,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, sdk)
	if err != nil {
		return nil, err
	}

	var result *User
	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return result, err
}

func (sdk *SDK) DropInternalPrivilege(user any, issuer any) (*User, error) {
	return sdk.dropPrivilege(user, issuer, "/v1/users/dropInternal")
}

func (sdk *SDK) DropPermissionPrivilege(user any, issuer any) (*User, error) {
	return sdk.dropPrivilege(user, issuer, "/v1/users/dropPermission")
}

func (sdk *SDK) SetVoucher(debtor any, voucher any, issuer any) (*VoucherUpdate, error) {
	if err := checkStrOrPosInt(debtor, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(voucher, true); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(VoucherUpdateRequest{
		Debtor:  debtor,
		Voucher: voucher,
		Issuer:  issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/users/setVoucher", content, sdk)
	if err != nil {
		return nil, err
	}

	var update *VoucherUpdate
	if err = json.Unmarshal(body, &update); err != nil {
		log.Println("No valid JSON body:", err)
		return update, err
	}
	return update, err
}

func (sdk *SDK) DeleteUser(userID uint, issuer any) (*User, error) {
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(IssuerIDBody{
		ID:     userID,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/users/delete", content, sdk)
	if err != nil {
		return nil, err
	}

	var user *User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return user, err
}

func (sdk *SDK) newAlias(userID uint, username string, confirmed *bool) (*Alias, error) {
	var content []byte
	if confirmed == nil {
		if c, err := json.Marshal(NewAlias{
			UserID:        userID,
			ApplicationID: sdk.ApplicationID,
			Username:      username,
			Confirmed:     false,
		}); err != nil {
			return nil, err
		} else {
			content = c
		}
	} else {
		if c, err := json.Marshal(NewAlias{
			UserID:        userID,
			ApplicationID: sdk.ApplicationID,
			Username:      username,
			Confirmed:     *confirmed,
		}); err != nil {
			return nil, err
		} else {
			content = c
		}
	}

	_, body, err := Post("/v1/aliases", content, sdk)
	if err != nil {
		return nil, err
	}

	var alias *Alias
	if err = json.Unmarshal(body, &alias); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return alias, err
}

func (sdk *SDK) NewAlias(userID uint, username string) (*Alias, error) {
	c := false
	return sdk.newAlias(userID, username, &c)
}

func (sdk *SDK) NewUserWithAlias(username string) (*User, error) {
	user, err := sdk.NewUser()
	if err != nil {
		return nil, err
	}

	c := true
	_, err = sdk.newAlias(user.ID, username, &c)
	if err != nil {
		return user, err
	}

	users, err := sdk.GetUsers(map[string]string{"id": strconv.Itoa(int(user.ID))})
	if err != nil {
		return user, err
	}
	return users[0], err
}

func (sdk *SDK) ConfirmAlias(aliasID uint, issuer any) (*Alias, error) {
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(IssuerIDBody{
		ID:     aliasID,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/aliases/confirm", content, sdk)
	if err != nil {
		return nil, err
	}

	var alias *Alias
	if err = json.Unmarshal(body, &alias); err != nil {
		log.Println("No valid JSON body:", err)
		return alias, err
	}
	return alias, err
}

func (sdk *SDK) DeleteAlias(aliasID uint, issuer any) (*AliasDeletion, error) {
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(IssuerIDBody{
		ID:     aliasID,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/aliases/delete", content, sdk)
	if err != nil {
		return nil, err
	}

	var deletion *AliasDeletion
	if err = json.Unmarshal(body, &deletion); err != nil {
		log.Println("No valid JSON body:", err)
		return deletion, err
	}
	return deletion, err
}

func (sdk *SDK) SendTransaction(sender any, receiver any, amount uint, reason string) (*Transaction, error) {
	if err := checkStrOrPosInt(sender, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(receiver, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(NewTransaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Reason:   reason,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/transactions/send", content, sdk)
	if err != nil {
		return nil, err
	}

	var transaction *Transaction
	if err = json.Unmarshal(body, &transaction); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return transaction, err
}

func (sdk *SDK) ConsumeTransaction(consumer any, amount uint, consumable string) (*Transaction, error) {
	if err := checkStrOrPosInt(consumer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(NewConsumption{
		User:       consumer,
		Amount:     amount,
		Consumable: consumable,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/transactions/consume", content, sdk)
	if err != nil {
		return nil, err
	}

	var transaction *Transaction
	if err = json.Unmarshal(body, &transaction); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return transaction, err
}

func (sdk *SDK) NewCommunism(creator any, amount uint, description string) (*Communism, error) {
	if err := checkStrOrPosInt(creator, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(NewCommunism{
		Creator:     creator,
		Amount:      amount,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/communisms", content, sdk)
	if err != nil {
		return nil, err
	}

	var communism *Communism
	if err = json.Unmarshal(body, &communism); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return communism, err
}

func (sdk *SDK) abortOrCloseCommunism(communismID uint, issuer any, endpoint string) (*Communism, error) {
	body, err := sdk.abortSomething(communismID, issuer, endpoint)
	if err != nil {
		return nil, err
	}

	var communism *Communism
	if err = json.Unmarshal(body, &communism); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return communism, err
}

func (sdk *SDK) AbortCommunism(communismID uint, issuer any) (*Communism, error) {
	return sdk.abortOrCloseCommunism(communismID, issuer, "/v1/communisms/abort")
}

func (sdk *SDK) CloseCommunism(communismID uint, issuer any) (*Communism, error) {
	return sdk.abortOrCloseCommunism(communismID, issuer, "/v1/communisms/close")
}

func (sdk *SDK) changeCommunismParticipation(communismID uint, user any, endpoint string) (*Communism, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(CommunismParticipationUpdate{
		ID:   communismID,
		User: user,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, sdk)
	if err != nil {
		return nil, err
	}

	var communism *Communism
	if err = json.Unmarshal(body, &communism); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return communism, err
}

func (sdk *SDK) IncreaseCommunismParticipation(communismID uint, user any) (*Communism, error) {
	return sdk.changeCommunismParticipation(communismID, user, "/v1/communisms/increaseParticipation")
}

func (sdk *SDK) DecreaseCommunismParticipation(communismID uint, user any) (*Communism, error) {
	return sdk.changeCommunismParticipation(communismID, user, "/v1/communisms/decreaseParticipation")
}

func (sdk *SDK) NewPoll(user any, issuer any, variant string) (*Poll, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(NewPoll{
		User:    user,
		Issuer:  issuer,
		Variant: variant,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/polls", content, sdk)
	if err != nil {
		return nil, err
	}

	var poll *Poll
	if err = json.Unmarshal(body, &poll); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return poll, err
}

func (sdk *SDK) NewRefund(creator any, amount uint, description string) (*Refund, error) {
	if err := checkStrOrPosInt(creator, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(NewRefund{
		Creator:     creator,
		Amount:      amount,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/refunds", content, sdk)
	if err != nil {
		return nil, err
	}

	var refund *Refund
	if err = json.Unmarshal(body, &refund); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return refund, err
}

func (sdk *SDK) AbortPoll(pollID uint, issuer any) (*Poll, error) {
	body, err := sdk.abortSomething(pollID, issuer, "/v1/polls/abort")
	if err != nil {
		return nil, err
	}

	var poll *Poll
	if err = json.Unmarshal(body, &poll); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return poll, err
}

func (sdk *SDK) AbortRefund(refundID uint, issuer any) (*Refund, error) {
	body, err := sdk.abortSomething(refundID, issuer, "/v1/refunds/abort")
	if err != nil {
		return nil, err
	}

	var refund *Refund
	if err = json.Unmarshal(body, &refund); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return refund, err
}

func (sdk *SDK) vote(ballotID uint, user any, vote bool, endpoint string) ([]byte, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(NewVote{
		User:     user,
		BallotID: ballotID,
		Vote:     vote,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, sdk)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (sdk *SDK) VoteOnPollBallot(ballotID uint, user any, vote bool) (*PollVote, error) {
	body, err := sdk.vote(ballotID, user, vote, "/v1/polls/vote")
	if err != nil {
		return nil, err
	}

	var pollVote *PollVote
	if err = json.Unmarshal(body, &pollVote); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return pollVote, err
}

func (sdk *SDK) VoteOnRefundBallot(ballotID uint, user any, vote bool) (*RefundVote, error) {
	body, err := sdk.vote(ballotID, user, vote, "/v1/refunds/vote")
	if err != nil {
		return nil, err
	}

	var refundVote *RefundVote
	if err = json.Unmarshal(body, &refundVote); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return refundVote, err
}

func (sdk *SDK) NewCallback(url string, applicationID uint, sharedSecret string) (*Callback, error) {
	content, err := json.Marshal(NewCallback{
		Url:           url,
		ApplicationID: applicationID,
		SharedSecret:  sharedSecret,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/callbacks", content, sdk)
	if err != nil {
		return nil, err
	}

	var callback *Callback
	if err = json.Unmarshal(body, &callback); err != nil {
		log.Println("No valid JSON body:", err)
		return nil, err
	}
	return callback, err
}

func (sdk *SDK) DeleteCallback(id uint) (bool, error) {
	content, err := json.Marshal(IDBody{ID: id})
	if err != nil {
		return false, err
	}

	code, _, err := Delete("/v1/callbacks", content, sdk)
	if err != nil {
		return code == 204, err
	}
	return code == 204, err
}
