package entity

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `form:"id" firestore:"type:uuid"`

	Name  string `json:"name" firestore:"name"`
	Email string `json:"email" firestore:"email"`
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
	if u.Name == "" {
		return errors.New("name cannot be null")
	}
	if u.Email == "" {
		return errors.New("email cannot be null")
	}
	return nil
}