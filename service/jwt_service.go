package service

import (
	"os"
	"time"

	"github.com/agusbasari29/xjx-code/entity"
	"github.com/agusbasari29/xjx-code/response"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(user entity.Users) response.ResponseCredential
}

type jwtService struct {
	secret string
	issuer string
}

type jwtCustomClaim struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer: "xjx",
		secret: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "apaajahlah"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(user entity.Users) response.ResponseCredential {
	claim := &jwtCustomClaim{}
	claim.UserID = user.ID
	claim.Email = user.Email
	claim.Username = user.Username
	claim.ExpiresAt = time.Now().AddDate(1, 0, 0).Unix()
	claim.Issuer = j.issuer
	claim.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	t, err := token.SignedString([]byte(j.secret))
	if err != nil {
		panic(err)
	}
	credential := response.ResponseCredential{}
	credential.UserID = claim.UserID
	credential.Token = t
	credential.Issuer = claim.Issuer
	credential.IssuedAt = claim.IssuedAt
	credential.ExpiresAt = claim.ExpiresAt

	return credential
}
