package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/darkphotonKN/online-trade/internal/auth"
	"github.com/darkphotonKN/online-trade/internal/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetUserByIdService(id uuid.UUID) (*models.User, error) {
	return s.Repo.GetById(id)
}

func (s *UserService) CreateUserService(user models.User) error {
	hashedPw, err := s.HashPassword(user.Password)

	if err != nil {
		return fmt.Errorf("Error when attempting to hash password.")
	}

	// update user's password with hashed password.
	user.Password = hashedPw

	return s.Repo.Create(user)
}

func (s *UserService) LoginUserService(loginReq UserLoginRequest) (*UserLoginResponse, error) {
	user, err := s.Repo.GetUserByEmail(loginReq.Email)

	if err != nil {
		return nil, errors.New("Could not get user with provided email.")
	}

	// extract password, and compare hashes
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return nil, errors.New("The credentials provided was incorrect.")
	}

	// construct response with both user info and auth credentials
	accessExpiryTime := time.Minute * 60
	accessToken, err := auth.GenerateJWT(*user, auth.Access, accessExpiryTime)
	refreshExpiryTime := time.Hour * 24 * 7
	refreshToken, err := auth.GenerateJWT(*user, auth.Refresh, refreshExpiryTime)

	res := &UserLoginResponse{
		AccessToken:      accessToken,
		AccessExpiresIn:  int(accessExpiryTime),
		RefreshToken:     refreshToken,
		RefreshExpiresIn: int(refreshExpiryTime),
		UserInfo:         user,
	}

	return res, nil
}

// HashPassword hashes the given password using bcrypt.
func (s *UserService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
