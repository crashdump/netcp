package repository

import (
	"github.com/crashdump/netcp/internal/model"

	"github.com/google/uuid"
)

type UserRepository struct {
}

func (u *UserRepository) FindAll() model.Users {
	var users model.Users
	Firebase() //.Find(&users)

	return users
}

func (u *UserRepository) FindByID(id uuid.UUID) model.User {
	var user model.User
	Firebase() //.First(&user, id)

	return user
}

func (u *UserRepository) Save(user model.User) model.User {
	Firebase() //.Save(&user)

	return user
}

func (u *UserRepository) Delete(user model.User) {
	Firebase() //.Delete(&user)
}