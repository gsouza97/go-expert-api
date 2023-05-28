package database

import "github.com/gsouza97/go-expert-api/internal/entity"

type UserDBInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
