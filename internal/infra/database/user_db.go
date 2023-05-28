package database

import (
	"github.com/gsouza97/go-expert-api/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (db *UserDB) CreateUser(user *entity.User) error {
	return db.DB.Create(user).Error
}

func (db *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
