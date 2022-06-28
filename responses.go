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
	Major uint `json:"major"`
	Minor uint `json:"minor"`
	Micro uint `json:"micro"`
}

type Status struct {
	Startup        *uint       `json:"startup"`
	ApiVersion     uint        `json:"api_version"`
	ProjectVersion VersionInfo `json:"project_version"`
	Timezone       *string     `json:"timezone"`
	Localtime      string      `json:"localtime"`
	Timestamp      uint        `json:"timestamp"`
}

type Settings struct {
	MinRefundApproves          uint `json:"min_refund_approves"`
	MinRefundDisapproves       uint `json:"min_refund_disapproves"`
	MinMembershipApproves      uint `json:"min_membership_approves"`
	MinMembershipDisapproves   uint `json:"min_membership_disapproves"`
	MaxParallelDebtors         uint `json:"max_parallel_debtors"`
	MaxSimultaneousConsumption uint `json:"max_simultaneous_consumption"`
	MaxTransactionAmount       uint `json:"max_transaction_amount"`
}

type Application struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Created uint   `json:"created"`
}

type Callback struct {
	Id            uint    `json:"id"`
	Url           string  `json:"url"`
	ApplicationId *uint   `json:"application_id"`
	SharedSecret  *string `json:"shared_secret"`
}

type Alias struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ApplicationId uint   `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type AliasDeletion struct {
	UserId  uint    `json:"user_id"`
	Aliases []Alias `json:"aliases"`
}

type User struct {
	Id         uint    `json:"id"`
	Balance    int     `json:"balance"`
	Permission bool    `json:"permission"`
	Active     bool    `json:"active"`
	External   bool    `json:"external"`
	VoucherId  *uint   `json:"voucher_id"`
	Aliases    []Alias `json:"aliases"`
	Created    uint    `json:"created"`
	Modified   uint    `json:"modified"`
}

type Transaction struct {
	Id                 uint    `json:"id"`
	Sender             User    `json:"sender"`
	Receiver           User    `json:"receiver"`
	Amount             uint    `json:"amount"`
	Reason             *string `json:"reason"`
	MultiTransactionId *uint   `json:"multi_transaction_id"`
	Timestamp          uint    `json:"timestamp"`
}

type VoucherUpdate struct {
	Debtor      User         `json:"debtor"`
	Voucher     *User        `json:"voucher"`
	Transaction *Transaction `json:"transaction"`
}

type MultiTransaction struct {
	Id           uint          `json:"id"`
	BaseAmount   uint          `json:"base_amount"`
	TotalAmount  uint          `json:"total_amount"`
	Transactions []Transaction `json:"transactions"`
	Timestamp    uint          `json:"timestamp"`
}

type Consumable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

type CommunismParticipant struct {
	UserId   uint `json:"user_id"`
	Quantity uint `json:"quantity"`
}

type Communism struct {
	Id               uint                   `json:"id"`
	Amount           uint                   `json:"amount"`
	Description      string                 `json:"description"`
	CreatorId        uint                   `json:"creator_id"`
	Active           bool                   `json:"active"`
	Created          uint                   `json:"created"`
	Modified         uint                   `json:"modified"`
	Participants     []CommunismParticipant `json:"participants"`
	MultiTransaction *MultiTransaction      `json:"multi_transaction"`
}

type Vote struct {
	Id       uint `json:"id"`
	UserId   uint `json:"user_id"`
	BallotId uint `json:"ballot_id"`
	Vote     bool `json:"vote"`
	Modified uint `json:"modified"`
}

type Poll struct {
	Id        uint   `json:"id"`
	Active    bool   `json:"active"`
	Accepted  *bool  `json:"accepted"`
	Variant   string `json:"variant"`
	User      User   `json:"user"`
	CreatorId uint   `json:"creator_id"`
	BallotId  uint   `json:"ballot_id"`
	Votes     []Vote `json:"votes"`
	Created   uint   `json:"created"`
	Modified  uint   `json:"modified"`
}

type PollVote struct {
	Poll Poll `json:"poll"`
	Vote Vote `json:"vote"`
}

type Refund struct {
	Id          uint         `json:"id"`
	Amount      uint         `json:"amount"`
	Description string       `json:"description"`
	Creator     User         `json:"creator"`
	Active      bool         `json:"active"`
	Allowed     *bool        `json:"allowed"`
	BallotId    uint         `json:"ballot_id"`
	Votes       []Vote       `json:"votes"`
	Transaction *Transaction `json:"transaction"`
	Created     *uint        `json:"created"`
	Modified    *uint        `json:"modified"`
}

type RefundVote struct {
	Refund Refund `json:"refund"`
	Vote   Vote   `json:"vote"`
}
