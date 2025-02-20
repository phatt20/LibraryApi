package server
import(
	_repository "github.com/phatt20/LibraryApi/pkg/book/repository"
	_service "github.com/phatt20/LibraryApi/pkg/book/service"
	_controller "github.com/phatt20/LibraryApi/pkg/book/controller"

)

func(s *echoServer) initBookRouter(){
	router := s.app.Group("v1/book")
	bookRepository := _repository.NewbookRepositoryImpl(s.db,s.app.Logger)
	bookService := 

}