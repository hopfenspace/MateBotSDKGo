package MateBotSDKGo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
)

type sdk struct {
	BaseUrl           string
	Username          string
	password          string
	applicationID     uint64
	accessToken       string
	Callbacks         []*Callback
	Currency          Currency
	communityUserID   uint64
	CommunityUsername *string
}

func (s *sdk) GetThisApplicationID() uint64 {
	return s.applicationID
}

func (s *sdk) GetThisApplicationName() string {
	return s.Username
}

func (s *sdk) GetCommunityID() uint64 {
	return s.communityUserID
}

func (s *sdk) GetCommunityUsername() *string {
	return s.CommunityUsername
}

func (s *sdk) GetCurrency() Currency {
	return s.Currency
}

func (s *sdk) FormatBalance(balance int64) string {
	v := float64(balance) / float64(s.Currency.Factor)
	return fmt.Sprintf("%."+strconv.FormatInt(s.Currency.Digits, 10)+"f %s", v, s.Currency.Symbol)
}

func (s *sdk) GetHealth() (bool, error) {
	uri := s.BaseUrl + "/v1/health"
	request, err := http.NewRequest("GET", uri, bytes.NewBuffer([]byte{}))
	if err != nil {
		return false, err
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(fmt.Sprintf("Error performing 'GET %s' request:", uri), err)
		return false, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Unexpected error while closing response buffer:", err)
		}
	}(response.Body)

	return response.StatusCode == 200, err
}

