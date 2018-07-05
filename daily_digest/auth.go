package daily_digest

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// TODO: Really really need to hash this password, eventually and extract the sessions and invalidate and avoid collisions
type User struct {
	username      string
	password      string
	subscriptions []*Subscription
	sessionKeys   []string
}

type AuthService struct {
	users []User
}

func (s *AuthService) Register(username, password string) {
	user := User{
		username:      username,
		password:      password,
		subscriptions: make([]*Subscription, 0),
		sessionKeys:   make([]string, 0),
	}
	s.users = append(s.users, user)
}

func (s *AuthService) Login(username, password string) (*string, error) {
	var user *User
	for i, u := range s.users {
		if u.username == username && u.password == password {
			user = &s.users[i]
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

func (s *AuthService) UserForSession(sessionKey string) *User {
	for _, user := range s.users {
		for _, s := range user.sessionKeys {
			if sessionKey == s {
				return &user
			}
		}
	}
	return nil
}
