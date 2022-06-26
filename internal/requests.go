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
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Amount   uint   `json:"amount"`
	Reason   string `json:"reason"`
}

type NewCommunism struct {
	Amount      uint   `json:"amount"`
	Description string `json:"description"`
	Creator     uint   `json:"creator"`
}

type NewConsumption struct {
	User       uint   `json:"user"`
	Amount     uint   `json:"amount"`
	Consumable string `json:"consumable"`
}

type NewPoll struct {
	User    int    `json:"user"`
	Issuer  int    `json:"issuer"`
	Variant string `json:"variant"`
}

type NewRefund struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Creator     int    `json:"creator"`
}

type NewVote struct {
	User     int  `json:"user"`
	BallotId int  `json:"ballot_id"`
	Vote     bool `json:"vote"`
}

type UpdateIdIssuer struct {
	Id     int `json:"id"`
	Issuer int `json:"issuer"`
}

type UpdateIdIssuerName struct {
	Id         int    `json:"id"`
	IssuerName string `json:"issuer"`
}

type UpdateIdUser struct {
	Id   int `json:"id"`
	User int `json:"user"`
}

type UpdateUserIssuer struct {
	User   int  `json:"user"`
	Issuer *int `json:"issuer"`
}

type UpdateUserNameIssuer struct {
	User   string `json:"user"`
	Issuer *int   `json:"issuer"`
}

type UpdateUserIssuerName struct {
	User   int     `json:"user"`
	Issuer *string `json:"issuer"`
}

type UpdateUserNameIssuerName struct {
	User   string  `json:"user"`
	Issuer *string `json:"issuer"`
}
