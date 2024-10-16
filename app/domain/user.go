package domain

type User struct {
	Username string `json:"username"`
	IsActive string `json:"is_active"`
}

type Auth struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
