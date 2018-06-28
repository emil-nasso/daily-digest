package server

type User struct {
	Id       string
	Username string
	Password string
	Sessions []*Session
}

type Session struct {
	Id   string
	Key  string
	User *User
}
