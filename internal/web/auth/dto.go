package auth

type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponseDTO struct {
	Token string `json:"token"`
}
