package MateBotSDKGo

type simpleID struct {
	ID uint `json:"id"`
}

type newCallback struct {
	Url           string `json:"url"`
	ApplicationID uint   `json:"application_id"`
	SharedSecret  string `json:"shared_secret"`
}

type newUser struct {
	Name string `json:"name"`
}

type newAlias struct {
	UserID        uint   `json:"user_id"`
	ApplicationID uint   `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type newTransaction struct {
	Sender   any    `json:"sender"`
	Receiver any    `json:"receiver"`
	Amount   uint   `json:"amount"`
	Reason   string `json:"reason"`
}

type newConsumption struct {
	User       any    `json:"user"`
	Amount     uint   `json:"amount"`
	Consumable string `json:"consumable"`
}

type newCommunism struct {
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
	Creator     any    `json:"creator"`
}

type newPoll struct {
	User    any    `json:"user"`
	Issuer  any    `json:"issuer"`
	Variant string `json:"variant"`
}

type newRefund struct {
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
	Creator     any    `json:"creator"`
}

type newVote struct {
	User     any  `json:"user"`
	BallotID uint `json:"ballot_id"`
	Vote     bool `json:"vote"`
}

type issuerID struct {
	ID     uint `json:"id"`
	Issuer any  `json:"issuer"`
}

type userPrivilegeDrop struct {
	User   any `json:"user"`
	Issuer any `json:"issuer"`
}

type usernameUpdate struct {
	Name   string `json:"name"`
	Issuer any    `json:"issuer"`
}

type voucherUpdate struct {
	Debtor  any `json:"debtor"`
	Voucher any `json:"voucher"`
	Issuer  any `json:"issuer"`
}

type communismParticipationUpdate struct {
	ID   uint `json:"id"`
	User any  `json:"user"`
}
