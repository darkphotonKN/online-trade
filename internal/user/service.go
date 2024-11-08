package user

import (
	"errors"
	"fmt"

	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
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

func (s *UserService) LoginUserService(loginReq UserLoginRequest) (*models.User, error) {

	user, err := s.Repo.GetUserByEmail(loginReq.Email)

	if err != nil {
		return nil, errors.New("Could not get user with provided email.")
	}

	// extract password, and compare hashes
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return nil, errors.New("The credentials provided was incorrect.")
	}

	return user, nil
}

// HashPassword hashes the given password using bcrypt.
func (s *UserService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
