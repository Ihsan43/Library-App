package utils

import (
	"library_app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserIdInContext(ctx *gin.Context) string {
	userId, exist := ctx.Get("userId")
	if !exist || userId == nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "userId not found in context")
		return ""
	}

	userIdStr, ok := userId.(string)
	if !ok || userIdStr == "" {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "invalid userId format")
		return ""
	}
	return userIdStr
}
