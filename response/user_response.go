package response

import (
	"time"

	"github.com/agusbasari29/xjx-code/entity"
	"gorm.io/gorm"
)

type ReesponseUserData struct {
	User       interface{} `json:"user_data"`
	Credential interface{} `json:"credential"`
}

type ResponseUser struct {
	ID             uint           `json:"id"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	UserCredential interface{}    `json:"user_credential"`
}

func ResponseUserFormatter(user entity.Users) ResponseUser {
	formatter := ResponseUser{}
	formatter.ID = user.ID
	formatter.Username = user.Username
	formatter.Email = user.Email
	formatter.CreatedAt = user.CreatedAt
	formatter.UpdatedAt = user.UpdatedAt
	formatter.DeletedAt = user.DeletedAt

	return formatter
}

func ResponseUserDataFormatter(user interface{}, credential interface{}) ReesponseUserData {
	userData := ReesponseUserData{
		User:       user,
		Credential: credential,
	}

	return userData
}
