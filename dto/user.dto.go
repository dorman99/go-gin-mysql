package dto

/**
`json:"id"` -> binding validation model
*/
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `jsong:"password" form:"password" validate:"min:6"`
}

type UserCreareDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `jsong:"password" form:"password" binding:"password" validate:"min:6"`
}
