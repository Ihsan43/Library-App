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

func ValidateCompareId(ctx *gin.Context) (string, error) {
	id, exist := ctx.Get("userId")
	if !exist {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "id not found in context")
		return "", fmt.Errorf("admin not found in context")
	}

	IdStr, _ := id.(string)

	idParam := ctx.Param("id")

	if IdStr != idParam {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "You are not authorized to access this resource")
		return "", fmt.Errorf("You are not authorized to access this resource")
	}

	return IdStr, nil
}
