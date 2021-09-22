package dto

type LoginDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"passsword" binding:"required"`
}

type LoginResponseDto struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Refresh  string `json:"Refresh"`
}
