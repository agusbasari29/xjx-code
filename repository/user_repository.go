package repository

import (
	"github.com/agusbasari29/xjx-code/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.Users) (entity.Users, error)
	GetByUsername(username string) interface{}
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user entity.Users) (entity.Users, error) {
	err := r.db.Raw("", user).Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetByUsername(username string) interface{} {
	user := entity.Users{}
	res := r.db.Raw("SELECT * FROM users WHERE username=@Username", map[string]interface{}{"Username": username}).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
