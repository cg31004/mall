package bo

type MemberSession struct {
	Id      string
	Account string
	Name    string
	Token   string
}

type MemberToken struct {
	Token string
}

type MemberSessionCond struct {
	Account  string
	Password string
}
