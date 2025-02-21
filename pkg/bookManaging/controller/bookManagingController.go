package controller

import "github.com/labstack/echo/v4"

type BookManagingController interface {
	Creating(pctx echo.Context) error
	Editing(pctx echo.Context) error
	ChangeStatus(pctx echo.Context) error
}
