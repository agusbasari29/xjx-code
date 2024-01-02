package routes

import (
	"github.com/agusbasari29/xjx-code/database"
	"github.com/agusbasari29/xjx-code/helper"
)

type AuthRoutes struct{}

func (r AuthRoutes) Route() []helper.Route {
	db := database.SetupDBConnection()
	useRepo
}
