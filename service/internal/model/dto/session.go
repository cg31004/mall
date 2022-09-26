package dto

type MemberSessionCond struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type MemberToken struct {
	Token string `json:"token"`
}
