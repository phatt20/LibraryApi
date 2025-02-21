package model

type BookCreatingReq struct {
	AdminId     string
	Name        string `json:"name" validate:"required,max=64"`
	Description string `json:"description" validate:"required,max=128"`
	Picture     string `json:"picture" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
}

type BookEditingReq struct {
	AdminId     string
	Name        string `json:"name" validate:"omitempty,max=64"`
	Description string `json:"description" validate:"omitempty,max=128"`
	Picture     string `json:"picture" validate:"omitempty"`
	Price       uint   `json:"price" validate:"omitempty"`
}
