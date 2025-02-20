package controller

import (
	_service "github.com/phatt20/LibraryApi/pkg/book/service"
)

type BookControllerImpl struct {
	service _service.BookService
}

func NewBookController(service _service.BookService) BookController {
	return &BookControllerImpl{service}
}
