package models

type User struct {
	UserId       int32  `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}
