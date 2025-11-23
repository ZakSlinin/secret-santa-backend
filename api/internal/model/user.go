package model

type User struct {
	UserId       string   `json:"userId"`
	Name         string   `json:"name"`
	Username     string   `json:"username"`
	GroupsMember []string `json:"groupsMember"`
}
