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

func (c *AuthController) Create(ctx *gin.Context) {
	var account model.Account

	if err := ctx.ShouldBindJSON(&account); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRes, err := c.auhtUc.RegisterAccount(account)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Succesfully Created Account", newRes)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var payload dto.AuthRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	account, err := c.auhtUc.Login(payload.Username, payload.Password)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succesfully Login", account.Username)
}
