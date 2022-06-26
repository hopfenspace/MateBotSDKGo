package internal

type IdBody struct {
	Id int `json:"id"`
}

type NewCallback struct {
	Url           string `json:"url"`
	ApplicationId int    `json:"application_id"`
	SharedSecret  string `json:"shared_secret"`
}

type NewAlias struct {
	UserId        int    `json:"user_id"`
	ApplicationId int    `json:"application_id"`
	Username      string `json:"username"`
	Confirmed     bool   `json:"confirmed"`
}

type NewTransaction struct {
	Sender   int    `json:"sender"`
	Receiver int    `json:"receiver"`
	Amount   int    `json:"amount"`
	Reason   string `json:"reason"`
}

type NewCommunism struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Creator     int    `json:"creator"`
}

type NewConsumption struct {
	User       int    `json:"user"`
	Amount     int    `json:"amount"`
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
