package dto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type GetCurrentUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
}
