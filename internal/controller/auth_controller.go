package controller

import (
	"library_app/internal/service"
	"library_app/model"
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

	newRes,err := c.auhtUc.RegisterAccount(account)

	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Succesfully Created Account", newRes)
}