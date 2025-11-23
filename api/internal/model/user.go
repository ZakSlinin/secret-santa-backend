package model

// d UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//                       user_id TEXT NOT NULL,
//                       name TEXT NOT NULL,
//                       username TEXT NOT NULL,
//                       groups_member TEXT[] DEFAULT '{}'

type User struct {
	UserId       string   `json:"userId"`
	Name         string   `json:"name"`
	Username     string   `json:"username"`
	GroupsMember []string `json:"groupsMember"`
}
