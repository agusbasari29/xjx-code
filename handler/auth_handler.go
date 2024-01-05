package handler

import (
	"net/http"

	"github.com/agusbasari29/xjx-code/entity"
	"github.com/agusbasari29/xjx-code/helper"
	"github.com/agusbasari29/xjx-code/request"
	"github.com/agusbasari29/xjx-code/response"
	"github.com/agusbasari29/xjx-code/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type authHandler struct {
	//service

	userService service.UserService
	jwtService  service.JWTService
}

func NewAuthHadnler(userService service.UserService, jwtService service.JWTService) *authHandler {
	return &authHandler{userService, jwtService}
}

func (h *authHandler) Login(c *gin.Context) {
	req := request.RequestAuthLogin{}

	//input error
	err := c.ShouldBind(&req)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	//validation
	validationErr := validate.Struct(req)
	if validationErr != nil {
		errorFormatter := helper.ErrorFormatter(validationErr)
		errorMessage := helper.M{"error": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	credential := h.userService.VerifyCredential(req.Username, req.Password)
	if v, ok := credential.(entity.Users); ok {
		generatedToken := h.jwtService.GenerateToken(v)
		userData := response.ResponseUserFormatter(v)
		data := response.ResponseUserDataFormatter(userData, generatedToken)
		response := helper.ResponseFormatter(http.StatusOK, "success", "User successfully Loggedin.", data)
		c.JSON(http.StatusOK, response)

	}
}
