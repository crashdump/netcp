package model

import (
	"encoding/json"
	"errors"
)

type User struct {
	Model

	Auth0ID  string `json:"-" gorm:"auth0_id"`
	APIToken string `json:"-" gorm:"api_token"`

	Name  string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
}

func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

type Users []User

func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

func (u *User) Validate() error {
	if u.Auth0ID == "" {
		return errors.New("auth0_id cannot be null")
	}
	return nil
}