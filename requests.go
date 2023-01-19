package MateBotSDKGo

type simpleID struct {
	ID uint64 `json:"id"`
}

type newCallback struct {
	Url           string `json:"url"`
	ApplicationID uint64 `json:"application_id"`
	SharedSecret  string `json:"shared_secret"`
}

type newUser struct {
	Name string `json:"name"`
}

type newAlias struct {
	UserID        uint64 `json:"user_id"`
	ApplicationID uint64 `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type newTransaction struct {
	Sender   any    `json:"sender"`
	Receiver any    `json:"receiver"`
	Amount   uint64 `json:"amount"`
	Reason   string `json:"reason"`
}

type newConsumption struct {
	User       any    `json:"user"`
	Amount     uint64 `json:"amount"`
	Consumable string `json:"consumable"`
}

type newCommunism struct {
	Amount      uint64 `json:"amount"`
	Description string `json:"description"`
	Creator     any    `json:"creator"`
}

type newPoll struct {
	User    any    `json:"user"`
	Issuer  any    `json:"issuer"`
	Variant string `json:"variant"`
}

type newRefund struct {
	Amount      uint64 `json:"amount"`
	Description string `json:"description"`
	Creator     any    `json:"creator"`
}

type newVote struct {
	User     any    `json:"user"`
	BallotID uint64 `json:"ballot_id"`
	Vote     bool   `json:"vote"`
}

type issuerID struct {
	ID     uint64 `json:"id"`
	Issuer any    `json:"issuer"`
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
	ID   uint64 `json:"id"`
	User any    `json:"user"`
}
