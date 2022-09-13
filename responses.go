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
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Created uint   `json:"created"`
}

type Callback struct {
	ID            uint    `json:"id"`
	Url           string  `json:"url"`
	ApplicationID *uint   `json:"application_id"`
	SharedSecret  *string `json:"shared_secret"`
}

type Alias struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ApplicationID uint   `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type AliasDeletion struct {
	UserID  uint    `json:"user_id"`
	Aliases []Alias `json:"aliases"`
}

type User struct {
	ID         uint    `json:"id"`
	Balance    int     `json:"balance"`
	Name       string  `json:"name"`
	Permission bool    `json:"permission"`
	Active     bool    `json:"active"`
	External   bool    `json:"external"`
	VoucherID  *uint   `json:"voucher_id"`
	Aliases    []Alias `json:"aliases"`
	Created    uint    `json:"created"`
	Modified   uint    `json:"modified"`
}

type PrivilegeLevel uint8

const (
	Disabled PrivilegeLevel = iota
	External
	Vouched
	Internal
	Permitted
)

func (u *User) Privilege() PrivilegeLevel {
	if !u.Active {
		return Disabled
	}
	if u.External {
		if u.VoucherID == nil {
			return External
		} else {
			return Vouched
		}
	}
	if !u.Permission {
		return Internal
	} else {
		return Permitted
	}
}

type Transaction struct {
	ID                 uint    `json:"id"`
	Sender             User    `json:"sender"`
	Receiver           User    `json:"receiver"`
	Amount             uint    `json:"amount"`
	Reason             *string `json:"reason"`
	MultiTransactionID *uint   `json:"multi_transaction_id"`
	Timestamp          uint    `json:"timestamp"`
}

type VoucherUpdate struct {
	Debtor      User         `json:"debtor"`
	Voucher     *User        `json:"voucher"`
	Transaction *Transaction `json:"transaction"`
}

type MultiTransaction struct {
	ID           uint          `json:"id"`
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
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
	Quantity uint   `json:"quantity"`
}

type Communism struct {
	ID               uint                   `json:"id"`
	Amount           uint                   `json:"amount"`
	Description      string                 `json:"description"`
	CreatorID        uint                   `json:"creator_id"`
	Active           bool                   `json:"active"`
	Created          uint                   `json:"created"`
	Modified         uint                   `json:"modified"`
	Participants     []CommunismParticipant `json:"participants"`
	MultiTransaction *MultiTransaction      `json:"multi_transaction"`
}

type Vote struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
	BallotID uint   `json:"ballot_id"`
	Vote     bool   `json:"vote"`
	Modified uint   `json:"modified"`
}

type PollVariant string

const (
	GetInternal     PollVariant = "get_internal"
	LooseInternal               = "loose_internal"
	GetPermission               = "get_permission"
	LoosePermission             = "loose_permission"
)

type Poll struct {
	ID        uint        `json:"id"`
	Active    bool        `json:"active"`
	Accepted  *bool       `json:"accepted"`
	Variant   PollVariant `json:"variant"`
	User      User        `json:"user"`
	CreatorID uint        `json:"creator_id"`
	BallotID  uint        `json:"ballot_id"`
	Votes     []Vote      `json:"votes"`
	Created   uint        `json:"created"`
	Modified  uint        `json:"modified"`
}

type PollVote struct {
	Poll Poll `json:"poll"`
	Vote Vote `json:"vote"`
}

type Refund struct {
	ID          uint         `json:"id"`
	Amount      uint         `json:"amount"`
	Description string       `json:"description"`
	Creator     User         `json:"creator"`
	Active      bool         `json:"active"`
	Allowed     *bool        `json:"allowed"`
	BallotID    uint         `json:"ballot_id"`
	Votes       []Vote       `json:"votes"`
	Transaction *Transaction `json:"transaction"`
	Created     *uint        `json:"created"`
	Modified    *uint        `json:"modified"`
}

type RefundVote struct {
	Refund Refund `json:"refund"`
	Vote   Vote   `json:"vote"`
}

type EventType string

const (
	ServerStarted              EventType = "server_started"
	AliasConfirmationRequested           = "alias_confirmation_requested"
	AliasConfirmed                       = "alias_confirmed"
	CommunismCreated                     = "communism_created"
	CommunismUpdated                     = "communism_updated"
	CommunismClosed                      = "communism_closed"
	PollCreated                          = "poll_created"
	PollUpdated                          = "poll_updated"
	PollClosed                           = "poll_closed"
	RefundCreated                        = "refund_created"
	RefundUpdated                        = "refund_updated"
	RefundClosed                         = "refund_closed"
	TransactionCreated                   = "transaction_created"
	VoucherUpdated                       = "voucher_updated"
	UserSoftlyDeleted                    = "user_softly_deleted"
	UserUpdated                          = "user_updated"
)

// The EventData struct may contain any of the available fields, depending on the event type
type EventData struct {
	ID                  *uint        `json:"id"`
	App                 *string      `json:"app"`
	ServerBaseUrl       *string      `json:"base_url"`
	TransactionSender   *uint        `json:"sender"`
	TransactionReceiver *uint        `json:"receiver"`
	Amount              *uint        `json:"amount"`
	Transaction         *uint        `json:"transaction"`
	Voucher             *uint        `json:"voucher"`
	User                *uint        `json:"user"`
	Participants        *uint        `json:"participants"`
	Aborted             *bool        `json:"aborted"`
	CountTransactions   *uint        `json:"transactions"`
	CurrentBallotResult *int         `json:"current_result"`
	LastBallotVote      *uint        `json:"last_vote"`
	Accepted            *bool        `json:"accepted"`
	PollVariant         *PollVariant `json:"variant"`
}

type Event struct {
	Event     EventType `json:"event"`
	Timestamp int       `json:"timestamp"`
	Data      EventData `json:"data"`
}

type EventsNotification struct {
	Number int     `json:"number"`
	Events []Event `json:"events"`
}
