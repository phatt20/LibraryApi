package service

import (
	"github.com/phatt20/LibraryApi/enum"
	_bookModel "github.com/phatt20/LibraryApi/pkg/book/model"
	_bookManagingModel "github.com/phatt20/LibraryApi/pkg/bookManaging/model"
)

type BookManagingService interface {
	Creating(bookCreatingReq _bookManagingModel.BookCreatingReq) (*_bookModel.Book, error)
	Editing(bookID uint64, bookEditingReq *_bookManagingModel.BookEditingReq) (*_bookModel.Book, error)
	ChangeStatus(bookID uint64, status enum.BookStatus) error
}
