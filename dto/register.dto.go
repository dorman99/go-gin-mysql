package dto

type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:6"`
	Username string `json:"username" form:"username" binding:"required" validate:"min:6"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
