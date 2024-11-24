package common

import (
	modelutil "library_app/utils/model_util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCreateResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusCreated, modelutil.SingleResponse{
		Status: modelutil.Status{
			Status:      true,
			Code:        http.StatusCreated,
			Description: description,
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusOK, modelutil.SingleResponse{
		Status: modelutil.Status{
			Status:      true,
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, description string) {
	ctx.JSON(code, modelutil.SingleResponse{
		Status: modelutil.Status{
			Status:      false,
			Code:        code,
			Description: description,
		},
	})
}

func SendPagedResponse(ctx *gin.Context, description string, data []any, paging any) {
	ctx.JSON(http.StatusOK, modelutil.PagedResponse{
		Status: modelutil.Status{
			Status:      true,
			Code:        http.StatusOK,
			Description: description,
		},
		Data:   data,
		Paging: paging,
	})
}

