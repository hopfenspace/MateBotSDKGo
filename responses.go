package MateBotSDKGo

type Error struct {
	IsError bool   `json:"error"`
	Status  int    `json:"status"`
	Method  string `json:"method"`
	Request string `json:"request"`
	Repeat  bool   `json:"repeat"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e Error) Error() string {
	return e.Message
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type VersionInfo struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Micro int `json:"micro"`
}

type Status struct {
	Startup        *int        `json:"startup"`
	ApiVersion     int         `json:"api_version"`
	ProjectVersion VersionInfo `json:"project_version"`
	Timezone       *string     `json:"timezone"`
	Localtime      string      `json:"localtime"`
	Timestamp      int         `json:"timestamp"`
}

type Settings struct {
	MinRefundApproves          int `json:"min_refund_approves"`
	MinRefundDisapproves       int `json:"min_refund_disapproves"`
	MinMembershipApproves      int `json:"min_membership_approves"`
	MinMembershipDisapproves   int `json:"min_membership_disapproves"`
	MaxParallelDebtors         int `json:"max_parallel_debtors"`
	MaxSimultaneousConsumption int `json:"max_simultaneous_consumption"`
	MaxTransactionAmount       int `json:"max_transaction_amount"`
}

type Application struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
}

type Callback struct {
	Id            int     `json:"id"`
	Url           string  `json:"url"`
	ApplicationId *int    `json:"application_id"`
	SharedSecret  *string `json:"shared_secret"`
}

type Alias struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	ApplicationId int    `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type AliasDeletion struct {
	UserId  int     `json:"user_id"`
	Aliases []Alias `json:"aliases"`
}

type User struct {
	Id         int     `json:"id"`
	Balance    int     `json:"balance"`
	Permission bool    `json:"permission"`
	Active     bool    `json:"active"`
	External   bool    `json:"external"`
	VoucherId  *int    `json:"voucher_id"`
	Aliases    []Alias `json:"aliases"`
	Created    int     `json:"created"`
	Modified   int     `json:"modified"`
}

type Transaction struct {
	Id                 int     `json:"id"`
	Sender             User    `json:"sender"`
	Receiver           User    `json:"receiver"`
	Amount             int     `json:"amount"`
	Reason             *string `json:"reason"`
	MultiTransactionId *int    `json:"multi_transaction_id"`
	Timestamp          int     `json:"timestamp"`
}

type VoucherUpdate struct {
	Debtor      User         `json:"debtor"`
	Voucher     *User        `json:"voucher"`
	Transaction *Transaction `json:"transaction"`
}

type MultiTransaction struct {
	Id           int           `json:"id"`
	BaseAmount   int           `json:"base_amount"`
	TotalAmount  int           `json:"total_amount"`
	Transactions []Transaction `json:"transactions"`
	Timestamp    int           `json:"timestamp"`
}

type Consumable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type CommunismParticipant struct {
	UserId   int `json:"user_id"`
	Quantity int `json:"quantity"`
}

type Communism struct {
	Id               int                    `json:"id"`
	Amount           int                    `json:"amount"`
	Description      string                 `json:"description"`
	CreatorId        int                    `json:"creator_id"`
	Active           bool                   `json:"active"`
	Created          int                    `json:"created"`
	Modified         int                    `json:"modified"`
	Participants     []CommunismParticipant `json:"participants"`
	MultiTransaction *MultiTransaction      `json:"multi_transaction"`
}

type Vote struct {
	Id       int  `json:"id"`
	UserId   int  `json:"user_id"`
	BallotId int  `json:"ballot_id"`
	Vote     bool `json:"vote"`
	Modified int  `json:"modified"`
}

type Poll struct {
	Id        int    `json:"id"`
	Active    bool   `json:"active"`
	Accepted  *bool  `json:"accepted"`
	Variant   string `json:"variant"`
	User      User   `json:"user"`
	CreatorId int    `json:"creator_id"`
	BallotId  int    `json:"ballot_id"`
	Votes     []Vote `json:"votes"`
	Created   int    `json:"created"`
	Modified  int    `json:"modified"`
}

type Refund struct {
	Id          int          `json:"id"`
	Amount      int          `json:"amount"`
	Description string       `json:"description"`
	Creator     User         `json:"creator"`
	Active      bool         `json:"active"`
	Allowed     *bool        `json:"allowed"`
	BallotId    int          `json:"ballot_id"`
	Votes       []Vote       `json:"votes"`
	Transaction *Transaction `json:"transaction"`
	Created     *int         `json:"created"`
	Modified    *int         `json:"modified"`
}
