package server

import (
	_controller "github.com/phatt20/LibraryApi/pkg/book/controller"
	_repository "github.com/phatt20/LibraryApi/pkg/book/repository"
	_service "github.com/phatt20/LibraryApi/pkg/book/service"
)

func (s *echoServer) initBookRouter() {
	router := s.app.Group("LibrarySoCool/book")
	bookRepository := _repository.NewbookRepositoryImpl(s.db, s.app.Logger)
	bookService := _service.NewBookServiceImpl(bookRepository)
	bookController := _controller.NewBookControllerImpl(bookService)

	router.GET("", bookController.Listing)

}
