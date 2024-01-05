package service

import (
	"log"
	"time"

	"github.com/agusbasari29/xjx-code/entity"
	"github.com/agusbasari29/xjx-code/repository"
	"github.com/agusbasari29/xjx-code/request"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) CreateUser(req request.RequestAuthLogin) (entity.Users, error) {
	user := entity.Users{}
	err := smapping.FillStruct(&user, smapping.MapFields(&req))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return user, err
	}
	return newUser, err
}
