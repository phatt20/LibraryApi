package service

import (
	"github.com/phatt20/LibraryApi/entities"
	"github.com/phatt20/LibraryApi/enum"
	_bookModel "github.com/phatt20/LibraryApi/pkg/book/model"
	_bookRepository "github.com/phatt20/LibraryApi/pkg/book/repository"
	_bookManagingModel "github.com/phatt20/LibraryApi/pkg/bookManaging/model"
	_repository "github.com/phatt20/LibraryApi/pkg/bookManaging/repository"
)

type BookManagingServiceImpl struct {
	bookManagingRepository _repository.BookManagingRepository
	bookRepository         _bookRepository.BookRepository
}

func NewBookManagingServiceImpl(bookManagingRepository _repository.BookManagingRepository, bookRepository _bookRepository.BookRepository) BookManagingService {
	return &BookManagingServiceImpl{bookManagingRepository, bookRepository}
}

func (s *BookManagingServiceImpl) Creating(bookCreatingReq _bookManagingModel.BookCreatingReq) (*_bookModel.Book, error) {
	BookEntity := &entities.Book{ //รับมาเป้น dto creating  เเปลงเป้น entity เเล้วโยนเข้่า repo
		Name:        bookCreatingReq.Name,
		Description: bookCreatingReq.Description,
		Picture:     bookCreatingReq.Picture,
		Price:       bookCreatingReq.Price,
	}

	result, err := s.bookManagingRepository.Creating(BookEntity)
	if err != nil {
		return nil, err
	}

	return result.ToModel(), nil
}

func (s *BookManagingServiceImpl) Editing(bookID uint64, bookEditingReq *_bookManagingModel.BookEditingReq) (*_bookModel.Book, error) {

	_, err := s.bookManagingRepository.Editing(bookID, bookEditingReq)
	if err != nil {
		return nil, err
	}

	reslut, err := s.bookRepository.FindByID(bookID)
	if err != nil {
		return nil, err
	}
	return reslut.ToModel(), nil
}

func (s *BookManagingServiceImpl) ChangeStatus(BookID uint64, status enum.BookStatus) error {
	return s.bookManagingRepository.ChangeStatus(BookID, status)
}
