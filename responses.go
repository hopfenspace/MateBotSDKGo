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

type Settings struct {
	MinRefundApproves          uint64 `json:"min_refund_approves"`
	MinRefundDisapproves       uint64 `json:"min_refund_disapproves"`
	MinMembershipApproves      uint64 `json:"min_membership_approves"`
	MinMembershipDisapproves   uint64 `json:"min_membership_disapproves"`
	MaxParallelDebtors         uint64 `json:"max_parallel_debtors"`
	MaxSimultaneousConsumption uint64 `json:"max_simultaneous_consumption"`
	MaxTransactionAmount       uint64 `json:"max_transaction_amount"`
}

type Application struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Created uint64 `json:"created"`
}

type Callback struct {
	ID            uint64  `json:"id"`
	Url           string  `json:"url"`
	ApplicationID *uint64 `json:"application_id"`
}

type Alias struct {
	ID            uint64 `json:"id"`
	UserID        uint64 `json:"user_id"`
	ApplicationID uint64 `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type AliasDeletion struct {
	UserID  uint64  `json:"user_id"`
	Aliases []Alias `json:"aliases"`
}

type User struct {
	ID         uint64  `json:"id"`
	Balance    int64   `json:"balance"`
	Name       string  `json:"name"`
	Permission bool    `json:"permission"`
	Active     bool    `json:"active"`
	External   bool    `json:"external"`
	VoucherID  *uint64 `json:"voucher_id"`
	Aliases    []Alias `json:"aliases"`
	Created    uint64  `json:"created"`
	Modified   uint64  `json:"modified"`
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
	ID                 uint64  `json:"id"`
	Sender             User    `json:"sender"`
	Receiver           User    `json:"receiver"`
	Amount             uint64  `json:"amount"`
	Reason             *string `json:"reason"`
	MultiTransactionID *uint64 `json:"multi_transaction_id"`
	Timestamp          uint64  `json:"timestamp"`
}

type VoucherUpdate struct {
	Debtor      User         `json:"debtor"`
	Voucher     *User        `json:"voucher"`
	Transaction *Transaction `json:"transaction"`
}

type MultiTransaction struct {
	ID           uint64        `json:"id"`
	BaseAmount   uint64        `json:"base_amount"`
	TotalAmount  uint64        `json:"total_amount"`
	Transactions []Transaction `json:"transactions"`
	Timestamp    uint64        `json:"timestamp"`
}

type Consumable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
}

type CommunismParticipant struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"user_name"`
	Quantity uint64 `json:"quantity"`
}

type Communism struct {
	ID               uint64                 `json:"id"`
	Amount           uint64                 `json:"amount"`
	Description      string                 `json:"description"`
	CreatorID        uint64                 `json:"creator_id"`
	Active           bool                   `json:"active"`
	Created          uint64                 `json:"created"`
	Modified         uint64                 `json:"modified"`
	Participants     []CommunismParticipant `json:"participants"`
	MultiTransaction *MultiTransaction      `json:"multi_transaction"`
}

type Vote struct {
	ID       uint64 `json:"id"`
	UserID   uint64 `json:"user_id"`
	Username string `json:"user_name"`
	BallotID uint64 `json:"ballot_id"`
	Vote     bool   `json:"vote"`
	Modified uint64 `json:"modified"`
}

type PollVariant string

const (
	GetInternal     PollVariant = "get_internal"
	LooseInternal               = "loose_internal"
	GetPermission               = "get_permission"
	LoosePermission             = "loose_permission"
)

type Poll struct {
	ID        uint64      `json:"id"`
	Active    bool        `json:"active"`
	Accepted  *bool       `json:"accepted"`
	Variant   PollVariant `json:"variant"`
	User      User        `json:"user"`
	CreatorID uint64      `json:"creator_id"`
	BallotID  uint64      `json:"ballot_id"`
	Votes     []Vote      `json:"votes"`
	Created   uint64      `json:"created"`
	Modified  uint64      `json:"modified"`
}

type PollVote struct {
	Poll Poll `json:"poll"`
	Vote Vote `json:"vote"`
}

type Refund struct {
	ID          uint64       `json:"id"`
	Amount      uint64       `json:"amount"`
	Description string       `json:"description"`
	Creator     User         `json:"creator"`
	Active      bool         `json:"active"`
	Allowed     *bool        `json:"allowed"`
	BallotID    uint64       `json:"ballot_id"`
	Votes       []Vote       `json:"votes"`
	Transaction *Transaction `json:"transaction"`
	Created     *uint64      `json:"created"`
	Modified    *uint64      `json:"modified"`
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
	ID                  *uint64      `json:"id"`
	App                 *string      `json:"app"`
	ServerBaseUrl       *string      `json:"base_url"`
	TransactionSender   *uint64      `json:"sender"`
	TransactionReceiver *uint64      `json:"receiver"`
	Amount              *uint64      `json:"amount"`
	Transaction         *uint64      `json:"transaction"`
	Voucher             *uint64      `json:"voucher"`
	User                *uint64      `json:"user"`
	Participants        *uint64      `json:"participants"`
	Aborted             *bool        `json:"aborted"`
	CountTransactions   *uint64      `json:"transactions"`
	CurrentBallotResult *int         `json:"current_result"`
	LastBallotVote      *uint64      `json:"last_vote"`
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
