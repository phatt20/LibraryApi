package server

import (
	_Bookrepository "github.com/phatt20/LibraryApi/pkg/book/repository"
	_controller "github.com/phatt20/LibraryApi/pkg/bookManaging/controller"
	_repository "github.com/phatt20/LibraryApi/pkg/bookManaging/repository"
	_service "github.com/phatt20/LibraryApi/pkg/bookManaging/serviceBook"
)

func (s *echoServer) initBookRouterManaging() {
	router := s.app.Group("LibrarySoCool/bookManaging")
	bookRepository := _Bookrepository.NewbookRepositoryImpl(s.db, s.app.Logger)
	bookManagingRepository := _repository.NewBookManagingRepositoryImpl(s.db, s.app.Logger)
	bookManagingService := _service.NewBookManagingServiceImpl(bookManagingRepository, bookRepository)
	bookManagingController := _controller.NewBookManagingControllerImpl(bookManagingService)

	router.PATCH("/ChangingStatus/:bookID", bookManagingController.ChangeStatus)
	router.PATCH("/Editing/:bookID", bookManagingController.Editing) //อย่าลืมทำ params
	router.POST("/Creating", bookManagingController.Creating)

}
