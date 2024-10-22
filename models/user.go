package models

type UserBase struct {
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	IDModel
	TimestampModel
	UserBase
	Password string `json:"password"`
}

type UserAuth struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
