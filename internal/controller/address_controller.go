package controller

import (
	"fmt"
	"library_app/internal/service"
	"library_app/model"
	"library_app/model/dto"
	"library_app/utils/common"
	"library_app/utils/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addressController struct {
	addressService service.AddressService
}

func NewAddressController(addressService service.AddressService) *addressController {
	return &addressController{
		addressService: addressService,
	}
}

func (c *addressController) CreateAddress(ctx *gin.Context) {
	var address model.Address

	userIdStr, err := validation.ValidateCompareId(ctx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	address.UserID = userIdStr

	if err := ctx.ShouldBindJSON(&address); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRes, err := c.addressService.CreateAddress(address)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Succesfully Create Address", newRes)

}

func (c *addressController) UpdateAddress(ctx *gin.Context) {
	var payload dto.AddressDto

	_, err := validation.ValidateCompareId(ctx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")

	newRes, err := c.addressService.UpdateAddress(id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfuly Update Address", newRes)

}

func (c *addressController) GetAddress(ctx *gin.Context) {

	_, err := validation.ValidateCompareId(ctx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id := ctx.Param("id")

	address, err := c.addressService.GetAddress(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfuly Get Address", address)

}

func (c *addressController) DeleteAddress(ctx *gin.Context) {

	_, err := validation.ValidateCompareId(ctx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id := ctx.Param("id")

	if err := c.addressService.DeleteAddrees(id); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfuly Get Address", fmt.Sprintf("Address with id:%s", id))
}
