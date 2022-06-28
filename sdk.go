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

func (sdk *SDK) abortSomething(obj uint, issuer any, endpoint string) ([]byte, error) {
	err := checkStrOrPosInt(issuer, false)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(IssuerIdBody{
		Id:     obj,
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

func (sdk *SDK) dropPrivilege(user any, issuer any, endpoint string) (User, error) {
	result := User{}
	err := checkStrOrPosInt(user, false)
	if err != nil {
		return result, err
	}
	err = checkStrOrPosInt(issuer, true)
	if err != nil {
		return result, err
	}

	content, err := json.Marshal(UserPrivilegeDrop{
		User:   user,
		Issuer: issuer,
	})
	if err != nil {
		return result, err
	}

	_, body, err := Post(endpoint, content, sdk)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("No valid JSON body:", err)
		return result, err
	}
	return result, err
}

func (sdk *SDK) DropInternalPrivilege(user any, issuer any) (User, error) {
	return sdk.dropPrivilege(user, issuer, "/v1/users/dropInternal")
}

func (sdk *SDK) DropPermissionPrivilege(user any, issuer any) (User, error) {
	return sdk.dropPrivilege(user, issuer, "/v1/users/dropPermission")
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

	content, err := json.Marshal(VoucherUpdateRequest{
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

	content, err := json.Marshal(IssuerIdBody{
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
	content, err := json.Marshal(NewAlias{
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

	content, err := json.Marshal(IssuerIdBody{
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

	content, err := json.Marshal(IssuerIdBody{
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

	content, err := json.Marshal(NewTransaction{
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

	content, err := json.Marshal(NewConsumption{
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

func (sdk *SDK) NewCommunism(creator any, amount uint, description string) (Communism, error) {
	communism := Communism{}
	err := checkStrOrPosInt(creator, false)
	if err != nil {
		return communism, err
	}

	content, err := json.Marshal(NewCommunism{
		Creator:     creator,
		Amount:      amount,
		Description: description,
	})
	if err != nil {
		return communism, err
	}

	_, body, err := Post("/v1/communisms", content, sdk)
	if err != nil {
		return communism, err
	}

	if err = json.Unmarshal(body, &communism); err != nil {
		log.Println("No valid JSON body:", err)
		return communism, err
	}
	return communism, err
}

func (sdk *SDK) abortOrCloseCommunism(communismID uint, issuer any, endpoint string) (Communism, error) {
	communism := Communism{}
	body, err := sdk.abortSomething(communismID, issuer, endpoint)
	if err != nil {
		return communism, err
	}

	if err = json.Unmarshal(body, &communism); err != nil {
		log.Println("No valid JSON body:", err)
		return communism, err
	}
	return communism, err
}

func (sdk *SDK) AbortCommunism(communismID uint, issuer any) (Communism, error) {
	return sdk.abortOrCloseCommunism(communismID, issuer, "/v1/communisms/abort")
}

func (sdk *SDK) CloseCommunism(communismID uint, issuer any) (Communism, error) {
	return sdk.abortOrCloseCommunism(communismID, issuer, "/v1/communisms/close")
}

func (sdk *SDK) changeCommunismParticipation(communismID uint, user any, endpoint string) (Communism, error) {
	communism := Communism{}
	err := checkStrOrPosInt(user, false)
	if err != nil {
		return communism, err
	}

	content, err := json.Marshal(CommunismParticipationUpdate{
		Id:   communismID,
		User: user,
	})
	if err != nil {
		return communism, err
	}

	_, body, err := Post(endpoint, content, sdk)
	if err != nil {
		return communism, err
	}

	if err = json.Unmarshal(body, &communism); err != nil {
		log.Println("No valid JSON body:", err)
		return communism, err
	}
	return communism, err
}

func (sdk *SDK) IncreaseCommunismParticipation(communismID uint, user any) (Communism, error) {
	return sdk.changeCommunismParticipation(communismID, user, "/v1/communisms/increaseParticipation")
}

func (sdk *SDK) DecreaseCommunismParticipation(communismID uint, user any) (Communism, error) {
	return sdk.changeCommunismParticipation(communismID, user, "/v1/communisms/decreaseParticipation")
}

func (sdk *SDK) NewPoll(user any, issuer any, variant string) (Poll, error) {
	poll := Poll{}
	err := checkStrOrPosInt(user, false)
	if err != nil {
		return poll, err
	}
	err = checkStrOrPosInt(issuer, false)
	if err != nil {
		return poll, err
	}

	content, err := json.Marshal(NewPoll{
		User:    user,
		Issuer:  issuer,
		Variant: variant,
	})
	if err != nil {
		return poll, err
	}

	_, body, err := Post("/v1/polls", content, sdk)
	if err != nil {
		return poll, err
	}

	if err = json.Unmarshal(body, &poll); err != nil {
		log.Println("No valid JSON body:", err)
		return poll, err
	}
	return poll, err
}

func (sdk *SDK) NewRefund(creator any, amount uint, description string) (Refund, error) {
	refund := Refund{}
	err := checkStrOrPosInt(creator, false)
	if err != nil {
		return refund, err
	}

	content, err := json.Marshal(NewRefund{
		Creator:     creator,
		Amount:      amount,
		Description: description,
	})
	if err != nil {
		return refund, err
	}

	_, body, err := Post("/v1/refunds", content, sdk)
	if err != nil {
		return refund, err
	}

	if err = json.Unmarshal(body, &refund); err != nil {
		log.Println("No valid JSON body:", err)
		return refund, err
	}
	return refund, err
}

func (sdk *SDK) AbortPoll(pollID uint, issuer any) (Poll, error) {
	poll := Poll{}
	body, err := sdk.abortSomething(pollID, issuer, "/v1/polls/abort")

	if err = json.Unmarshal(body, &poll); err != nil {
		log.Println("No valid JSON body:", err)
		return poll, err
	}
	return poll, err
}

func (sdk *SDK) AbortRefund(refundID uint, issuer any) (Refund, error) {
	refund := Refund{}
	body, err := sdk.abortSomething(refundID, issuer, "/v1/refunds/abort")

	if err = json.Unmarshal(body, &refund); err != nil {
		log.Println("No valid JSON body:", err)
		return refund, err
	}
	return refund, err
}

func (sdk *SDK) vote(ballotID uint, user any, vote bool, endpoint string) ([]byte, error) {
	err := checkStrOrPosInt(user, false)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(NewVote{
		User:     user,
		BallotId: ballotID,
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

func (sdk *SDK) VoteOnPollBallot(ballotID uint, user any, vote bool) (PollVote, error) {
	pollVote := PollVote{}
	body, err := sdk.vote(ballotID, user, vote, "/v1/polls/vote")

	if err = json.Unmarshal(body, &pollVote); err != nil {
		log.Println("No valid JSON body:", err)
		return pollVote, err
	}
	return pollVote, err
}

func (sdk *SDK) VoteOnRefundBallot(ballotID uint, user any, vote bool) (RefundVote, error) {
	refundVote := RefundVote{}
	body, err := sdk.vote(ballotID, user, vote, "/v1/refunds/vote")

	if err = json.Unmarshal(body, &refundVote); err != nil {
		log.Println("No valid JSON body:", err)
		return refundVote, err
	}
	return refundVote, err
}

func (sdk *SDK) NewCallback(url string, applicationID uint, sharedSecret string) (Callback, error) {
	callback := Callback{}
	content, err := json.Marshal(NewCallback{
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
	content, err := json.Marshal(IdBody{Id: id})
	if err != nil {
		return false, err
	}

	code, _, err := Delete("/v1/callbacks", content, sdk)
	if err != nil {
		return code == 204, err
	}
	return code == 204, err
}
