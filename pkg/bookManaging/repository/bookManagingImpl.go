package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/phatt20/LibraryApi/databases"
	"github.com/phatt20/LibraryApi/entities"
	enum "github.com/phatt20/LibraryApi/enum"
	_bookManagingException "github.com/phatt20/LibraryApi/pkg/bookManaging/exception"
	_bookManagingModel "github.com/phatt20/LibraryApi/pkg/bookManaging/model"
)

type BookManagingRepositoryImpl struct {
	db     databases.Databases
	Logger echo.Logger
}

func NewBookManagingImpl(db databases.Databases, Logger echo.Logger) BookManagingRepository {
	return &BookManagingRepositoryImpl{db, Logger}
}

func (r *BookManagingRepositoryImpl) Creating(bookEntity *entities.Book) (*entities.Book, error) {
	book := new(entities.Book)

	if err := r.db.Connect().Create(bookEntity).Scan(book).Error; err != nil {
		r.Logger.Errorf("Creating book failed: %s", err.Error())
		return nil, &_bookManagingException.BookCreating{}
	}
	return book, nil
}

func (r *BookManagingRepositoryImpl) Editing(bookID uint64, bookEditingReq *_bookManagingModel.BookEditingReq) (uint64, error) {

	if err := r.db.Connect().Model(&entities.Book{}).Where("id = ?", bookID).Updates(bookEditingReq).Error; err != nil {
		r.Logger.Errorf("Creating book failed: %s", err.Error())
		return 0, &_bookManagingException.BookEditing{BookID: bookID}
	}
	return 0, nil
}

func (r *BookManagingRepositoryImpl) ChangeStatus(bookID uint64, status enum.BookStatus) error {
	if err := r.db.Connect().Table("books").Where("id = ?", bookID).Update("status", status).Error; err != nil {
		r.Logger.Errorf("Status Changing failed: %s", err.Error())
		return &_bookManagingException.BookEditing{BookID: bookID}
	}
	return nil
}
