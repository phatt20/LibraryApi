package controller

import "github.com/labstack/echo/v4"

type BookController interface {
	Listing(pctx echo.Context) error
}
