package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	__bookModel "github.com/phatt20/LibraryApi/pkg/book/model"
	_bookService "github.com/phatt20/LibraryApi/pkg/book/service"
	"github.com/phatt20/LibraryApi/pkg/custom"
)

type BookControllerImpl struct {
	bookService _bookService.BookService
}

func NewBookControllerImpl(bookService _bookService.BookService) BookController {
	return &BookControllerImpl{bookService}
}

func (c *BookControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(__bookModel.BookFilter)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	itemListModel, err := c.bookService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, itemListModel)
}
