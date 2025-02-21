package service

import (
	_bookModel "github.com/phatt20/LibraryApi/pkg/book/model"
)

type BookService interface {
	Listing(itemFilter *_bookModel.BookFilter) (*_bookModel.BookResult, error)
}
