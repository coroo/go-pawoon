package dto

type Credentials struct {
	Transactionname string `form:"transactionname"`
	Password string `form:"password"`
}

type Token struct {
	Token 	string `json:"token"`
}
