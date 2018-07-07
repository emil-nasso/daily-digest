package daily_digest

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/emil-nasso/daily-digest/util"
)

// TODO: Really really need to hash this password, eventually and extract the sessions and invalidate and avoid collisions
type User struct {
	username      string
	password      string
	subscriptions []*Subscription
	sessionKeys   []string
}

var users []User

func init() {
	users = make([]User, 0)
}

func RegisterUser(username, password string) {
	user := User{
		username:      username,
		password:      password,
		subscriptions: make([]*Subscription, 0),
		sessionKeys:   make([]string, 0),
	}
	users = append(users, user)
}

func Login(username, password string) (*string, error) {
	var user *User
	for i, u := range users {
		if u.username == username && u.password == password {
			user = &users[i]
			break
		}
	}

	if user == nil {
		return nil, errors.New("Invalid user")
	}
	b := make([]byte, 48)
	_, err := rand.Read(b)
	// TODO: Proper error handling
	if err != nil {
		return nil, err
	}
	sessionKey := base64.StdEncoding.EncodeToString(b)

	user.sessionKeys = append(user.sessionKeys, sessionKey)
	return &sessionKey, nil
}

func GetUserForSession(sessionKey string) *User {
	for _, user := range users {
		for _, s := range user.sessionKeys {
			if sessionKey == s {
				return &user
			}
		}
	}
	return nil
}

func DebugUsers() {
	util.Dd(users)
}
