package repository

import (
	"github.com/phatt20/LibraryApi/entities"
	__bookModel "github.com/phatt20/LibraryApi/pkg/book/model"
)

type BookRepository interface {
	Listing(itemFilter *__bookModel.BookFilter) ([]*entities.Book, error)
	FindByID(bookID uint64) (*entities.Book, error)
	Counting(itemFilter *__bookModel.BookFilter) (int64, error)
}
