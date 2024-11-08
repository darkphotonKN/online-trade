package user

import "github.com/darkphotonKN/ecommerce-server-go/internal/models"

type UserResponse struct {
	models.BaseDBDateModel
	Email string `db:"email" json:"email"`
	Name  string `db:"name" json:"name"`
}

type UserLoginRequest struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json: "password"`
}
