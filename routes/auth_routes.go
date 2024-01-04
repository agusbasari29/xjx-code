package routes

import (
	"github.com/agusbasari29/xjx-code/database"
	"github.com/agusbasari29/xjx-code/handler"
	"github.com/agusbasari29/xjx-code/helper"
	"github.com/agusbasari29/xjx-code/repository"
	"github.com/agusbasari29/xjx-code/service"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct{}

func (r AuthRoutes) Route() []helper.Route {
	db := database.SetupDBConnection("")
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	jwtService := service.NewJWTService()
	authHandler := handler.NewAuthHadnler(userService, jwtService)

	return []helper.Route{
		{
			Method:  "POST",
			Path:    "/login",
			Handler: []gin.HandlerFunc{authHandler.Login()},
		},
	}
}
