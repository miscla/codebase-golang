package services

import (
	"github.com/miscla/codebase-golang/internal/models"
	"github.com/miscla/codebase-golang/internal/repository"
)

// UserService defines business logic for users
type UserService interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Delete(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
