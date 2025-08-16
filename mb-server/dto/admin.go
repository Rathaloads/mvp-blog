package dto

type LoginRequest struct {
	Email    string `json:"email" xml:"email" binding:"required,max=32"`
	Password string `json:"password" xml:"password" binding:"required,max=16,min=6"`
}

type LoginResponse struct {
	Token string `json:"token" xml:"token"`
}
