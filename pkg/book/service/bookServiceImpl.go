package service

import (
	_repository "github.com/phatt20/LibraryApi/pkg/book/repository"
)

type BookServiceImpl struct {
	bookRepository _repository.BookRepository
}

func NewBookServiceImpl(bookRepository _repository.BookRepository) BookService {
	return &BookServiceImpl{bookRepository}
}
