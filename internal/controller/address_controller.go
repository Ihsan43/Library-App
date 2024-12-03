package controller

import (
	"fmt"
	"library_app/internal/service"
	"library_app/model/dto"
	"library_app/utils"
	"library_app/utils/common"
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
	var payload dto.AddressRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := utils.GetUserIdInContext(ctx)

	payload.UserId = userId

	newRes, err := c.addressService.CreateAddress(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Successfully created address", newRes)
}

func (c *addressController) UpdateAddress(ctx *gin.Context) {
	var payload dto.AddressRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := utils.GetUserIdInContext(ctx)

	address, _ := c.addressService.FindAddressByUserId(userId)
	if address.UserID != userId {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "You are not authorized to access this resource")
		return
	}

	id := ctx.Param("id")

	res, err := c.addressService.UpdateAddress(id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfuly Update Address", res)

}

func (c *addressController) GetAddress(ctx *gin.Context) {

	userId := utils.GetUserIdInContext(ctx)

	address, _ := c.addressService.FindAddressByUserId(userId)
	if address.UserID != userId {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "You are not authorized to access this resource")
		return
	}

	id := ctx.Param("id")

	res, err := c.addressService.GetAddress(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfuly Get Address", res)

}

func (c *addressController) DeleteAddress(ctx *gin.Context) {

	userId := utils.GetUserIdInContext(ctx)

	address, _ := c.addressService.FindAddressByUserId(userId)
	if address.UserID != userId {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "You are not authorized to access this resource")
		return
	}

	id := ctx.Param("id")

	if err := c.addressService.DeleteAddrees(id); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfuly Get Address", fmt.Sprintf("Address with id:%s", id))
}

func (c *addressController) GetAddresses(ctx *gin.Context) {

	userId := utils.GetUserIdInContext(ctx)

	addresses, err := c.addressService.FindAddresses(userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var newAddresses []any
	for _, item := range addresses {
		newAddresses = append(newAddresses, item)
	}

	common.SendMultipleResponse(ctx, "Successfuly Get Addresses", newAddresses)
}
