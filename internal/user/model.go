package user

import "github.com/darkphotonKN/online-trade/internal/models"

type UserResponse struct {
	models.BaseDBDateModel
	Email string `db:"email" json:"email"`
	Name  string `db:"name" json:"name"`
}

type UserLoginRequest struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type UserLoginResponse struct {
	RefreshToken     string `json:"refreshToken"`
	AccessToken      string `json:"accessToken"`
	AccessExpiresIn  int    `json:"accessExpiresIn"`
	RefreshExpiresIn int    `json:"refreshExpiresIn"`

	UserInfo *models.User `json:"userInfo"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}
