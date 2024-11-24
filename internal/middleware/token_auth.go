package middleware

import (
	"fmt"
	"library_app/utils/common"
	"library_app/utils/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthorizationHeader struct {
	AuthHeader string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var au AuthorizationHeader

		if err := ctx.ShouldBindHeader(&au); err != nil {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, "Failed to read Authorization header")
			ctx.Abort()
			return
		}

		if au.AuthHeader == "" {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "Missing Authorization header")
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(au.AuthHeader, "Bearer ")
		token = strings.TrimSpace(token)
		if token == "" {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid Authorization token")
			ctx.Abort()
			return
		}

		claims, err := security.VerifyToken(token)
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %s", err.Error()))
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims["userId"])
		ctx.Set("role", claims["role"])

		ctx.Next()
	}
}

func ValidationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, existRole := ctx.Get("role")
		if !existRole {
			common.SendErrorResponse(ctx, http.StatusForbidden, "Role not found in context")
			ctx.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, "Role in context is not a string")
			ctx.Abort()
			return
		}

		jwtUserID, existUserID := ctx.Get("userId")
		if !existUserID {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, "User ID not found in context")
			ctx.Abort()
			return
		}

		jwtUserIDStr, ok := jwtUserID.(string)
		if !ok {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, "User ID in context is not a string")
			ctx.Abort()
			return
		}

		paramID := ctx.Param("id")

		if roleStr == "admin" {
			ctx.Next()
			return
		}

		if paramID != jwtUserIDStr {
			common.SendErrorResponse(ctx, http.StatusForbidden, "You are not authorized to access this resource")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
