package exception

type ItemCounting struct {
}

func (e *ItemCounting) Error() string {
	return "counting item failed"
}
