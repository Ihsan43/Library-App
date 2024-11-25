package controller

import (
	"fmt"
	"library_app/internal/service"
	"library_app/model/dto"
	"library_app/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUSerController(userService service.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetUserId(ctx *gin.Context) {
	userId := ctx.Param("id")

	userRes, err := c.userService.FindUserById(userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfully Get User", userRes)
}

func (c *userController) UpdatedUserById(ctx *gin.Context) {
	var payload dto.UserDto

	userId := ctx.Param("id")

	user, err := c.userService.UpdatedUser(userId, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succesfully Update User", user)
}

func (c *userController) GetUsersWithPagination(ctx *gin.Context) {
	// Mengambil query parameter page dan limit dari URL

	page, limit, err := common.GetLimitAndPage(ctx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Mendapatkan data dan total menggunakan service
	users, total, err := c.userService.FindUsers(page, limit)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Membuat objek PagingInfo
	var newUsers []any
	for _, item := range users {
		newUsers = append(newUsers, item)
	}

	// Membuat objek PagingInfo
	paging := common.PagingInfo{
		Total: total,
		Page:  page,
		Limit: limit,
	}

	// Menggunakan SendPagedResponse untuk mengirimkan response
	common.SendPagedResponse(ctx, "Success", newUsers, paging)

}

func (c *userController) DeleteUserID(ctx *gin.Context) {

	userId := ctx.Param("id")

	if userId == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id not found")
		return
	}

	res, err := c.userService.DeleteUserById(userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfully delete user", fmt.Sprintf("Account user with id : %s is delete", res.ID))
}
