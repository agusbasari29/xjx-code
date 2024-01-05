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

type UserService interface {
	CreateUser(req request.RequestAuthLogin) (entity.Users, error)
	VerifyCredential(username string, password string) interface{}
}

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

func (s *userService) VerifyCredential(username string, password string) interface{} {
	res := s.repository.GetByUsername(username)
	if v, ok := res.(entity.Users); ok {
		comparedPassword := comparePassword(v.Password, password)
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPassword))
	return err == nil
}
