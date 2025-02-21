package repository

import (
	"github.com/phatt20/LibraryApi/entities"
	enum "github.com/phatt20/LibraryApi/enum"
	_bookManagingModel "github.com/phatt20/LibraryApi/pkg/bookManaging/model"
)

type BookManagingRepository interface {
	Creating(bookEntity *entities.Book) (*entities.Book, error)
	Editing(bookID uint64, bookEditingReq *_bookManagingModel.BookEditingReq) (uint64, error)
	ChangeStatus(bookID uint64, status enum.BookStatus) error
}
