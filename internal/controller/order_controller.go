package controller

import (
	"library_app/internal/service"
	"library_app/model"
	"library_app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *orderController {
	return &orderController{
		orderService: orderService,
	}
}

func (o *orderController) CreateOrder(ctx *gin.Context) {

	var payload model.Orders

	id, exist := ctx.Get("userId")
	if !exist {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "id not found in context")
		return
	}

	userId, _ := id.(string)

	payload.UserID = userId

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRes, err := o.orderService.CreateOrder(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Succesfully Create Order", newRes)
}
