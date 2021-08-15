package dto

type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      int64  `json:"userId,omitempty" form:"userId,omitempty"`
}

type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      int64  `json:"userId,omitempty" form:"userId,omitempty"`
}
