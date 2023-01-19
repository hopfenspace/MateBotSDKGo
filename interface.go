package MateBotSDKGo

type SDK interface {
	GetThisApplicationID() uint64
	GetThisApplicationName() string
	GetCommunityUsername() *string
	GetCurrency() Currency
	FormatBalance(balance int64) string

	GetHealth() (bool, error)
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
	GetVerifiedUser(userId uint64, minimalLevel *PrivilegeLevel) (*User, error)
	IsUserConfirmed(user *User) bool

	FindSponsoringUser(issuer *User) (*User, error)
	GetCommunityBalance(issuer *User) (int64, error)

	NewUser(username string) (*User, error)
	DropInternalPrivilege(user any, issuer any) (*User, error)
	DropPermissionPrivilege(user any, issuer any) (*User, error)
	SetUsername(issuer any, newName string) (*User, error)
	SetVoucher(debtor any, voucher any, issuer any) (*VoucherUpdate, error)
	DeleteUser(userID uint64, issuer any) (*User, error)

	NewAlias(userID uint64, username string) (*Alias, error)
	NewUserWithAlias(username string) (*User, error)
	ConfirmAlias(aliasID uint64, issuer any) (*Alias, error)
	DeleteAlias(aliasID uint64, issuer any) (*AliasDeletion, error)

	SendTransaction(sender any, receiver any, amount uint64, reason string) (*Transaction, error)
	ConsumeTransaction(consumer any, amount uint64, consumable string) (*Transaction, error)

	NewCommunism(creator any, amount uint64, description string) (*Communism, error)
	AbortCommunism(communismID uint64, issuer any) (*Communism, error)
	CloseCommunism(communismID uint64, issuer any) (*Communism, error)
	IncreaseCommunismParticipation(communismID uint64, user any) (*Communism, error)
	DecreaseCommunismParticipation(communismID uint64, user any) (*Communism, error)

	NewPoll(user any, issuer any, variant string) (*Poll, error)
	AbortPoll(pollID uint64, issuer any) (*Poll, error)
	VoteOnPollBallot(ballotID uint64, user any, vote bool) (*PollVote, error)

	NewRefund(creator any, amount uint64, description string) (*Refund, error)
	AbortRefund(refundID uint64, issuer any) (*Refund, error)
	VoteOnRefundBallot(ballotID uint64, user any, vote bool) (*RefundVote, error)

	NewCallback(url string, applicationID uint64, sharedSecret string) (*Callback, error)
	DeleteCallback(id uint64) (bool, error)
}
