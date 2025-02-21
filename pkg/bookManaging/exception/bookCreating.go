package exception

type BookCreating struct{}

func (e *BookCreating) Error() string {
	return "Creating book failed"
}
