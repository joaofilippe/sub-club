package auth

type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsernameLoginRequestDTO struct {
	Username    string `json:"username"`
	AccountSlug string `json:"account_slug"`
	Password    string `json:"password"`
}

type LookupRequestDTO struct {
	Email string `json:"email"`
}

type AccountInfoDTO struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type TokenResponseDTO struct {
	Token string `json:"token"`
}
