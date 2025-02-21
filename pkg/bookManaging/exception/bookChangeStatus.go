package exception

import "fmt"

type BookChangeStatus struct {
	BookID uint64
}

func (e *BookChangeStatus) Error() string {
	return fmt.Sprintf("Changing Status book id : %d failed", e.BookID)
}
