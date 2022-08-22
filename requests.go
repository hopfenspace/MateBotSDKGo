package MateBotSDKGo

type IDBody struct {
	ID uint `json:"id"`
}

type NewCallback struct {
	Url           string `json:"url"`
	ApplicationID uint   `json:"application_id"`
	SharedSecret  string `json:"shared_secret"`
}

type NewAlias struct {
	UserID        uint   `json:"user_id"`
	ApplicationID uint   `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type NewTransaction struct {
	Sender   any    `json:"sender"`
	Receiver any    `json:"receiver"`
	Amount   uint   `json:"amount"`
	Reason   string `json:"reason"`
}

type NewConsumption struct {
	User       any    `json:"user"`
	Amount     uint   `json:"amount"`
	Consumable string `json:"consumable"`
}

type NewCommunism struct {
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
	Creator     any    `json:"creator"`
}

type NewPoll struct {
	User    any    `json:"user"`
	Issuer  any    `json:"issuer"`
	Variant string `json:"variant"`
}

type NewRefund struct {
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
	Creator     any    `json:"creator"`
}

type NewVote struct {
	User     any  `json:"user"`
	BallotID uint `json:"ballot_id"`
	Vote     bool `json:"vote"`
}

type IssuerIDBody struct {
	ID     uint `json:"id"`
	Issuer any  `json:"issuer"`
}

type UserPrivilegeDrop struct {
	User   any `json:"user"`
	Issuer any `json:"issuer"`
}

type VoucherUpdateRequest struct {
	Debtor  any `json:"debtor"`
	Voucher any `json:"voucher"`
	Issuer  any `json:"issuer"`
}

type CommunismParticipationUpdate struct {
	ID   uint `json:"id"`
	User any  `json:"user"`
}
