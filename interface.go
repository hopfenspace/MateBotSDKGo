package MateBotSDKGo

type SDK interface {
	FormatBalance(balance int) string

	GetStatus() (*Status, error)
	GetSettings() (*Settings, error)
	GetAliases(filter map[string]string) ([]*Alias, error)
	GetApplications(filter map[string]string) ([]*Application, error)
	GetCallbacks(filter map[string]string) ([]*Callback, error)
	GetCommunisms(filter map[string]string) ([]*Communism, error)
	GetConsumables(filter map[string]string) ([]*Consumable, error)
	GetPolls(filter map[string]string) ([]*Poll, error)
	GetRefunds(filter map[string]string) ([]*Refund, error)
	GetTransactions(filter map[string]string) ([]*Transaction, error)
	GetUsers(filter map[string]string) ([]*User, error)
	GetVotes(filter map[string]string) ([]*Vote, error)

	GetUser(userIdOrUsername any, extendedFilter *map[string]string) (*User, error)
	FindSponsoringUser(issuer *User) (*User, error)
	GetCommunityBalance(issuer *User) (int, error)

	NewUser(username string) (*User, error)
	DropInternalPrivilege(user any, issuer any) (*User, error)
	DropPermissionPrivilege(user any, issuer any) (*User, error)
	SetUsername(issuer any, newName string) (*User, error)
	SetVoucher(debtor any, voucher any, issuer any) (*VoucherUpdate, error)
	DeleteUser(userID uint, issuer any) (*User, error)

	NewAlias(userID uint, username string) (*Alias, error)
	NewUserWithAlias(username string) (*User, error)
	ConfirmAlias(aliasID uint, issuer any) (*Alias, error)
	DeleteAlias(aliasID uint, issuer any) (*AliasDeletion, error)

	SendTransaction(sender any, receiver any, amount uint, reason string) (*Transaction, error)
	ConsumeTransaction(consumer any, amount uint, consumable string) (*Transaction, error)

	NewCommunism(creator any, amount uint, description string) (*Communism, error)
	AbortCommunism(communismID uint, issuer any) (*Communism, error)
	CloseCommunism(communismID uint, issuer any) (*Communism, error)
	IncreaseCommunismParticipation(communismID uint, user any) (*Communism, error)
	DecreaseCommunismParticipation(communismID uint, user any) (*Communism, error)

	NewPoll(user any, issuer any, variant string) (*Poll, error)
	AbortPoll(pollID uint, issuer any) (*Poll, error)
	VoteOnPollBallot(ballotID uint, user any, vote bool) (*PollVote, error)

	NewRefund(creator any, amount uint, description string) (*Refund, error)
	AbortRefund(refundID uint, issuer any) (*Refund, error)
	VoteOnRefundBallot(ballotID uint, user any, vote bool) (*RefundVote, error)

	NewCallback(url string, applicationID uint, sharedSecret string) (*Callback, error)
	DeleteCallback(id uint) (bool, error)
}
