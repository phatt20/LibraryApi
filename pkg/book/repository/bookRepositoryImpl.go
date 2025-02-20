package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/phatt20/LibraryApi/databases"
	"github.com/phatt20/LibraryApi/entities"
	enum "github.com/phatt20/LibraryApi/enum"
	_bookException "github.com/phatt20/LibraryApi/pkg/book/exception"
	__bookModel "github.com/phatt20/LibraryApi/pkg/book/model"
)

type BookRepositoryImpl struct {
	db     databases.Databases
	Logger echo.Logger
}

func NewbookRepositoryImpl(db databases.Databases, Logger echo.Logger) BookRepository {
	return &BookRepositoryImpl{db, Logger}
}

func (r *BookRepositoryImpl) Listing(itemFilter *__bookModel.BookFilter) ([]*entities.Book, error) {
	book := make([]*entities.Book, 0)
	query := r.db.Connect().Model(entities.Book{}).Where("Status = ?", enum.StatusAvailable)

	if itemFilter != nil {
		if itemFilter.Name != "" {
			query = query.Where("name ILIKE ?", "%"+itemFilter.Name+"%")
		}
		if itemFilter.Description != "" {
			query = query.Where("description ILIKE ?", "%"+itemFilter.Description+"%")
		}
	}

	offset := int((itemFilter.Page - 1) * itemFilter.Size)
	limit := int(itemFilter.Size)
	if err := query.Offset(offset).Limit(limit).Find(&book).Error; err != nil {
		r.Logger.Errorf("Failed to list items: %s", err.Error())
		return nil, &_bookException.ItemListing{}
	}
	return book, nil
}
