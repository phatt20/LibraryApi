package exception

import "fmt"

type BookEditing struct {
	BookID uint64
}

func (e *BookEditing) Error() string {
	return fmt.Sprintf("editing book id : %d failed", e.BookID)
}
