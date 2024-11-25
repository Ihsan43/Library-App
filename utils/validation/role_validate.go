package validation

import (
	"fmt"
	"library_app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidatRoleAdmin(ctx *gin.Context) error {
	role, exist := ctx.Get("role")
	if !exist {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "admin not found in context")
		return fmt.Errorf("admin not found in context")
	}

	roleStr, _ := role.(string)

	if roleStr != "admin" {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "You are not authorized to access this resource")
		return fmt.Errorf("You are not authorized to access this resource")
	}

	return nil
}
