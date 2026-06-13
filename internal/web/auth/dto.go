package auth

type LoginRequestDTO struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccountSlug string `json:"account_slug"` // optional — empty means system login
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
