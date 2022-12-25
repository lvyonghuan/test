package model

type User struct {
	UserName string `json:"userName"`
	PassWord string `json:"passWord"`
	Token    string `json:"token"`
}
