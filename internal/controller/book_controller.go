package controller

import (
	"fmt"
	"library_app/internal/service"
	"library_app/model"
	"library_app/utils/common"
	"library_app/utils/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) *bookController {
	return &bookController{
		bookService: bookService,
	}
}

func (c *bookController) CreateBook(ctx *gin.Context) {
	var payload model.Book

	if err := validation.ValidatRoleAdmin(ctx); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRes, err := c.bookService.CreateBook(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Successfully create book", newRes)
}

func (c *bookController) GetBook(ctx *gin.Context) {

	id := ctx.Param("id")

	newRes, err := c.bookService.GetBookById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Successfully get book", newRes)

}

func (c *bookController) GetBooksWithPagination(ctx *gin.Context) {

	page, limit, err := common.GetLimitAndPage(ctx)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	books, total, err := c.bookService.FindBooks(page, limit)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return

	}

	var newBooks []any
	for _, item := range books {
		newBooks = append(newBooks, item)
	}

	paging := common.PagingInfo{
		Total: total,
		Page:  page,
		Limit: limit,
	}

	common.SendPagedResponse(ctx, "Success", newBooks, paging)

}

func (c *bookController) DeleteBook(ctx *gin.Context) {

	if err := validation.ValidatRoleAdmin(ctx); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id := ctx.Param("id")

	newRes, err := c.bookService.DeleteBookById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succesfully delete book", fmt.Sprintf("Book with id : %s is delete", newRes.ID))
}
