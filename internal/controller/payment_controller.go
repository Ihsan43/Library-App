package controller

import (
	"library_app/internal/service"
	"library_app/model"
	"library_app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentController struct {
	paymentService service.PaymentService
}

func (p *paymentController) CreatePayment(ctx *gin.Context) {
	var payment model.Payment

	if err := ctx.ShouldBindJSON(&payment); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRes, err := p.paymentService.CreatePayment(payment)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Succesfully Create Payment", newRes)

}

func NewPaymentController(paymentService service.PaymentService) *paymentController {
	return &paymentController{
		paymentService: paymentService,
	}
}
