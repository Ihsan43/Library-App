package controller

import (
	"library_app/internal/service"
	"library_app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *transactionController {
	return &transactionController{
		TransactionService: transactionService,
	}
}

func (c *transactionController) GetTransactionHistories(ctx *gin.Context) {
	id, exist := ctx.Get("userId")
	if !exist {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, "User ID not found in context")
		return
	}

	IdStr, ok := id.(string)
	if !ok || IdStr == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Panggil service untuk mendapatkan riwayat transaksi
	histories, err := c.TransactionService.GetTransactionHistoriesByUser(IdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"code":  "FETCH_FAILED",
		})
		return
	}

	// Berikan response
	ctx.JSON(http.StatusOK, gin.H{"data": histories})
}
