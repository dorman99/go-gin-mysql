package dto

type LoginDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"passsword" binding:"required"`
}
