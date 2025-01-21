package marsGin

import (
	"errors"
)

const UserKey = "user"

func (g *Gin) SetUser(username string) {
	g.Set(UserKey, username)
}

func (g *Gin) GetUser() (string, error) {
	value, exists := g.Get(UserKey)
	if !exists {
		return "", errors.New("user not exists")
	}
	return value.(string), nil
}
