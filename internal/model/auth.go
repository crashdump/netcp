package model

import (
	"encoding/json"
	"errors"
)

type Auth struct {
	Type   string     `json:"type"`
	Config AuthConfig `json:"config"`
}

type AuthConfig struct {
	Domain   string `json:"domain"`
	ClientID string `json:"client_id"`
}

func (a Auth) String() string {
	ju, _ := json.Marshal(a)
	return string(ju)
}

type Auths []Auth

func (a Auths) String() string {
	ju, _ := json.Marshal(a)
	return string(ju)
}

func (u *Auth) Validate() error {
	if u.Type == "" {
		return errors.New("type cannot be null")
	}
	return nil
}