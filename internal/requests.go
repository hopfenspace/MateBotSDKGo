package internal

type IdBody struct {
	Id uint `json:"id"`
}

type NewCallback struct {
	Url           string `json:"url"`
	ApplicationId uint   `json:"application_id"`
	SharedSecret  string `json:"shared_secret"`
}

type NewAlias struct {
	UserId        uint   `json:"user_id"`
	ApplicationId uint   `json:"application_id"`
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
	BallotId uint `json:"ballot_id"`
	Vote     bool `json:"vote"`
}

type IssuerIdBody struct {
	Id     uint `json:"id"`
	Issuer any  `json:"issuer"`
}

type UserPrivilegeDrop struct {
	User   any `json:"user"`
	Issuer any `json:"issuer"`
}

type VoucherUpdateRequest struct {
	Debtor  any `json:"debtor"`
	Voucher any `json:"voucher"`
}

type CommunismParticipationUpdate struct {
	Id   uint `json:"id"`
	User any  `json:"user"`
}
