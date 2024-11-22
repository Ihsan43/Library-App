package controller

import (
	"library_app/internal/service"
	"library_app/model"
	"library_app/model/dto"
	"library_app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auhtUc service.AuthService
}

func NewAuthController(auhtUc service.AuthService) *AuthController {
	return &AuthController{
		auhtUc: auhtUc,
	}
}

func (c *AuthController) CreateUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRes, err := c.auhtUc.Register(user)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Succesfully Created Account", newRes)
}

func (c *AuthController) LoginUser(ctx *gin.Context) {

	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.auhtUc.Login(payload.Username, payload.Password)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfully Login", token)
}
