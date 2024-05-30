package service

import (
	"time"
	"tirigo/internal/user-management/domain"
	"tirigo/internal/user-management/dtos"
	"tirigo/internal/user-management/repository"
	"tirigo/pkg/jwt"
	"tirigo/pkg/redis"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user dtos.UserRegisterParameter) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	newUser := domain.User{
		ID:        uuid.NewString(),
		Username:  user.Username,
		Password:  string(hashedPassword),
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.repo.Save(newUser)
	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}

func (s *UserService) Authenticate(userInfo dtos.UserAuthParameter) (string, error) {

	user, err := s.repo.FindByUsername(userInfo.Username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password))
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	err = redis.Client.Set(token, user.Username, time.Hour*72).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *UserService) Logout(token string) error {
	err := redis.Client.Del(token).Err()
	if err != nil {
		return err
	}
	return nil
}
