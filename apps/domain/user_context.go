package domain

type UserContext struct {
	Id       string `json:"id"`
	FullName string `json:"fullname"`
	Exp      int64  `json:"exp"`
}
