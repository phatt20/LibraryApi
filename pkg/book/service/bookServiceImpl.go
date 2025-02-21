package service

import (
	"github.com/phatt20/LibraryApi/entities"
	_bookModel "github.com/phatt20/LibraryApi/pkg/book/model"
	_repository "github.com/phatt20/LibraryApi/pkg/book/repository"
)

type BookServiceImpl struct {
	bookRepository _repository.BookRepository
}

func NewBookServiceImpl(bookRepository _repository.BookRepository) BookService {
	return &BookServiceImpl{bookRepository}
}
func (s *BookServiceImpl) Listing(itemFilter *_bookModel.BookFilter) (*_bookModel.BookResult, error) {
	itemList, err := s.bookRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}
	counting, err := s.bookRepository.Counting(itemFilter)
	if err != nil {
		return nil, err
	}
	totalPage := s.totalPage(counting, itemFilter.Size)
	result := s.toItemResult(itemList, itemFilter.Page, totalPage)

	return result, nil
}

func (s *BookServiceImpl) totalPage(totalBook int64, size int64) int64 {
	totalPage := totalBook / size
	if totalBook%size != 0 {
		totalPage++
	}
	return totalPage
}

func (s *BookServiceImpl) toItemResult(itemEntityList []*entities.Book, page, totalpage int64) *_bookModel.BookResult {
	BookList := make([]*_bookModel.Book, 0)

	for _, book := range itemEntityList {
		BookList = append(BookList, book.ToModel())
	}
	return &_bookModel.BookResult{
		Books: BookList,
		Paginate: _bookModel.PaginateResult{
			Page:      page,
			TotalPage: totalpage,
		},
	}
}
