package service

import (
	"errors"
	"forum/internal/repository"
	"forum/models"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

const salt = "hfg43eghkdfjkn3264k"

var (
	err         error
	emptyUser   models.User
	user        models.User
	ErrMail     = errors.New("email not registered")
	ErrPassword = errors.New("wrong password")
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user models.User) error {
	if err := checkUsername(user.Username); err != nil {
		return err
	}
	if err := checkEmail(user.Email); err != nil {
		return err
	}
	if err := checkPassword(user.Password); err != nil {
		return err
	}
	user.Password, err = getHash(user.Password)
	if err != nil {
		return err
	}
	if err := s.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func getHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("service: %s\n", err)
		return "", err
	}
	return string(hash), nil
}

func (s *AuthService) GenerateToken(email, password string) (models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return emptyUser, ErrMail
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+salt)); err != nil {
		return emptyUser, ErrPassword
	}
	token, err := uuid.NewV4()
	if err != nil {
		return emptyUser, err
	}
	user.Token = token.String()
	user.TokenDuration = time.Now().Add(12 * time.Hour)
	if err := s.repo.SaveToken(user.Username, user.Token, user.TokenDuration); err != nil {
		return emptyUser, err
	}
	return user, nil
}

func (s *AuthService) GetUserByToken(token string) (models.User, error) {
	user, err := s.repo.GetUserByToken(token)
	if err != nil {
		return emptyUser, err
	}
	return user, nil
}

func (s *AuthService) DeleteToken(token string) error {
	return s.repo.DeleteToken(token)
}
