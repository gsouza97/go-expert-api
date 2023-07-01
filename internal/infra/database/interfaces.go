package database

import "github.com/gsouza97/go-expert-api/internal/entity"

type UserDBInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductDBInterface interface {
	CreateProduct(product *entity.Product) error
	FindAll(page int, limit int, sort string) ([]*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
