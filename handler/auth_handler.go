package handler

import "github.com/gin-gonic/gin"

type authHandler struct {
	//service
	/*
		user service
		jwt service
	*/
}

func NewAuthHadnler( /*service*/ ) *authHandler {
	return &authHandler{ /*service*/ }
}

func (h *authHandler) Login(c *gin.Context) {

}