func (s *sdk) GetSettings() (*Settings, error) {
	_, body, err := Get("/v1/settings", nil, s)
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

func (s *sdk) GetAliases(filter map[string]string) ([]*Alias, error) {
	_, body, err := Get("/v1/aliases", filter, s)
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

func (s *sdk) GetApplications(filter map[string]string) ([]*Application, error) {
	_, body, err := Get("/v1/applications", filter, s)
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

func (s *sdk) GetCallbacks(filter map[string]string) ([]*Callback, error) {
	_, body, err := Get("/v1/callbacks", filter, s)
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

func (s *sdk) GetCommunisms(filter map[string]string) ([]*Communism, error) {
	_, body, err := Get("/v1/communisms", filter, s)
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

func (s *sdk) GetConsumables(filter map[string]string) ([]*Consumable, error) {
	_, body, err := Get("/v1/consumables", filter, s)
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

func (s *sdk) GetPolls(filter map[string]string) ([]*Poll, error) {
	_, body, err := Get("/v1/polls", filter, s)
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

func (s *sdk) GetRefunds(filter map[string]string) ([]*Refund, error) {
	_, body, err := Get("/v1/refunds", filter, s)
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

func (s *sdk) GetTransactions(filter map[string]string) ([]*Transaction, error) {
	_, body, err := Get("/v1/transactions", filter, s)
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

func (s *sdk) GetUsers(filter map[string]string) ([]*User, error) {
	_, body, err := Get("/v1/users", filter, s)
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

func (s *sdk) GetVotes(filter map[string]string) ([]*Vote, error) {
	_, body, err := Get("/v1/votes", filter, s)
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

func (s *sdk) lookupUsers(name string, extendedFilter *map[string]string, isAlias bool) ([]*User, error) {
	filter := map[string]string{"active": "true"}
	if isAlias {
		filter["alias_confirmed"] = "true"
		filter["alias_username"] = name
		filter["alias_application_id"] = strconv.FormatUint(s.applicationID, 10)
	} else {
		filter["name"] = name
	}
	if extendedFilter != nil {
		for k, v := range *extendedFilter {
			filter[k] = v
		}
	}
	return s.GetUsers(filter)
}

func (s *sdk) GetUser(userIdOrUsername any, extendedFilter *map[string]string) (*User, error) {
	err := checkStrOrPosInt(userIdOrUsername, false)
	if err != nil {
		return nil, err
	}

	var userID uint64
	switch userIdOrUsername.(type) {
	case uint, uint16, uint32, uint64:
		userID = userIdOrUsername.(uint64)
	case int, int16, int32, int64:
		userID = userIdOrUsername.(uint64)
	case string:
		users, err := s.lookupUsers(userIdOrUsername.(string), extendedFilter, false)
		if err != nil {
			return nil, err
		} else if len(users) == 1 {
			return users[0], nil
		} else {
			users, err = s.lookupUsers(userIdOrUsername.(string), extendedFilter, true)
			if err != nil {
				return nil, err
			} else if len(users) < 1 {
				return nil, errors.New("no user found")
			} else if len(users) > 1 {
				return nil, errors.New("ambiguous username")
			}
			return users[0], nil
		}
	default:
		return nil, errors.New("invalid")
	}

	filter := map[string]string{"active": "true", "id": strconv.FormatUint(userID, 10)}
	if extendedFilter != nil {
		for k, v := range *extendedFilter {
			filter[k] = v
		}
	}
	users, err := s.GetUsers(filter)
	if err != nil {
		return nil, err
	} else if len(users) < 1 {
		return nil, errors.New("no user found")
	} else if len(users) > 1 {
		return nil, errors.New("ambiguous username")
	}
	return users[0], nil
}

// GetVerifiedUser returns a verified core user instance.
// Verification steps include the active check and that an alias
// for the current application exists, which has to be confirmed.
// The minimal privilege level will be External implicitly if unspecified.
// The verification includes a check that the user has at least
// the specified privilege level. Use this function with
// result(s) from the GetUser function. The difference between those
// two functions is that GetUser (and GetUsers, respectively) should be
// used for lookups of foreign users, while this should be used for app users.
func (s *sdk) GetVerifiedUser(coreUserID uint64, minimalLevel *PrivilegeLevel) (*User, error) {
	users, err := s.GetUsers(map[string]string{
		"id":                   strconv.FormatUint(coreUserID, 10),
		"active":               "true",
		"alias_application_id": strconv.FormatUint(s.GetThisApplicationID(), 10),
	})
	if err != nil {
		return nil, err
	} else if len(users) > 1 {
		return nil, errors.New("ambiguous user ID")
	} else if len(users) < 1 {
		return nil, errors.New("no user was found for the given query")
	}

	user := users[0]
	verifiedAliasFound := false
	for i := range user.Aliases {
		if user.Aliases[i].Confirmed && user.Aliases[i].ApplicationID == s.GetThisApplicationID() {
			verifiedAliasFound = true
		}
	}
	if !verifiedAliasFound {
		return nil, errors.New(fmt.Sprintf("The user alias for %s is not confirmed yet. It can't be used while the connection to the other MateBot apps wasn't verified.", user.Name))
	}

	if minimalLevel == nil {
		l := External
		minimalLevel = &l
	}
	if user.Privilege() < *minimalLevel {
		return nil, errors.New("you don't have the required privileges to perform this operation, maybe open a poll to request them")
	}
	return user, nil
}

func (s *sdk) IsUserConfirmed(user *User) bool {
	if user == nil {
		return false
	}
	for _, alias := range user.Aliases {
		if alias.Confirmed && alias.ApplicationID == s.GetThisApplicationID() {
			return true
		}
	}
	return false
}

func (s *sdk) FindSponsoringUser(issuer *User) (*User, error) {
	if issuer == nil {
		return nil, errors.New("invalid user account")
	} else if !issuer.Active {
		return nil, errors.New("this user account has been disabled")
	} else if issuer.Privilege() < Internal {
		return nil, errors.New("you don't have the permission to request this information")
	}

	users, err := s.GetUsers(map[string]string{"community": "false", "active": "true"})
	if err != nil {
		return nil, err
	}
	sort.Slice(users, func(i int, j int) bool { return users[i].Balance < users[j].Balance })
	if users[0].Balance >= 0 {
		return nil, nil
	}
	return users[0], nil
}

func (s *sdk) GetCommunityBalance(issuer *User) (int64, error) {
	if issuer == nil {
		return 0, errors.New("invalid user account")
	} else if !issuer.Active {
		return 0, errors.New("this user account has been disabled")
	} else if issuer.Privilege() < Internal {
		return 0, errors.New("you don't have the permission to request this information")
	}
	community, err := s.GetUser(s.communityUserID, &map[string]string{"community": "true"})
	if err != nil {
		return 0, err
	}
	return community.Balance, nil
}

func (s *sdk) abortSomething(obj uint64, issuer any, endpoint string) ([]byte, error) {
	err := checkStrOrPosInt(issuer, false)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(issuerID{
		ID:     obj,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, s)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (s *sdk) NewUser(username string) (*User, error) {
	content, err := json.Marshal(newUser{
		Name: username,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/users", content, s)
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

func (s *sdk) dropPrivilege(user any, issuer any, endpoint string) (*User, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(issuer, true); err != nil {
		return nil, err
	}

	content, err := json.Marshal(userPrivilegeDrop{
		User:   user,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, s)
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

func (s *sdk) DropInternalPrivilege(user any, issuer any) (*User, error) {
	return s.dropPrivilege(user, issuer, "/v1/users/dropInternal")
}

func (s *sdk) DropPermissionPrivilege(user any, issuer any) (*User, error) {
	return s.dropPrivilege(user, issuer, "/v1/users/dropPermission")
}

func (s *sdk) SetUsername(issuer any, newName string) (*User, error) {
	if err := checkStrOrPosInt(issuer, true); err != nil {
		return nil, err
	}

	content, err := json.Marshal(usernameUpdate{
		Name:   newName,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/users/setName", content, s)
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

func (s *sdk) SetVoucher(debtor any, voucher any, issuer any) (*VoucherUpdate, error) {
	if err := checkStrOrPosInt(debtor, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(voucher, true); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(voucherUpdate{
		Debtor:  debtor,
		Voucher: voucher,
		Issuer:  issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/users/setVoucher", content, s)
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

func (s *sdk) DeleteUser(userID uint64, issuer any) (*User, error) {
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(issuerID{
		ID:     userID,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/users/delete", content, s)
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

func (s *sdk) newAlias(userID uint64, username string, confirmed *bool) (*Alias, error) {
	var content []byte
	if confirmed == nil {
		if c, err := json.Marshal(newAlias{
			UserID:        userID,
			ApplicationID: s.applicationID,
			Username:      username,
			Confirmed:     false,
		}); err != nil {
			return nil, err
		} else {
			content = c
		}
	} else {
		if c, err := json.Marshal(newAlias{
			UserID:        userID,
			ApplicationID: s.applicationID,
			Username:      username,
			Confirmed:     *confirmed,
		}); err != nil {
			return nil, err
		} else {
			content = c
		}
	}

	_, body, err := Post("/v1/aliases", content, s)
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

func (s *sdk) NewAlias(userID uint64, username string) (*Alias, error) {
	c := false
	return s.newAlias(userID, username, &c)
}

func (s *sdk) NewUserWithAlias(username string) (*User, error) {
	user, err := s.NewUser(username)
	if err != nil {
		return nil, err
	}

	c := true
	_, err = s.newAlias(user.ID, username, &c)
	if err != nil {
		return user, err
	}

	users, err := s.GetUsers(map[string]string{"id": strconv.FormatUint(user.ID, 10)})
	if err != nil {
		return user, err
	}
	return users[0], err
}

func (s *sdk) ConfirmAlias(aliasID uint64, issuer any) (*Alias, error) {
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(issuerID{
		ID:     aliasID,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/aliases/confirm", content, s)
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

func (s *sdk) DeleteAlias(aliasID uint64, issuer any) (*AliasDeletion, error) {
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(issuerID{
		ID:     aliasID,
		Issuer: issuer,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/aliases/delete", content, s)
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

func (s *sdk) SendTransaction(sender any, receiver any, amount uint64, reason string) (*Transaction, error) {
	if err := checkStrOrPosInt(sender, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(receiver, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(newTransaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Reason:   reason,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/transactions/send", content, s)
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

func (s *sdk) ConsumeTransaction(consumer any, amount uint64, consumable string) (*Transaction, error) {
	if err := checkStrOrPosInt(consumer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(newConsumption{
		User:       consumer,
		Amount:     amount,
		Consumable: consumable,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/transactions/consume", content, s)
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

func (s *sdk) NewCommunism(creator any, amount uint64, description string) (*Communism, error) {
	if err := checkStrOrPosInt(creator, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(newCommunism{
		Creator:     creator,
		Amount:      amount,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/communisms", content, s)
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

func (s *sdk) abortOrCloseCommunism(communismID uint64, issuer any, endpoint string) (*Communism, error) {
	body, err := s.abortSomething(communismID, issuer, endpoint)
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

func (s *sdk) AbortCommunism(communismID uint64, issuer any) (*Communism, error) {
	return s.abortOrCloseCommunism(communismID, issuer, "/v1/communisms/abort")
}

func (s *sdk) CloseCommunism(communismID uint64, issuer any) (*Communism, error) {
	return s.abortOrCloseCommunism(communismID, issuer, "/v1/communisms/close")
}

func (s *sdk) changeCommunismParticipation(communismID uint64, user any, endpoint string) (*Communism, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(communismParticipationUpdate{
		ID:   communismID,
		User: user,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, s)
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

func (s *sdk) IncreaseCommunismParticipation(communismID uint64, user any) (*Communism, error) {
	return s.changeCommunismParticipation(communismID, user, "/v1/communisms/increaseParticipation")
}

func (s *sdk) DecreaseCommunismParticipation(communismID uint64, user any) (*Communism, error) {
	return s.changeCommunismParticipation(communismID, user, "/v1/communisms/decreaseParticipation")
}

func (s *sdk) NewPoll(user any, issuer any, variant string) (*Poll, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}
	if err := checkStrOrPosInt(issuer, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(newPoll{
		User:    user,
		Issuer:  issuer,
		Variant: variant,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/polls", content, s)
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

func (s *sdk) NewRefund(creator any, amount uint64, description string) (*Refund, error) {
	if err := checkStrOrPosInt(creator, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(newRefund{
		Creator:     creator,
		Amount:      amount,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/refunds", content, s)
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

func (s *sdk) AbortPoll(pollID uint64, issuer any) (*Poll, error) {
	body, err := s.abortSomething(pollID, issuer, "/v1/polls/abort")
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

func (s *sdk) AbortRefund(refundID uint64, issuer any) (*Refund, error) {
	body, err := s.abortSomething(refundID, issuer, "/v1/refunds/abort")
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

func (s *sdk) vote(ballotID uint64, user any, vote bool, endpoint string) ([]byte, error) {
	if err := checkStrOrPosInt(user, false); err != nil {
		return nil, err
	}

	content, err := json.Marshal(newVote{
		User:     user,
		BallotID: ballotID,
		Vote:     vote,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post(endpoint, content, s)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (s *sdk) VoteOnPollBallot(ballotID uint64, user any, vote bool) (*PollVote, error) {
	body, err := s.vote(ballotID, user, vote, "/v1/polls/vote")
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

func (s *sdk) VoteOnRefundBallot(ballotID uint64, user any, vote bool) (*RefundVote, error) {
	body, err := s.vote(ballotID, user, vote, "/v1/refunds/vote")
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

func (s *sdk) NewCallback(url string, applicationID uint64, sharedSecret string) (*Callback, error) {
	content, err := json.Marshal(newCallback{
		Url:           url,
		ApplicationID: applicationID,
		SharedSecret:  sharedSecret,
	})
	if err != nil {
		return nil, err
	}

	_, body, err := Post("/v1/callbacks", content, s)
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

func (s *sdk) DeleteCallback(id uint64) (bool, error) {
	content, err := json.Marshal(simpleID{ID: id})
	if err != nil {
		return false, err
	}

	code, _, err := Delete("/v1/callbacks", content, s)
	if err != nil {
		return code == 204, err
	}
	return code == 204, err
}
