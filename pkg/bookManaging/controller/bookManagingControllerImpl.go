package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/phatt20/LibraryApi/enum"
	_bookManagingModel "github.com/phatt20/LibraryApi/pkg/bookManaging/model"
	_bookManagingService "github.com/phatt20/LibraryApi/pkg/bookManaging/serviceBook"
	"github.com/phatt20/LibraryApi/pkg/custom"
)

type BookManagingControllerImpl struct {
	bookManagingService _bookManagingService.BookManagingService
}

func NewBookManagingControllerImpl(bookManagingService _bookManagingService.BookManagingService) BookManagingController {
	return &BookManagingControllerImpl{bookManagingService}
}

func (c *BookManagingControllerImpl) Creating(pctx echo.Context) error {
	BookCreatingReq := new(_bookManagingModel.BookCreatingReq)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(BookCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	book, err := c.bookManagingService.Creating(*BookCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusCreated, book)
}

func (c *BookManagingControllerImpl) Editing(pctx echo.Context) error {
	BookID, err := c.getBookID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}
	BookCreatingReq := new(_bookManagingModel.BookEditingReq)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(BookCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}
	book, err := c.bookManagingService.Editing(BookID, BookCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, book)

}
func (c *BookManagingControllerImpl) ChangeStatus(pctx echo.Context) error {
	bookID, err := c.getBookID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	var req struct {
		Status enum.BookStatus `json:"status"`
	}
	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(req); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}
	if err := c.bookManagingService.ChangeStatus(bookID, req.Status); err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.NoContent(http.StatusNoContent)

}

func (c *BookManagingControllerImpl) getBookID(pctx echo.Context) (uint64, error) {
	bookID := pctx.Param("BookID")
	BookIDint64, err := strconv.ParseUint(bookID, 10, 64)
	if err != nil {
		return 0, err
	}

	return BookIDint64, nil
}
